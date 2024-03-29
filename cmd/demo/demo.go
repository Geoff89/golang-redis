package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Geoff89/go_redis/internal/utility"
	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Without caching...")
	start := time.Now()
	getDataExpensive()
	elapsed := time.Since(start)
	fmt.Printf("Without caching took %s\n\n", elapsed)

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
	ctx := context.Background()
	// Ensure you have redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     utility.Address(),
		Password: utility.Password(),
		DB:       utility.Database(),
	})
	// Ensure database is closed gracefully
	defer rdb.Close()

	for i := 0; i < 3; i++ {
		fmt.Println("\tBefore query")
		val, err := rdb.Get(ctx, "query").Result()
		if err != nil {
			// Database query was not cached yet
			// Make database call and cache the value
			val = databaseQuery()
			rdb.Set(ctx, "query", val, 0)
		}
		fmt.Printf("\tAfter query with result %s\n", val)
	}
}

func databaseQuery() string {
	fmt.Println("\tDatabase queried")
	// Intentionally sleep for 5 seconds
	time.Sleep(5 * time.Second)
	return "bar"
}
