package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/jackc/pgx/v5"
	"links/config"
	"links/internal/links"
	"links/internal/links/repository"
	"links/pkg/posgresql"
	"links/pkg/redis"
	"log"
	"math/big"
	"os"
	"time"
)

const salt = "4uj4fj4thj"
const lengthLink = 6
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const prefix = "link.ru/"

type Service struct {
	Log   *log.Logger
	Db    links.DatabaseService
	Cache links.CacheService
}

// NewLinkService Create new instance Postgres
func NewLinkService(ctx context.Context, conf *config.Config) (links.LinkService, error) {
	Log := log.New(os.Stdout, "LinkService ", log.Lshortfile|log.LstdFlags)
	pool, err := posgresql.GetPool(ctx, conf)
	if err != nil {
		Log.Print(err)
		return nil, err
	}
	database := repository.NewPostgres(pool)

	client := redis.InitRedisDB(conf)
	_, err = client.Ping(ctx).Result()
	if err != nil {
		Log.Print(err)
		return nil, err
	}
	cache := repository.NewClient(client)

	return &Service{
		Log:   Log,
		Db:    database,
		Cache: cache,
	}, nil
}

// PostLink Checks for the presence of a link in the database, if it does not exist,
// then creates a shortened version. If it already exists, it will return an abbreviated version of it
func (s *Service) PostLink(ctx context.Context, link string) (string, *links.MyError) {
	if shortLink, err := s.Db.GetShortLink(ctx, link); err == nil {
		return shortLink, nil
	} else {
		tempLink := link
		for {
			shortLink = convertHashToLink(tempLink)

			exist, err := s.Db.ExistShortLink(ctx, shortLink)
			if err != nil {
				return "", err
			}
			if !exist {
				break
			}
			tempLink += salt
		}
		redisErr := s.Cache.SetLinkOnCache(link, prefix+shortLink)
		if redisErr != nil {
			s.Log.Print(redisErr)
		}
		_, err = s.Db.AddNote(ctx, link, prefix+shortLink, time.Now())
		if err != nil {
			return "", err
		}

		return prefix + shortLink, nil
	}
}

// GetLink Checks for an abbreviated link in the cache,
// if it is empty or an error has occurred, then tries to get the full version from the database.
func (s *Service) GetLink(ctx context.Context, shortLink string) (string, *links.MyError) {
	link, redisErr := s.Cache.GetLinkOnCache(shortLink)
	if redisErr == nil && link != "" {
		return link, nil
	} else if redisErr != nil {
		s.Log.Print(redisErr)
	}

	link, err := s.Db.GetFullLink(ctx, shortLink)
	if err != nil && err.Err == pgx.ErrNoRows {
		err.Code = 404
		err.Err = errors.New("link not found")
		return "", err
	} else if err != nil {
		return "", err
	}
	redisErr = s.Cache.SetLinkOnCache(link, shortLink)
	if redisErr != nil {
		s.Log.Print(redisErr)
	}
	return link, nil
}

// convertHashToLink converts the full link to a cache, and then to an abbreviated link
// by dividing the cache by 62 and getting the number of the number that
// is replaced by a character in the alphabet.
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
