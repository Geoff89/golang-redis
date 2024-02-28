package main

import (
	"context"
	"fmt"
	"log"

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

	// Perform basic check to see if conection to redis is successful with PING:PONG
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println(status)
}
