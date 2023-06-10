package repository

import (
	"links/pkg/db"
	"time"
)

func ExistLink(link string) (bool, error) {
	res := false
	err := db.Connection.QueryRowx(`select exists(select * from links where link=$1)`, link).Scan(&res)
	if err != nil {
		db.Log.Print(err)
		return false, err
	}
	return res, nil
}

func GetShortLink(link string) (string, error) {
	short := ""
	err := db.Connection.QueryRowx(`select shortLink form links where link=$1`, link).Scan(&short)
	if err != nil {
		db.Log.Print(err)

		return "", err
	}

	return short, err
}

func AddNote(link string, time time.Time) (int64, error) {
	var id int64
	err := db.Connection.QueryRowx(`INSERT INTO links (link, creationTime) values ($1, $2) returning id`, link, time).Scan(&id)
	if err != nil {
		db.Log.Print(err)

		return -1, err
	}
	return id, nil
}

func AddShortLink(id int64, link string) error {
	_, err := db.Connection.Exec(`update links set shortLink=$2 where id=$1`, id, link)
	if err != nil {
		db.Log.Print(err)

		return err
	}
	return nil
}

func FindLink(shortLink string) (string, error) {
	fullLink := ""
	err := db.Connection.QueryRowx(`select link from links where shortLink=$1`, shortLink).Scan(&fullLink)
	if err != nil {
		db.Log.Print(err)

		return "", err
	}

	return fullLink, nil
}
