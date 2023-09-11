package test

import (
	"github.com/Rascal0814/boot/cache"
	"github.com/go-redis/redis"
)

const PrefixUserSession = "user:"

type UserCache = cache.Cache[*UserSession]

func NewUserCache(rdb *redis.Client) (UserCache, error) {
	return cache.NewWithPrefix[*UserSession](rdb, PrefixUserSession)
}
