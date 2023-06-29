package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	fmt.Println("Without caching...")
	start := time.Now()
	getDataExpensive()
	elapsed := time.Since(start)
	fmt.Printf("Without caching took %s\n", elapsed)

	fmt.Println("With caching...")
	start = time.Now()
	getDataCached()
	elapsed = time.Since(start)
	fmt.Printf("With caching took %s\n", elapsed)
}

func getDataExpensive() {
	for i := 0; i < 3; i++ {
		fmt.Println("\tBefore query")
		result := databaseQuery()
		fmt.Printf("\tAfter query with result %s\n", result)
	}
}

func getDataCached() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()

	for i := 0; i < 3; i++ {
		fmt.Println("\tBefore query")
		val, err := rdb.Get("query").Result()
		if err != nil {
			// Database query was not cached yet
			// Make database call and cache the value
			val = databaseQuery()
			rdb.Set("query", val, 0)
		}
		fmt.Printf("\tAfter query with result %s\n", val)
	}
}

func databaseQuery() string {
	fmt.Println("\tDatabase queried")
	// Intentionally sleep for 5 seconds to simulate a long database query
	time.Sleep(5 * time.Second)
	return "bar"
}
