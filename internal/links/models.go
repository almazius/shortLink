package links

import (
	"context"
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

type LinkService interface {
	PostLink(ctx context.Context, link string) (string, *MyError)
	GetLink(ctx context.Context, shortLink string) (string, *MyError)
}

type DatabaseService interface {
	ExistShortLink(ctx context.Context, shortLink string) (bool, *MyError)
	GetShortLink(ctx context.Context, link string) (string, *MyError)
	AddNote(ctx context.Context, link, shortLink string, time time.Time) (int64, *MyError)
	GetFullLink(ctx context.Context, shortLink string) (string, *MyError)
}

type CacheService interface {
	GetLinkOnCache(shortLink string) (string, error)
	SetLinkOnCache(link, shortLink string) error
}
