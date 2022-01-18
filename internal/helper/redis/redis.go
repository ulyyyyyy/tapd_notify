package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ulyyyyyy/tapd_notify/configs"
	"time"
)

const Nil = redis.Nil

var (
	client         *redis.Client
	operateTimeout time.Duration
)

func Initialize() error {
	operateTimeout = configs.DB.Timeout.Operate
	client = redis.NewClient(&redis.Options{
		Addr:     configs.DB.Redis.Addr,
		Password: configs.DB.Redis.Password,
		DB:       configs.DB.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return fmt.Errorf("redis connection fail: %s", pong)
	}

	return nil
}

func Keys(pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.Keys(ctx, pattern).Result()
}

func Set(key string, value interface{}, expiration ...time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	if len(expiration) == 0 {
		expiration = []time.Duration{0}
	}

	return client.Set(ctx, key, value, expiration[0]).Err()
}

func Exists(key string, keys ...string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	if len(keys) == 0 {
		keys = []string{key}
	} else {
		keys = append(keys, key)
	}

	result, err := client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}
	if result == 0 {
		return false, nil
	}
	if result != int64(len(keys)) {
		return false, fmt.Errorf("expected count '%d': actual count '%d'", len(keys), result)
	}
	return true, nil
}

func get(key string) *redis.StringCmd {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.Get(ctx, key)
}

func Get(key string) (result string, err error) {
	defer func() {
		result, err = get(key).Result()
	}()
	return
}

func Get2Bytes(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.Get(ctx, key).Bytes()
}

func RPush(key string, value ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.RPush(ctx, key, value).Err()
}

func ListGetAll(key string) ([]string, error) {
	return ListGetRange(key, 0, -1)
}

func LRem(key string, count int64, value string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()
	return client.LRem(ctx, key, count, value).Result()
}

func ListGetRange(key string, start, end int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.LRange(ctx, key, start, end).Result()
}

func Del(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.Del(ctx, key).Err()
}

func HashSet(key string, values ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.HSet(ctx, key, values...).Err()
}

func hGetAll(key string) *redis.StringStringMapCmd {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.HGetAll(ctx, key)
}

func HashGetAll(key string) *redis.StringStringMapCmd {
	return hGetAll(key)
}

func HashGetAllToMap(key string) map[string]string {
	return hGetAll(key).Val()
}

func hGet(key string, field string) *redis.StringCmd {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.HGet(ctx, key, field)
}

func HashGet(key string, field string) (string, error) {
	return hGet(key, field).Result()
}

func HashGet2Bool(key string, field string) (bool, error) {
	return hGet(key, field).Bool()
}

func HashGet2Int(key string, field string) (int, error) {
	return hGet(key, field).Int()
}

func HashGet2Time(key string, field string) (time.Time, error) {
	return hGet(key, field).Time()
}

// func HashGetAll(key string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
// 	defer cancel()
//
// 	client.HGetAll(ctx, key)
// }

func Expire(key string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), operateTimeout)
	defer cancel()

	return client.Expire(ctx, key, expiration).Err()
}
