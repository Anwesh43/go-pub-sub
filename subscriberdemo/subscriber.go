package main

import (
	"context"
	"fmt"
	"subscriberdemo/constants"

	"github.com/go-redis/redis/v8"
)

func subscribe(rdb *redis.Client, ctx context.Context, ch chan int) {
	pubsub := rdb.Subscribe(ctx, constants.MY_CHANNEL)
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err == nil {
			fmt.Println(msg.Payload)
		}
		if msg != nil && string(msg.Payload) == "quit" {
			break
		}
	}
	ch <- 0

}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ch := make(chan int)
	go subscribe(rdb, ctx, ch)
	<-ch
}
