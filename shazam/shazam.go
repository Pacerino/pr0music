package shazam

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	fluentffmpeg "github.com/modfy/fluent-ffmpeg"
	"github.com/sirupsen/logrus"
)

type HubActions struct {
	Name string
	Type string
	ID   string
	URI  string
}

type ShazamResponse struct {
	Matches []struct {
		ID            string  `json:"id"`
		Offset        float64 `json:"offset"`
		Timeskew      float64 `json:"timeskew"`
		Frequencyskew float64 `json:"frequencyskew"`
	} `json:"matches"`
	Location struct {
		Accuracy float64 `json:"accuracy"`
	} `json:"location"`
	Timestamp int64  `json:"timestamp"`
	Timezone  string `json:"timezone"`
	Track     struct {
		Layout   string `json:"layout"`
		Type     string `json:"type"`
		Key      string `json:"key"`
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Images   struct {
			Background string `json:"background"`
			Coverart   string `json:"coverart"`
			Coverarthq string `json:"coverarthq"`
			Joecolor   string `json:"joecolor"`
		} `json:"images"`
		Share struct {
			Subject  string `json:"subject"`
			Text     string `json:"text"`
			Href     string `json:"href"`
			Image    string `json:"image"`
			Twitter  string `json:"twitter"`
			HTML     string `json:"html"`
			Avatar   string `json:"avatar"`
			Snapchat string `json:"snapchat"`
		} `json:"share"`
		Hub struct {
			Type    string `json:"type"`
			Image   string `json:"image"`
			Actions []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				ID   string `json:"id,omitempty"`
				URI  string `json:"uri,omitempty"`
			} `json:"actions"`
			Options []struct {
				Caption string `json:"caption"`
				Actions []struct {
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"actions"`
				Beacondata struct {
					Type         string `json:"type"`
					Providername string `json:"providername"`
				} `json:"beacondata"`
				Image               string `json:"image"`
				Type                string `json:"type"`
				Listcaption         string `json:"listcaption"`
				Overflowimage       string `json:"overflowimage"`
				Colouroverflowimage bool   `json:"colouroverflowimage"`
				Providername        string `json:"providername"`
			} `json:"options"`
			Providers []struct {
				Caption string `json:"caption"`
				Images  struct {
					Overflow string `json:"overflow"`
					Default  string `json:"default"`
				} `json:"images"`
				Actions []struct {
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"actions"`
				Type string `json:"type"`
			} `json:"providers"`
			Explicit    bool   `json:"explicit"`
			Displayname string `json:"displayname"`
		} `json:"hub"`
		Sections []struct {
			Type      string `json:"type"`
			Metapages []struct {
				Image   string `json:"image"`
				Caption string `json:"caption"`
			} `json:"metapages,omitempty"`
			Tabname  string `json:"tabname"`
			Metadata []struct {
				Title string `json:"title"`
				Text  string `json:"text"`
			} `json:"metadata,omitempty"`
			Youtubeurl string `json:"youtubeurl,omitempty"`
			URL        string `json:"url,omitempty"`
		} `json:"sections"`
		URL     string `json:"url"`
		Artists []struct {
			ID     string `json:"id"`
			Adamid string `json:"adamid"`
		} `json:"artists"`
		Isrc   string `json:"isrc"`
		Genres struct {
			Primary string `json:"primary"`
		} `json:"genres"`
		Urlparams struct {
			Tracktitle  string `json:"{tracktitle}"`
			Trackartist string `json:"{trackartist}"`
		} `json:"urlparams"`
		Myshazam struct {
			Apple struct {
				Actions []struct {
					Name string `json:"name"`
					Type string `json:"type"`
					URI  string `json:"uri"`
				} `json:"actions"`
			} `json:"apple"`
		} `json:"myshazam"`
		Highlightsurls struct {
			Artisthighlightsurl string `json:"artisthighlightsurl"`
		} `json:"highlightsurls"`
		Relatedtracksurl string `json:"relatedtracksurl"`
		Albumadamid      string `json:"albumadamid"`
	} `json:"track"`
	Tagid string `json:"tagid"`
}

func getAudioFromUrl(url string) (*io.PipeReader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("downloading video: %v", err)
	}

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
			fmt.Println(err)
			logrus.WithError(err).Error("Error while extracting the audio track")
		}
		writer.Close()
		resp.Body.Close()
	}()

	return reader, nil
}

func Recognize(audioUrl string) (ShazamResponse, error) {
	reader, err := getAudioFromUrl(audioUrl)
	if err != nil {
		return ShazamResponse{}, err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "audio.mp3")
	if err != nil {
		return ShazamResponse{}, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return ShazamResponse{}, err
	}
	err = writer.Close()
	if err != nil {
		return ShazamResponse{}, err
	}

	req, err := http.NewRequest("POST", os.Getenv("SHAZAM_ENDPOINT"), body)
	if err != nil {
		return ShazamResponse{}, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ShazamResponse{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShazamResponse{}, err
	}
	var shazamResponse ShazamResponse
	if err = json.Unmarshal(respBody, &shazamResponse); err != nil {
		return ShazamResponse{}, err
	}
	return shazamResponse, nil
}

func GetAppleMusicLink(shazamInfo ShazamResponse) (string, error) {
	for _, opt := range shazamInfo.Track.Hub.Options {
		if opt.Caption == "OPEN IN" {
			if len(opt.Actions) > 0 {
				for _, act := range opt.Actions {
					if act.Type == "applemusicopen" {
						return act.URI, nil
					}
				}
			}
		}
	}
	return "", errors.New("no apple music link found")
}

func GetAlbum(shazamInfo ShazamResponse) (string, error) {
	for _, sec := range shazamInfo.Track.Sections {
		if sec.Type == "SONG" {
			for _, meta := range sec.Metadata {
				if meta.Title == "Album" {
					return meta.Text, nil
				}
			}
		}
	}
	return "", errors.New("no album found")
}
