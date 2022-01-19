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
	links := &Links{}
	ids := &IDS{}

	if (&audd.AppleMusicResult{}) != songInfo.AppleMusic {
		links.AppleMusic = songInfo.AppleMusic.URL
	}

	if (&audd.DeezerResult{}) != songInfo.Deezer {
		links.Deezer = songInfo.Deezer.Link
		ids.Deezer = songInfo.Deezer.ID
	}

	if (&audd.SpotifyResult{}) != songInfo.Spotify {
		links.Spotify = songInfo.Spotify.ExternalUrls.Spotify
		ids.Spotify = songInfo.Spotify.ID
	}

	meta := Metadata{
		Title:  songInfo.Title,
		Album:  songInfo.Album,
		Artist: songInfo.Artist,
		Url:    songInfo.SongLink,
		Links:  *links,
		IDS:    *ids,
	}
	return meta
}

func uploadToService(soundBuff []byte) audd.RecognitionResult {
	api_token := os.Getenv("AUDD_API_TOKEN")
	if len(api_token) == 0 {
		log.Fatal("Missing AUDD_API_TOKEN from environment")
	}
	client := audd.NewClient(api_token)
	song, err := client.RecognizeByFile(bytes.NewReader(soundBuff), "apple_music,deezer,spotify", nil)
	if err != nil {
		log.WithError(err).Error("Error while recognizing the music")
	}
	return song
}

func downloadVideo(url string) (output []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.WithError(err).Error("Error downloading the video")
	}
	defer resp.Body.Close()

	outputBuffer := new(bytes.Buffer)
	_, err = io.Copy(outputBuffer, resp.Body)
	if err != nil {
		log.WithError(err).Error("Error downloading the video")
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
		log.WithError(err).Error("Error while extracting the audio track")
	}
	return outputBuffer.Bytes()
}
