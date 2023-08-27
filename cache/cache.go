package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type CacheObject struct {
	Key   string
	Value string
}

var ctx = context.Background()

func GetClient() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB

	})

	return client
}

func SaveToken(user string, token string) (string, error) {
	client := GetClient()

	_, err := client.Set(ctx, user, token, 0).Result()
	if err != nil {
		log.Fatal(err)
	}
	return client.Get(ctx, user).Result()

}

func CmpToken(user string, token string) bool {

	client := GetClient()

	cachedToken, err := client.Get(ctx, user).Result()
	if err == redis.Nil {
		client.Set(ctx, user, token, 0)
	}

	return cachedToken == token
}

func TokenExists(user string, token string) bool {
	client := GetClient()

	cachedToken, err := client.Get(ctx, user).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	return cachedToken != ""
}
