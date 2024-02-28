package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Geoff89/go_redis/internal/utility"
	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name string `redis:"name"`
	Age  int    `redis:"name"`
}

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

	_, err := rdb.Set(ctx, "FOO", "BAR", 0).Result()
	if err != nil {
		fmt.Println("Failed to add FOO <> BAR key value pair")
		return
	}
	rdb.Set(ctx, "INT", 5, 0)
	rdb.Set(ctx, "FLOAT", 5.5, 0)
	rdb.Set(ctx, "EXPIRING", 15, 30*time.Minute)
	rdb.Set(ctx, "LIST", []string{"Hello"}, 0)

	rdb.HSet(ctx, "STRUCT", Person{"JOHN DOE", 15})

}
