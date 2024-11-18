package redis

import (
	"github.com/gofiber/storage/redis/v3"
	"github.com/salawatbro/raxmet/config"
	"runtime"
	"sync"
	"time"
)

var (
	once         sync.Once
	redisStorage *redis.Storage
)

func SetUpRedisClient() {
	once.Do(func() {
		redisStorage = redis.New(redis.Config{
			Host:      config.Cfg.Redis.Host,
			Port:      config.Cfg.Redis.Port,
			Username:  config.Cfg.Redis.User,
			Password:  config.Cfg.Redis.Password,
			URL:       "",
			Database:  0,
			Reset:     false,
			TLSConfig: nil,
			PoolSize:  10 * runtime.GOMAXPROCS(0),
		})
	})
}

func SetRedisKey(key string, value []byte, epx time.Duration) error {
	return redisStorage.Set(key, value, epx)
}

func GetRedisKey(key string) ([]byte, error) {
	return redisStorage.Get(key)
}

func DeleteRedisKey(key string) error {
	return redisStorage.Delete(key)
}
