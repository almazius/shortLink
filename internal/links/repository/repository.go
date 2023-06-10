package repository

import (
	"links/pkg/db"
	"time"
)

func ExistLink(link string) (bool, error) {
	res := false
	err := db.Connection.QueryRowx(`select exists(select * from links where link=$1)`, link).Scan(&res)
	if err != nil {
		return false, err
	}
	return res, nil
}

func GetShortLink(link string) (string, error) {
	short := ""
	err := db.Connection.QueryRowx(`select shortlink form links where link=$1`, link).Scan(&short)
	if err != nil {
		return "", err
	}

	return short, err
}

func AddNote(link string, time time.Time) (int64, error) {
	var id int64
	err := db.Connection.QueryRowx(`INSERT INTO links (link, creationtime) values ($1, $2) returning id`, link, time).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AddShortLink(id int64, link string) error {
	_, err := db.Connection.Exec(`update links set shortlink=$2 where id=$1`, id, link)
	if err != nil {
		return err
	}
	return nil
}
