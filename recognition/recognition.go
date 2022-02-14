package recognition

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AudDMusic/audd-go"
	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	log "github.com/sirupsen/logrus"
)

func DetectMusic(url string) (*Metadata, error) {
	songInfo, err := analyzeVideo(url)
	if err != nil {
		return nil, err
	}

	if len(songInfo.Title) > 0 {
		m := Metadata{
			Title:  songInfo.Title,
			Album:  songInfo.Album,
			Artist: songInfo.Artist,
			Url:    songInfo.SongLink,
		}

		if songInfo.AppleMusic != nil {
			m.Links.AppleMusic = songInfo.AppleMusic.URL
		}

		if songInfo.Deezer != nil {
			m.Links.Deezer = songInfo.Deezer.Link
			m.IDS.Deezer = songInfo.Deezer.ID
		}

		if songInfo.Spotify != nil {
			m.Links.Spotify = songInfo.Spotify.ExternalUrls.Spotify
			m.IDS.Spotify = songInfo.Spotify.ID
		}

		return &m, nil
	}

	return nil, nil
}

var apiToken string
var client *audd.Client

func init() {
	apiToken = os.Getenv("AUDD_API_TOKEN")
	if len(apiToken) == 0 {
		log.Fatal("Missing AUDD_API_TOKEN from environment")
	}
	client = audd.NewClient(apiToken)
}

func analyzeVideo(url string) (*audd.RecognitionResult, error) {
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
			log.WithError(err).Error("Error while extracting the audio track")
		}
	}()

	song, err := client.RecognizeByFile(reader, "apple_music,deezer,spotify", nil)
	if err != nil {
		return nil, fmt.Errorf("recognizing music: %v", err)
	}

	return &song, nil
}
