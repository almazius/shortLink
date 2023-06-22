package links

import (
	"fmt"
	"time"
)

type Link struct {
	Id           int64     `json:"id"`
	FullLink     string    `json:"link"`
	ShortLink    string    `json:"shortLink"`
	CreationTime time.Time `json:"creationTime"`
}

type MyError struct {
	Code int32
	Err  error
}

func (err *MyError) Error() string {
	return fmt.Sprintf("Status: %d\nError: %s", err.Code, err.Err)
}
