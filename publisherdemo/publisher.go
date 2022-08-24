package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"publisherdemo/constants"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		data := scanner.Text()
		err := rdb.Publish(ctx, constants.MY_CHANNEL, data).Err()
		if err == nil {
			fmt.Println("published")
		} else {
			panic(err)
		}
		if data == "quit" {
			break
		}
	}
}
