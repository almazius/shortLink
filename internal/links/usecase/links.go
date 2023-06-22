package usecase

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"links/internal/links"
	"links/internal/links/repository"
	"links/pkg/db"
	"math/big"
	"time"
)

const salt = "4uj4fj4thj"
const lengthLink = 6
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const prefix = "link.ru/"

// храним ид записи
// транформируем его в алфавит мощностью 62(a-zA-Z0-9)
// туда-сюда и мы богаты
// ;(
// храним время создания ссылки
// допустим, бд, позволяющая хранить 56800235584 записей, позволит нормально существовать
// короткой ссылке 10 часов, затем, она будет удаляться отдельной утилитой
// в случае, если записи будут превышать максимум для 6 символов, временно будет увеличен максимум

//в последовательность впихнуть значение, которое было удалено!

// создает сокращенную ссылку по id
func PostLink(link string) (string, *links.MyError) {
	exist, err := repository.ExistLink(link)
	if err != nil {
		return "", err
	}
	if exist {
		return repository.GetShortLink(link)
	} else {
		tempLink := link
		shortLink := ""
		for {
			shortLink = convertHashToLink(tempLink)

			exist, err = repository.ExistShortLink(shortLink)
			if err != nil {
				return "", err
			}
			if !exist {
				break
			}
			tempLink += salt
		}

		redisErr := repository.SetLinkOnCache(link, prefix+shortLink)
		if redisErr != nil {
			db.Log.Print(redisErr)
		}

		_, err = repository.AddNote(link, prefix+shortLink, time.Now())
		if err != nil {
			return "", err
		}

		return prefix + shortLink, nil
	}
}

func GetLink(shortLink string) (string, *links.MyError) {
	link, redisErr := repository.GetLinkOnCache(shortLink)
	if redisErr == nil && link != "" {
		return link, nil
	} else if redisErr != nil {
		db.Log.Print(redisErr)
	}

	link, err := repository.FindLink(shortLink)
	if err != nil && err.Err == sql.ErrNoRows {
		err.Code = 404
		err.Err = errors.New("link not found")
		return "", err
	} else if err != nil {
		return "", err
	}
	redisErr = repository.SetLinkOnCache(link, shortLink)
	if redisErr != nil {
		db.Log.Print(redisErr)
	}
	return link, nil
}

func convertHashToLink(link string) string {
	shortLink := make([]byte, lengthLink, lengthLink)
	v := big.Int{}
	h := sha256.New()
	h.Write([]byte(link))
	hash := h.Sum(nil)
	result := v.SetBytes([]byte(hex.EncodeToString(hash)))
	bigInt := result.Uint64()
	for i := 0; i < lengthLink; i++ {
		shortLink[i] = alphabet[bigInt%62]
		bigInt /= 62
	}
	//result.Mod(result, big.NewInt(lengthLink))
	return string(shortLink)
}
