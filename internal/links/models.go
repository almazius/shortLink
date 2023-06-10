package links

import "time"

type Link struct {
	Id           int64     `json:"id"`
	FullLink     string    `json:"link"`
	ShortLink    string    `json:"shortLink"`
	CreationTime time.Time `json:"creationTime"`
}
