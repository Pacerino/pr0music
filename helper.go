package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func ResolveThumb(thumb string) (string, error) {
	s := strings.Split(thumb, ".")
	url := fmt.Sprintf("https://vid.pr0gramm.com/%s.mp4", s[0])
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Error("Could not verify the URL")
	} else if resp.StatusCode != 200 {
		err = errors.New("status code is not 200")
		log.WithError(err)
		return "", err
	}
	return url, nil
}
