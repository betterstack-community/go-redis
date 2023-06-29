package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// Ensure that you have Redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()

	// Set "FOO" to be associated with "BAR"
	rdb.Set("FOO", "BAR", 0)
	fmt.Println(rdb.Get("FOO").Result())
	// Update "FOO" to be associated with 5
	rdb.Set("FOO", 5, 0)
	fmt.Println(rdb.Get("FOO").Result())
}
