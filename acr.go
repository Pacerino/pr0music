package main

type ACRRecognitionResult struct {
	Metadata struct {
		Music []struct {
			ExternalIds struct {
			} `json:"external_ids"`
			ExternalMetadata struct {
				Youtube struct {
					Vid string `json:"vid"`
				} `json:"youtube"`
				Deezer struct {
					Album struct {
						Name string `json:"name"`
					} `json:"album"`
					Track struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"track"`
					Artists []struct {
						Name string `json:"name"`
					} `json:"artists"`
				} `json:"deezer"`
				Spotify struct {
					Album struct {
						Name string `json:"name"`
					} `json:"album"`
					Track struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					} `json:"track"`
					Artists []struct {
						Name string `json:"name"`
					} `json:"artists"`
				} `json:"spotify"`
			} `json:"external_metadata"`
			ResultFrom int `json:"result_from"`
			Album      struct {
				Name string `json:"name"`
			} `json:"album"`
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Genres []struct {
				Name string `json:"name"`
			} `json:"genres"`
			DurationMs   int    `json:"duration_ms"`
			Score        int    `json:"score"`
			Title        string `json:"title"`
			PlayOffsetMs int    `json:"play_offset_ms"`
			ReleaseDate  string `json:"release_date"`
			Label        string `json:"label"`
			Acrid        string `json:"acrid"`
		} `json:"music"`
		TimestampUtc string `json:"timestamp_utc"`
	} `json:"metadata"`
	ResultType int `json:"result_type"`
	Status     struct {
		Version string `json:"version"`
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
	} `json:"status"`
	CostTime float64 `json:"cost_time"`
}

type RecognizedMetadata struct {
	// Title from the recognized song
	Title string
	// Album from the recognized song
	Album string
	// Artist from the recognized song
	Artist string
	// Respective ACR ID from the recognized song
	AcrID string
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
	YouTube string
}

type MetadataIDs struct {
	// Deezer ID
	Deezer string
	// Spotify ID
	Spotify string
	// YouTube ID
	YouTube string
}
