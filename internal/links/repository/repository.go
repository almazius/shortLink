package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"links/internal/links"
	"links/pkg/db"
	"time"
)

// Выбор ключом короткой ссылки обоснован тем, что пользователи чаще
// будут переходить по короткой ссылке, чем создавать ее
func GetLinkOnCache(shortLink string) (string, error) {
	val, err := db.Redis.Get(context.Background(), shortLink).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

func SetLinkOnCache(link, shortLink string) error {
	err := db.Redis.Set(context.Background(), shortLink, link, 10*time.Hour)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func ExistLink(link string) (bool, *links.MyError) {
	res := false
	err := db.Postgres.QueryRowx(`select exists(select * from links where link=$1)`, link).Scan(&res)
	if err != nil {
		db.Log.Print(err)
		return false, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	return res, nil
}

func ExistShortLink(shortLink string) (bool, *links.MyError) {
	res := false
	err := db.Postgres.QueryRowx(`select exists(select * from links where shortLink=$1)`, shortLink).Scan(&res)
	if err != nil {
		db.Log.Print(err)
		return false, &links.MyError{
			Code: 500,
			Err:  err,
		}
	}
	return res, nil
}

func GetShortLink(link string) (string, *links.MyError) {
	short := ""
	err := db.Postgres.QueryRowx(`select shortlink from links where link=$1`, link).Scan(&short)
	if err != nil {
		db.Log.Print(err)

		return "", &links.MyError{Code: 500, Err: err}
	}

	return short, nil
}

func AddNote(link, shortLink string, time time.Time) (int64, *links.MyError) {
	var id int64
	err := db.Postgres.QueryRowx(`INSERT INTO links (link, shortLink, creationTime) values ($1, $2, $3) returning id`, link, shortLink, time).Scan(&id)
	if err != nil {
		db.Log.Print(err)

		return -1, &links.MyError{Code: 500, Err: err}
	}
	return id, nil
}

func FindLink(shortLink string) (string, *links.MyError) {
	fullLink := ""
	err := db.Postgres.QueryRowx(`select link from links where shortLink=$1`, shortLink).Scan(&fullLink)
	if err != nil {
		db.Log.Print(err)

		return "", &links.MyError{Code: 500, Err: err}
	}

	return fullLink, nil
}
