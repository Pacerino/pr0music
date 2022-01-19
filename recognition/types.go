package recognition

type Metadata struct {
	// Title from the recognized song
	Title string
	// Album from the recognized song
	Album string
	// Artist from the recognized song
	Artist string
	// URL to Audd.io from the recognized song
	Url string
	// Respective links to known streaming platforms
	Links Links
	// Respective IDs to known streaming platforms
	IDS IDS
}

type Links struct {
	// Deezer link
	Deezer string
	// Spotify link
	Spotify string
	// Apple Music link
	AppleMusic string
}

type IDS struct {
	// Deezer ID
	Deezer int
	// Spotify ID
	Spotify string
}
