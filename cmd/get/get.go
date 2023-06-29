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

	fmt.Println(rdb.Get("FOO"))
	fmt.Println(rdb.Get("FOO").Result())

	fmt.Println(rdb.Get("BAZ"))
	fmt.Println(rdb.Get("BAZ").Result())
}
