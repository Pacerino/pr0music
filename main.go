package main

import (
	"fmt"
	"main/database"
	"main/pr0gramm"
	"main/recognition"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type SauceSession struct {
	session *pr0gramm.Session
	db      *database.DB
	msgChan chan pr0gramm.Message
}

func main() {
	godotenv.Load()
	logrus.SetLevel(logrus.DebugLevel)
	db, err := database.Connect()
	if err != nil {
		logrus.WithError(err).Error("Error while connecting to a database!")
	}

	session := pr0gramm.NewSession(http.Client{Timeout: 10 * time.Second})
	if resp, err := session.Login(os.Getenv("PR0_USER"), os.Getenv("PR0_PASSWORD")); err != nil {
		logrus.WithError(err).Fatal("Error logging in to pr0gramm")
		return
	} else {
		if !resp.Success {
			logrus.Fatal("Error logging in to pr0gramm")
			return
		}
	}

	ss := SauceSession{
		session: session,
		db:      db,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go ss.commentWorker(&wg)

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go ss.detectWorker(&wg)
	}

	wg.Wait()
}

func (s *SauceSession) commentWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	msgResp, err := s.session.GetComments()
	logrus.Debug("check Pr0gramm comments")
	if err != nil {
		logrus.WithError(err).Info("failed fetching comments")
		return
	}

	for _, msg := range msgResp.Messages {
		// Skip already read comments
		if msg.Read == 1 {
			continue
		}

		// Check if bot was pinged
		if !strings.Contains(strings.ToLower(msg.Message), "@sauce") {
			continue
		}

		logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).
			Debug(fmt.Sprintf("Bot was marked by %s", msg.Name))

		// create a copy of the message to take a valid pointer
		s.msgChan <- msg
	}
}

func (s *SauceSession) detectWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range s.msgChan {
		s.handleMessage(&msg)
	}
}

func (s *SauceSession) handleMessage(msg *pr0gramm.Message) {
	var item database.Items
	err := s.db.Find(&item, "item_id", msg.ItemID).Error
	if err != nil {
		logrus.WithError(err).Error("Could not check database for post")
		return
	}

	if item.ID != 0 {
		var message string
		//Post ist in der Datenbank
		if len(item.Title) > 0 {
			// Es ist ein Titel vorhanden, sende Nachricht mit den Metadaten
			logrus.WithFields(logrus.Fields{"item_id": item.ItemID}).Debug("Post has already been queried, information available")
			message = fmt.Sprintf("Hallo %s,\n\nDu hast bei https://pr0gramm.com/new/%d nach der Musik gefragt.\nJemand hat bereits danach gefragt, daher erhälst du hier nur eine Kopie.\n\nTitel: %s\nAlbum: %s\nArtist: %s\n\nHier ist ein Link: %s",
				msg.Name,
				msg.ItemID,
				item.Title,
				item.Album,
				item.Artist,
				item.Url,
			)
		} else {
			// es ist kein Titel vorhanden, sende Nachricht ohne Metadaten
			logrus.WithFields(logrus.Fields{"item_id": item.ID}).Debug("Post has already been queried, no information available")
			message = fmt.Sprintf("Hallo %s,\n\nDu hast bei https://pr0gramm.com/new/%d nach der Musik gefragt.\nLeider wurden keine Informationen gefunden.",
				msg.Name,
				msg.ItemID,
			)
		}

		_, err := s.session.SendMessage(msg.Name, message)
		if err != nil {
			logrus.WithError(err).Error("Could not send private message to user")
		}

		logrus.Info("Private message sent to user")
		return
	}

	// Post ist nicht in der Datenbank
	logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).Debug("Post has never been queried, searching for the music")

	sourceURL, err := fetchThumb(msg.Thumb)
	if err != nil {
		logrus.WithError(err).Error("Could not verify the video URL of the post")
		return
	}

	message, dbItem, err := findSong(msg, sourceURL)
	if err != nil {
		logrus.WithError(err).Error("could not fetch song metdata")
		return
	}

	if err := s.db.Create(dbItem).Error; err != nil {
		logrus.WithError(err).Error("Error saving metadata to the database!")
		return
	}

	_, err = s.session.PostComment(msg.ItemID, message, msg.ID)
	if err != nil {
		logrus.WithError(err).Error("Could not post comment")
		return
	}

	logrus.Info("Comment written")
	return
}

func fetchThumb(thumb string) (string, error) {
	url := fmt.Sprintf("https://vid.pr0gramm.com/%s.mp4", strings.Split(thumb, ".")[0])
	resp, err := http.Head(url)
	if err != nil {
		return "", fmt.Errorf("fetching thumb url: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid status %q: %v", resp.Status, url)
	}

	return url, nil
}

func findSong(msg *pr0gramm.Message, sourceURL string) (string, *database.Items, error) {
	dt := time.Now().Format("02.01.2006 um 15:04")
	meta, err := recognition.DetectMusic(sourceURL)
	if err != nil {
		return "", nil, err
	}

	if meta != nil {
		// Metadaten gefunden
		logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).Debug("Metadata was found")
		dbItem := database.Items{
			ItemID: msg.ItemID,
			Title:  meta.Title,
			Album:  meta.Album,
			Artist: meta.Artist,
			Url:    meta.Url,
			Metadata: database.Metadata{
				SpotifyURL: meta.Links.Spotify,
				SpotifyID:  meta.IDS.Spotify,
				DeezerURL:  meta.Links.Deezer,
				DeezerID:   strconv.Itoa(meta.IDS.Deezer),
			},
		}

		message := fmt.Sprintf("Es wurden folgende Informationen dazu gefunden:\n%s - %s\nAus dem Album: %s\n\nHier ist ein Link: %s\nZeitpunkt der Überprüfung %s",
			meta.Title,
			meta.Album,
			meta.Album,
			meta.Url,
			dt,
		)

		return message, &dbItem, nil
	}

	// Keine Metadaten gefunden
	logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).Debug("No metadata found")
	message := fmt.Sprintf("Es wurden keine Informationen zu dem Lied gefunden\n\nZeitpunkt der Überprüfung %s", dt)
	dbItem := database.Items{
		ItemID: msg.ItemID,
	}

	return message, &dbItem, nil
}
