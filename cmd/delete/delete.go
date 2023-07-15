package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/woojiahao/go_redis/internal/utility"
)

func main() {
	ctx := context.Background()
	// Ensure that you have Redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     utility.Address(),
		Password: utility.Password(), // no password set
		DB:       utility.Database(), // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()

	// Set "FOO" to be associated with "BAR"
	rdb.Set(ctx, "FOO", "BAR", 0)
	fmt.Println(rdb.Get(ctx, "FOO").Result())
	// Deleting the key "FOO" and its associated value
	rdb.Del(ctx, "FOO")
	fmt.Println(rdb.Get(ctx, "FOO").Result())
}
