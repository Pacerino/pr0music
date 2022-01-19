package main

import (
	"main/pr0gramm"
	"main/recognition"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Pr0Service struct {
	Session *pr0gramm.Session
	After   pr0gramm.Timestamp
}

func main() {
	godotenv.Load()
	log.Info("Download")
	_ = recognition.DetectMusic("https://vid.pr0gramm.com/2022/01/18/df0e313f122115d6.mp4")
	/* session := pr0gramm.NewSession(http.Client{Timeout: 10 * time.Second})
	if resp, err := session.Login(os.Getenv("PR0_USER"), os.Getenv("PR0_PASSWORD")); err != nil {
		log.WithError(err).Fatal("Could not login into pr0gramm!")
		return
	} else {
		if !resp.Success {
			log.WithError(err).Fatal("Could not login into pr0gramm!")
			return
		}
	}
	pr0 := Pr0Service{Session: session, After: pr0gramm.Timestamp{time.Unix(1623837600, 0)}}
	DoEvery(5*time.Second, pr0.ticker) */
}

/* func (pr0 *Pr0Service) ticker(t time.Time) {
	msgarr, err := pr0.Session.GetComments()
	if err != nil {
		log.WithError(err)
	}
	for _, msg := range msgarr.Messages {
		// Alle Kommentare
		if msg.Read != 1 {
			// Ungelesene Kommentare
			if strings.Contains(strings.ToLower(msg.Message), "@sauce") {
				// Bot wurde Markiert!
			}
		}
	}
}
*/
