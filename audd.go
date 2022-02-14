package main

import (
	"fmt"
	"github.com/AudDMusic/audd-go"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (s *SauceSession) AnalyzeVideo(url string) (*audd.RecognitionResult, error) {
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
			logrus.WithError(err).Error("Error while extracting the audio track")
		}
	}()

	song, err := s.audd.RecognizeByFile(reader, "apple_music,deezer,spotify", nil)
	if err != nil {
		return nil, fmt.Errorf("recognizing music: %v", err)
	}

	return &song, nil
}

type RecognizedMetadata struct {
	// Title from the recognized song
	Title string
	// Album from the recognized song
	Album string
	// Artist from the recognized song
	Artist string
	// URL to Audd.io from the recognized song
	Url string
	// Respective links to known streaming platforms
	Links MetadataLinks
	// Respective IDs to known streaming platforms
	IDs MetadataIDs
}

type MetadataLinks struct {
	// Deezer link
	Deezer string
	// Spotify link
	Spotify string
	// Apple Music link
	AppleMusic string
}

type MetadataIDs struct {
	// Deezer ID
	Deezer int
	// Spotify ID
	Spotify string
}
