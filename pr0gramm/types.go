package pr0gramm

import "time"

type Id uint64

type LoginResponse struct {
	Success    bool   `json:"success"`
	Identifier string `json:"identifier"`

	Ban *BanInfo `json:"ban,omitempty"`
}

type BanInfo struct {
	Banned  bool      `json:"banned"`
	Reason  string    `json:"reason"`
	EndTime Timestamp `json:"till"`
}

type Response struct {
	Timestamp    Timestamp     `json:"ts"`
	ResponseTime time.Duration `json:"rt"`
	QueryCount   uint          `json:"qt"`
}

type CommentResponse struct {
	Comments []Comment `json:"comments"`
	HasOlder bool      `json:"hasOlder"`
	HasNewer bool      `json:"hasNewer"`
	Response
}

type Comment struct {
	Id        Id        `json:"id"`
	Created   Timestamp `json:"created"`
	Up        int       `json:"up"`
	Down      int       `json:"down"`
	Content   string    `json:"content"`
	Thumbnail string    `json:"thumb"`
	ItemId    Id        `json:"itemId"`
}

type Message struct {
	Type       string      `json:"type"`
	ID         int         `json:"id"`
	ItemID     int         `json:"itemId"`
	Image      interface{} `json:"image"`
	Thumb      string      `json:"thumb"`
	Flags      int         `json:"flags"`
	Name       string      `json:"name"`
	Mark       int         `json:"mark"`
	SenderID   int         `json:"senderId"`
	Score      int         `json:"score"`
	Collection interface{} `json:"collection"`
	Owner      interface{} `json:"owner"`
	OwnerMark  interface{} `json:"ownerMark"`
	Keyword    interface{} `json:"keyword"`
	Created    int         `json:"created"`
	Message    string      `json:"message"`
	Read       int         `json:"read"`
	Blocked    int         `json:"blocked"`
}

type MessagesResponse struct {
	Messages []Message `json:"messages"`
}
