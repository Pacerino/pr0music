package recognition

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/AudDMusic/audd-go"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	log "github.com/sirupsen/logrus"
)

func DetectMusic(url string) (data Metadata) {
	buffer := downloadVideo(url)
	soundBuffer := extractSound(buffer)
	songInfo := uploadToService(soundBuffer)
	mbIds := make([]string, 0, len(songInfo.MusicBrainz))
	for _, mb := range songInfo.MusicBrainz {
		mbIds = append(mbIds, mb.ID)
	}
	meta := Metadata{
		Title:          songInfo.Title,
		Album:          songInfo.Album,
		Artist:         songInfo.Artist,
		Url:            songInfo.SongLink,
		MusicBrainzIDS: mbIds,
		Links: Links{
			AppleMusic: songInfo.AppleMusic.URL,
			Deezer:     songInfo.Deezer.Link,
			Spotify:    songInfo.Spotify.URI,
		},
		IDS: IDS{
			Deezer:  songInfo.Deezer.ID,
			Spotify: songInfo.Spotify.ID,
		},
	}
	log.Info(meta)
	return meta
}

func uploadToService(soundBuff []byte) audd.RecognitionResult {
	api_token := os.Getenv("AUDD_API_TOKEN")
	if len(api_token) == 0 {
		log.Fatal("Missing AUDD_API_TOKEN from environment!")
	}
	client := audd.NewClient("test")
	song, err := client.RecognizeByFile(bytes.NewReader(soundBuff), "apple_music,deezer,spotify", nil)
	if err != nil {
		log.WithError(err)
	}
	return song
}

func downloadVideo(url string) (output []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err)
	}
	defer resp.Body.Close()

	outputBuffer := new(bytes.Buffer)
	_, err = io.Copy(outputBuffer, resp.Body)
	if err != nil {
		log.WithError(err)
	}

	return outputBuffer.Bytes()
}

func extractSound(input []byte) (output []byte) {
	inputBuff := bytes.NewBuffer(input)
	outputBuffer := bytes.NewBuffer(nil)
	err := fluentffmpeg.NewCommand("").
		PipeInput(inputBuff).
		OutputFormat("mp3").
		PipeOutput(outputBuffer).
		Run()
	if err != nil {
		log.WithError(err)
	}
	return outputBuffer.Bytes()
}
