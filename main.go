package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/Pacerino/pr0music/acrcloud"
	"github.com/Pacerino/pr0music/odesli"
	"github.com/Pacerino/pr0music/pr0gramm"
	"github.com/Pacerino/pr0music/shazam"
	"github.com/mileusna/crontab"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const maxWorkers = 10

type SauceSession struct {
	session *pr0gramm.Session
	db      *gorm.DB
	msgChan chan pr0gramm.Message
	acr     *acrcloud.Recognizer
	after   pr0gramm.Timestamp
}

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	if len(os.Getenv("LOG_LEVEL")) > 0 {
		switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
		case "error":
			logrus.SetLevel(logrus.ErrorLevel)
		case "fatal":
			logrus.SetLevel(logrus.FatalLevel)
		case "info":
			logrus.SetLevel(logrus.InfoLevel)
		case "debug":
			logrus.SetLevel(logrus.DebugLevel)
		default:
			logrus.SetLevel(logrus.InfoLevel)
		}
	}

	for _, env := range []string{"DB_HOST", "DB_USER", "DB_PASS", "DB_DATABASE", "DB_PORT", "SHAZAM_ENDPOINT"} {
		if len(os.Getenv(env)) == 0 {
			logrus.Fatal(fmt.Sprintf("Missing %s from environment", env))
		}
	}
}

func main() {
	db, err := connectDB(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	))
	if err != nil {
		logrus.WithError(err).Error("Error while connecting to a database!")
	}

	apiKey := os.Getenv("ACR_API_KEY")
	if len(apiKey) == 0 {
		log.Fatal("Missing ACR_API_KEY from environment")
	}

	apiSecret := os.Getenv("ACR_API_SECRET")
	if len(apiSecret) == 0 {
		log.Fatal("Missing ACR_API_SECRET from environment")
	}

	apiHost := os.Getenv("ACR_API_HOST")
	if len(apiHost) == 0 {
		log.Fatal("Missing ACR_API_HOST from environment")
	}

	acrConfig := map[string]string{
		"access_key":     apiKey,
		"access_secret":  apiSecret,
		"host":           apiHost,
		"recognize_type": acrcloud.ACR_OPT_REC_AUDIO,
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
		acr:     acrcloud.NewRecognizer(acrConfig),
		msgChan: make(chan pr0gramm.Message),
		after:   pr0gramm.Timestamp{Time: time.Unix(1623837600, 0)},
	}

	ctab := crontab.New()
	err = ctab.AddJob(os.Getenv("CRONJOB"), ss.getBotComments)
	if err != nil {
		logrus.WithError(err).Fatal("Could not add cronjob!")
	}

	//ss.getBotComments()

	ctx, cancel := context.WithCancel(context.Background())
	var cwg sync.WaitGroup
	cwg.Add(1)
	go ss.commentWorker(ctx, &cwg)

	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for msg := range ss.msgChan {
				ss.handleMessage(&msg)
			}
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// cancel the context to close the comment worker
	cancel()
	// wait for comment worker to finish
	cwg.Wait()

	// close channel and wait for workers to finish
	close(ss.msgChan)
	wg.Wait()
}

func (s *SauceSession) commentWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for range time.Tick(5 * time.Second) {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msgResp, err := s.session.GetComments()
		logrus.Debug("check Pr0gramm comments")
		if err != nil {
			logrus.WithError(err).Info("failed fetching comments")
			continue
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

			// Check if bot was pinged by an unwanted individual
			unwantedUsers := strings.Split(os.Getenv("BANNED_USERS"), ",")
			if slices.Contains(unwantedUsers, msg.Name) {
				continue
			}

			logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).
				Debug(fmt.Sprintf("Bot was marked by %s", msg.Name))

			// create a copy of the message to take a valid pointer
			s.msgChan <- msg
		}
	}
}

func (s *SauceSession) handleMessage(msg *pr0gramm.Message) {
	var item Items
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

	message, tags, dbItem, err := s.findSong(msg)
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

	_, err = s.session.AddTag(msg.ItemID, tags)
	if err != nil {
		logrus.WithError(err).Error("Could not add tags")
		return
	}

	logrus.Info("Comment written")
}

func (s *SauceSession) findSong(msg *pr0gramm.Message) (string, []string, *Items, error) {
	url := fmt.Sprintf("https://vid.pr0gramm.com/%s.mp4", strings.Split(msg.Thumb, ".")[0])
	resp, err := http.Head(url)
	if err != nil {
		return "", nil, nil, fmt.Errorf("fetching thumb url: %v", err)
	}

	/* if resp.StatusCode == http.StatusNotFound {
		return "Sag mal, raffst du dat nicht? Dit ist kein Video! Nur Idioten im Internet...", &Items{ItemID: msg.ItemID}, nil
	} */

	if resp.StatusCode != http.StatusOK {
		return "", nil, nil, fmt.Errorf("invalid status %q: %v", resp.Status, url)
	}

	cstBer, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return "", nil, nil, fmt.Errorf("getting timezone: %v", err)
	}
	dt := time.Now().In(cstBer).Format("02.01.2006 um 15:04")

	meta, err := s.detectMusic(url)
	if err != nil {
		return "", nil, nil, err
	}

	if meta != nil {
		// Metadaten gefunden
		logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).Debug("Metadata was found")
		meta.Url = fmt.Sprintf("https://pr0sauce.info/%v", msg.ItemID) // Set URL with the ItemID, creates shorter links!
		dbItem := Items{
			ItemID: msg.ItemID,
			Title:  meta.Title,
			Album:  meta.Album,
			Artist: meta.Artist,
			Url:    meta.Url,
			AcrID:  meta.AcrID,
			Metadata: ItemMetadata{
				SpotifyURL: meta.Links.Spotify,
				SpotifyID:  meta.IDs.Spotify,
				YoutubeURL: meta.Links.YouTube,
				YoutubeID:  meta.IDs.YouTube,
			},
		}

		message := fmt.Sprintf("Es wurden folgende Informationen dazu gefunden:\n%s - %s\nAus dem Album: %s\n\nHier ist ein Link: %s\nZeitpunkt der Überprüfung %s",
			meta.Title,
			meta.Artist,
			meta.Album,
			meta.Url,
			dt,
		)
		tags := []string{meta.Title, meta.Artist, fmt.Sprintf("%s - %s", meta.Artist, meta.Title)}

		return message, tags, &dbItem, nil
	}

	// Keine Metadaten gefunden
	logrus.WithFields(logrus.Fields{"item_id": msg.ItemID}).Debug("No metadata found")
	message := fmt.Sprintf("Es wurden keine Informationen zu dem Lied gefunden\n\nZeitpunkt der Überprüfung %s", dt)
	return message, nil, &Items{ItemID: msg.ItemID}, nil
}

func (s *SauceSession) getBotComments() error {
	s.after = pr0gramm.Timestamp{Time: time.Unix(1623837600, 0)}
	for {
		data, err := s.session.GetUserComments("Sauce", 15, int(s.after.Unix()))
		if err != nil {
			return err
		}
		for _, c := range data.Comments {
			if !strings.Contains(c.Content, "Es wurden") {
				continue
			}
			comm := Comments{
				CommentID: int(c.Id),
				Up:        c.Up,
				Down:      c.Down,
				Content:   c.Content,
				Created:   &c.Created.Time,
				ItemID:    int(c.ItemId),
				Thumb:     c.Thumbnail,
			}
			s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "comment_id"}},
				UpdateAll: true,
			}).Create(&comm)
			logrus.Infof("Inserted Comment with ID: %d", c.Id)
		}
		if !data.HasNewer {
			break
		}
		s.after = data.Comments[len(data.Comments)-1].Created
	}
	return nil
}

func (s *SauceSession) convertToAudio(url string) (*ACRRecognitionResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("downloading video: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status %q: %v", resp.Status, url)
	}

	reader, writer := io.Pipe()
	go func() {
		err := fluentffmpeg.NewCommand("").
			PipeInput(resp.Body).
			OutputFormat("mp3").
			PipeOutput(writer).
			Run()
		if err != nil {
			logrus.WithError(writer.CloseWithError(err)).Error("Error while extracting the audio track")
		}
		writer.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	var song ACRRecognitionResult
	resultString, _, err := s.acr.RecognizeByFileBuffer(buf.Bytes(), 0, 30, nil)
	if err != nil {
		return nil, fmt.Errorf("recognizing music: %v", reader.CloseWithError(err))
	}
	json.Unmarshal([]byte(resultString), &song)
	return &song, nil
}

func (s *SauceSession) detectMusic(url string) (*RecognizedMetadata, error) {
	// Try it first with Shazam
	shazamInfo, err := shazam.Recognize(url)
	if err != nil {
		return nil, err
	} else if shazamInfo.Track.Title != "" {
		appleLink, err := shazam.GetAppleMusicLink(shazamInfo)
		if err != nil {
			return nil, err
		}
		links, err := odesli.GetLinks(appleLink)
		if err != nil {
			return nil, err
		}
		album, err := shazam.GetAlbum(shazamInfo)
		if err != nil {
			return nil, err
		}
		// Found with Shazam!
		m := &RecognizedMetadata{
			Title:  shazamInfo.Track.Title,
			Album:  album,
			Artist: shazamInfo.Track.Subtitle,
		}
		if links.LinksByPlatform.Spotify.URL != "" {
			idslice := strings.Split(links.LinksByPlatform.Spotify.EntityUniqueID, "::")
			m.Links.Spotify = links.LinksByPlatform.Spotify.URL
			m.IDs.Spotify = idslice[1]
		}

		if links.LinksByPlatform.Youtube.URL != "" {
			idslice := strings.Split(links.LinksByPlatform.Youtube.EntityUniqueID, "::")
			m.Links.YouTube = links.LinksByPlatform.Youtube.URL
			m.IDs.YouTube = idslice[1]
		}
		return m, nil
	} else {
		// Found nothing with Shazam, Try with ACRCloud
		// TODO: If Shazam is reliable enough, remove ACRCloud
		recognitionResult, err := s.convertToAudio(url)
		if err != nil {
			return nil, err
		}
		if len(recognitionResult.Metadata.Music) > 0 {
			acrInfo := recognitionResult.Metadata.Music[0]
			if len(acrInfo.Title) > 0 {
				m := &RecognizedMetadata{
					Title:  acrInfo.Title,
					Album:  acrInfo.Album.Name,
					Artist: acrInfo.Artists[0].Name,
					AcrID:  acrInfo.Acrid,
				}

				if acrInfo.ExternalMetadata.Spotify.Track.ID != "" {
					m.Links.Spotify = fmt.Sprintf("https://open.spotify.com/track/%s", acrInfo.ExternalMetadata.Spotify.Track.ID)
					m.IDs.Spotify = acrInfo.ExternalMetadata.Spotify.Track.ID
				}

				if acrInfo.ExternalMetadata.Youtube.Vid != "" {
					m.Links.YouTube = fmt.Sprintf("https://www.youtube.com/watch?v=%s", acrInfo.ExternalMetadata.Youtube.Vid)
					m.IDs.YouTube = acrInfo.ExternalMetadata.Youtube.Vid
				}

				return m, nil
			}
		}
	}
	return nil, nil
}
