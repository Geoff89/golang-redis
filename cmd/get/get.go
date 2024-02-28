package main

import (
	"context"
	"fmt"

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

	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("Key FOO not found in Redis Cache")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}

	intValue, err := rdb.Get(ctx, "INT").Int()
	if err != nil {
		fmt.Println("Key INT not found in Redis Cache")
	} else {
		fmt.Printf("INT has value %d\n", intValue)
	}

	var person Person
	err = rdb.HGetAll(ctx, "STRUCT").Scan(&person)
	if err != nil {
		fmt.Println("Key STRUCT not found in Redis Cache")
	} else {
		fmt.Printf("FOO has value %+v\n", person)
	}

	result, err = rdb.Get(ctx, "BAZ").Result()
	if err != nil {
		fmt.Println("Key BAZ not found in Redis Cache")
	} else {
		fmt.Printf("BAZ has value %s\n", result)
	}

	result, err = rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("Key FOO not found in Redis Cache")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}
}
