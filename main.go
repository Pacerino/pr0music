package main

import (
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
	Session *pr0gramm.Session
	DB      *database.DB
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
	pr0 := SauceSession{Session: session, DB: db}
	DoEvery(5*time.Second, pr0.ticker)
}

func (sauce *SauceSession) ticker(t time.Time) {
	msgarr, err := sauce.Session.GetComments()
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
					} else {
						// es ist kein Titel vorhanden, sende Nachricht ohne Metadaten
						log.WithFields(log.Fields{"item_id": item.ID}).Debug("Post has already been queried, no information available")
					}
				} else {
					// Post ist nicht in der Datenbank
					log.WithFields(log.Fields{"item_id": msg.ItemID}).Debug("Post has never been queried, searching for the music")
					sourceURL, err := ResolveThumb(msg.Thumb)
					if err != nil {
						log.WithError(err).Error("Could not verify the video URL of the post")
					} else {
						_ = recognition.DetectMusic(sourceURL)
					}
				}
			}
		}
	}
}
