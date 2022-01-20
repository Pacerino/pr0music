package main

import (
	"fmt"
	"main/database"
	"main/pr0gramm"
	"main/recognition"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type SauceSession struct {
	Pr0 *pr0gramm.Session
	DB  *database.DB
}

func main() {
	godotenv.Load()
	db, err := database.Connect()
	if err != nil {
		log.WithError(err).Error("Error while connecting to a database!")
	}
	session := pr0gramm.NewSession(http.Client{Timeout: 10 * time.Second})
	if resp, err := session.Login(os.Getenv("PR0_USER"), os.Getenv("PR0_PASSWORD")); err != nil {
		log.WithError(err).Fatal("Error logging in to pr0gramm")
		return
	} else {
		if !resp.Success {
			log.Fatal("Error logging in to pr0gramm")
			return
		}
	}
	pr0 := SauceSession{Pr0: session, DB: db}
	DoEvery(5*time.Second, pr0.ticker)
}

func (sauce *SauceSession) ticker(t time.Time) {
	msgarr, err := sauce.Pr0.GetComments()
	if err != nil {
		log.WithError(err)
	}
	for _, msg := range msgarr.Messages {
		// Alle Kommentare
		if msg.Read != 1 {
			// Ungelesene Kommentare
			if strings.Contains(strings.ToLower(msg.Message), "@sauce") {
				// Bot wurde Markiert
				log.Debug("Bot was marked by %s", msg.Name)
				var item database.Items
				sauce.DB.Find(&item, "item_id", msg.ItemID)
				if item.ItemID > 0 {
					//Post ist in der Datenbank
					if len(item.Title) > 0 {
						// Es ist ein Titel vorhanden, sende Nachricht mit den Metadaten
						log.WithFields(log.Fields{"item_id": item.ID}).Debug("Post has already been queried, information available")
						message := fmt.Sprintf("Hallo %s,\n\nDu hast bei https://pr0gramm.com/new/%d nach der Musik gefragt.\nJemand hat bereits danach gefragt, daher erhälst du hier nur eine Kopie.\n\nTitel: %s\nAlbum: %s\nArtist: %s\n\nHier ist ein Link: %s",
							msg.Name,
							msg.ItemID,
							item.Title,
							item.Album,
							item.Artist,
							item.Url,
						)
						_, err := sauce.Pr0.SendMessage(msg.Name, message)
						if err != nil {
							log.WithError(err).Error("Could not send private message to user")
						} else {
							log.Info("Private message sent to user")
						}
					} else {
						// es ist kein Titel vorhanden, sende Nachricht ohne Metadaten
						log.WithFields(log.Fields{"item_id": item.ID}).Debug("Post has already been queried, no information available")
						message := fmt.Sprintf("Hallo %s,\n\nDu hast bei https://pr0gramm.com/new/%d nach der Musik gefragt.\nLeider wurden keine Informationen gefunden.",
							msg.Name,
							msg.ItemID,
						)
						_, err := sauce.Pr0.SendMessage(msg.Name, message)
						if err != nil {
							log.WithError(err).Error("Could not send private message to user")
						} else {
							log.Info("Private message sent to user")
						}
					}
				} else {
					// Post ist nicht in der Datenbank
					log.WithFields(log.Fields{"item_id": msg.ItemID}).Debug("Post has never been queried, searching for the music")
					sourceURL, err := ResolveThumb(msg.Thumb)
					if err != nil {
						log.WithError(err).Error("Could not verify the video URL of the post")
					} else {
						dt := time.Now()
						meta := recognition.DetectMusic(sourceURL)
						if len(meta.Title) > 0 {
							// Metadaten gefunden
							log.WithFields(log.Fields{"item_id": msg.ItemID}).Debug("Metadata was found")
							message := fmt.Sprintf("Es wurden folgende Informationen dazu gefunden:\n%s - %sAus dem Album: %s\n\nHier ist ein Link: %s\nZeitpunkt der Überprüfung %s",
								meta.Title,
								meta.Album,
								meta.Album,
								meta.Url,
								dt.Format("02.01.2006 um 15:04"),
							)
							_, err := sauce.Pr0.PostComment(msg.ItemID, message, msg.ID)
							if err != nil {
								log.WithError(err).Error("Could not post comment")
							} else {
								log.Info("Comment written")
							}
						} else {
							// Keine Metadaten gefunden
							log.WithFields(log.Fields{"item_id": msg.ItemID}).Debug("No metadata found")
							message := fmt.Sprintf("Es wurden keine Informationen zu dem Lied gefunden\n\nZeitpunkt der Überprüfung %s", dt.Format("02.01.2006 um 15:04"))
							_, err := sauce.Pr0.PostComment(msg.ItemID, message, msg.ID)
							if err != nil {
								log.WithError(err).Error("Could not post comment")
							} else {
								log.Info("Comment written")
							}
						}
					}
				}
			}
		}
	}
}
