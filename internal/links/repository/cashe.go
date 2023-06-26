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

// NewClient Create new instance RedisCache
func NewClient(client *redis.Client) links.CacheService {
	return &RedisCache{
		Client: client,
		Log:    log.New(os.Stdout, "Redis ", log.LstdFlags|log.Lshortfile),
	}
}

// GetLinkOnCache gets the full link by the shortened version from the cache
func (r *RedisCache) GetLinkOnCache(shortLink string) (string, error) {
	val, err := r.Client.Get(context.Background(), shortLink).Result()
	if err != redis.Nil {
		return "", err
	}
	return val, nil
}

// SetLinkOnCache Adds the entry "full link - abbreviated link" to the cache
func (r *RedisCache) SetLinkOnCache(link, shortLink string) error {
	err := r.Client.Set(context.Background(), shortLink, link, 10*time.Hour)
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}
