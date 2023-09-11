package test_test

import (
	"github.com/Rascal0814/boot/cache/test"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestNewUserCache(t *testing.T) {
	cache, err := test.NewUserCache(redisFun())
	if err != nil {
		t.Fatal(err)
	}
	err = cache.Put("1", &test.UserSession{Id: 1}, time.Minute*30)
	if err != nil {
		t.Fatal(err)
	}
}

func redisFun() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		DialTimeout: time.Second * 2,
		PoolSize:    10,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}
