package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	ExampleClient()
}

func ExampleClient() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"localhost:7001"},
		//Password: "root123", // no password set
		//DB:       0,  // use default DB
	})
	cmd := rdb.Ping(context.Background())
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
