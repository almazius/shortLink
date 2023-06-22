package repository

import (
	"links/internal/links"
	"links/pkg/db"
	"time"
)

func ExistLink(link string) (bool, *links.MyError) {
	res := false
	err := db.Connection.QueryRowx(`select exists(select * from links where link=$1)`, link).Scan(&res)
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
	err := db.Connection.QueryRowx(`select exists(select * from links where shortLink=$1)`, shortLink).Scan(&res)
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
	err := db.Connection.QueryRowx(`select shortlink from links where link=$1`, link).Scan(&short)
	if err != nil {
		db.Log.Print(err)

		return "", &links.MyError{Code: 500, Err: err}
	}

	return short, nil
}

func AddNote(link, shortLink string, time time.Time) (int64, *links.MyError) {
	var id int64
	err := db.Connection.QueryRowx(`INSERT INTO links (link, shortLink, creationTime) values ($1, $2, $3) returning id`, link, shortLink, time).Scan(&id)
	if err != nil {
		db.Log.Print(err)

		return -1, &links.MyError{Code: 500, Err: err}
	}
	return id, nil
}

func FindLink(shortLink string) (string, *links.MyError) {
	fullLink := ""
	err := db.Connection.QueryRowx(`select link from links where shortLink=$1`, shortLink).Scan(&fullLink)
	if err != nil {
		db.Log.Print(err)

		return "", &links.MyError{Code: 500, Err: err}
	}

	return fullLink, nil
}
