package redis

import (
	"github.com/redis/go-redis/v9"
	"links/config"
)

func InitRedisDB(c *config.Config) *redis.Client {
	database := redis.NewClient(&redis.Options{
		Network:               *c.Redis.Network,
		Addr:                  *c.Redis.Host + ":" + *c.Redis.Port,
		ClientName:            "",
		Dialer:                nil,
		OnConnect:             nil,
		Protocol:              0,
		Username:              *c.Redis.User,
		Password:              *c.Redis.Password,
		CredentialsProvider:   nil,
		DB:                    0,
		MaxRetries:            0,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		ContextTimeoutEnabled: false,
		PoolFIFO:              false,
		PoolSize:              0,
		PoolTimeout:           0,
		MinIdleConns:          0,
		MaxIdleConns:          0,
		ConnMaxIdleTime:       0,
		ConnMaxLifetime:       0,
		TLSConfig:             nil,
		Limiter:               nil,
	})
	return database
}
