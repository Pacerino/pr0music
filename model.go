package main

import (
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ItemID   int `gorm:"primarykey"`
	Title    string
	Album    string
	Artist   string
	Url      string
	Metadata ItemMetadata `gorm:"embedded"`
}

type ItemMetadata struct {
	DeezerURL     string
	DeezerID      string
	SoundcloudURL string
	SoundcloudID  string
	SpotifyURL    string
	SpotifyID     string
	YoutubeURL    string
	YoutubeID     string
	TidalURL      string
	TidalID       string
	ApplemusicURL string
	ApplemusicID  string
}

type Comments struct {
	gorm.Model
	CommentID int `gorm:"primarykey"`
	Up        int
	Down      int
	Content   string
	Created   *time.Time
	ItemID    int `gorm:"not null;index;unique"`
	Thumb     string
}
