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
	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	// Update "FOO" to be associated with 5
	rdb.Set(ctx, "FOO", 5, 0)
	intResult, err := rdb.Get(ctx, "FOO").Int()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %d\n", intResult)
	}
}
