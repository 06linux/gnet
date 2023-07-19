package gbase

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type baseRedis struct{}

var Redis = &baseRedis{}

func (baseRedis) Test() {
	fmt.Println("baseRedis xxx")

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	_ = rdb.Set(ctx, "key_with_ttl", "bar", time.Minute).Err()
	_ = rdb.Set(ctx, "key_without_ttl_1", "", 0).Err()
	_ = rdb.Set(ctx, "key_without_ttl_2", "", 0).Err()

	rdb.HSet(ctx, "myhash", "key1", "value1", "key2", "value2")

}
