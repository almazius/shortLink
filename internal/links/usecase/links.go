package usecase

import (
	"errors"
	"links/internal/links/repository"
	"time"
)

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
func POST(link string) (string, error) {

	// 		POST
	//find link
	// if finded {
	//	return shortLink
	//} else {
	//	add link, return  id
	// create short link
	//  return shortLink
	// }
	//
	isExist, err := repository.ExistLink("https://github.com/jackc/pgx")
	if err != nil {
		return "", err
	}
	if isExist {
		return repository.GetShortLink(link)
	} else {
		id, err := repository.AddNote(link, time.Now())
		if err != nil {
			return "", err
		}
		shortLink, err := CreateShortLink(id)
		if err != nil {
			return "", err
		}

		err = repository.AddShortLink(id, shortLink)
		if err != nil {
			return "", err
		}

		return shortLink, nil
	}
}

func CreateShortLink(id int64) (string, error) {
	res := ""
	shortLinkOnNumbers := make([]int64, 0)

	// convert into 62-znachnyy system
	for id > 0 {
		shortLinkOnNumbers = append(shortLinkOnNumbers, id%62)
		id /= 62
	}

	// 0=a 25=z 26=A 51=Z 52=0 61=9
	for _, el := range shortLinkOnNumbers {
		if el >= 0 && el <= 25 {
			res += string(rune('a' + el))
		} else if el >= 26 && el <= 51 {
			res += string(rune('A' + el - 26))
		} else {
			res += string(rune('0' + el - 52))
		}
	}
	if len(res) < 6 {
		for i := len(res); i < 6; i++ {
			res = "a" + res
		}
	} else if len(res) > 6 {
		return "", errors.New("big id")
	}
	res = "link.ru/" + res
	return res, nil
}
