package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"links/internal/links"
	"log"
	"os"
	"time"
)

type RedisCache struct {
	Client *redis.Client
	Log    *log.Logger
}

func NewClient(client *redis.Client) links.CacheService {
	return &RedisCache{
		Client: client,
		Log:    log.New(os.Stdout, "Redis ", log.LstdFlags|log.Lshortfile),
	}
}

func (r *RedisCache) GetLinkOnCache(shortLink string) (string, error) {
	val, err := r.Client.Get(context.Background(), shortLink).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCache) SetLinkOnCache(link, shortLink string) error {
	err := r.Client.Set(context.Background(), shortLink, link, 10*time.Hour)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
