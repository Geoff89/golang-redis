package main

import (
	"context"
	"fmt"

	"github.com/Geoff89/go_redis/internal/utility"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	// Ensure you have redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     utility.Address(),
		Password: utility.Password(),
		DB:       utility.Database(),
	})
	// Ensure database is closed gracefully
	defer rdb.Close()

	rdb.Set(ctx, "FOO", "BAR", 0)
	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO Not Found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	// Update "FOO" to be associated with 5
	rdb.Del(ctx, "FOO")
	result, err = rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO Not Found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

}
