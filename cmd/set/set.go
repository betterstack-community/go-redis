package main

import "github.com/go-redis/redis"

func main() {
	// Ensure that you have Redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()

	rdb.Set("FOO", "BAR", 0)
	rdb.Set("INT", 5, 0)
	rdb.Set("FLOAT", 5.5, 0)
}
