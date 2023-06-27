// Demonstration of how to connect to a Redis server from Go
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

	// Perform basic diagnostic to check if the connection is working
	// Expected result > ping: PONG
	fmt.Println(rdb.Ping())
}
