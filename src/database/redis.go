package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetSettings() map[string]string {

	settings, err := rdb.HGetAll(ctx, "settings").Result()
	if err != nil {
		log.Println(err)
		return map[string]string{}
	}

	return settings
}

func UpdateSettings(key string, value string) {

	err := rdb.HSet(ctx, "settings", key, value).Err()
	if err != nil {
		log.Println(err)
	}
}

// UpdateFurryScore is a function that updates the furry score of a user and returns the new score
func UpdateFurryScore(userID string, score int) int {

	err := rdb.HIncrBy(ctx, "furry", userID, int64(score)).Err()
	if err != nil {
		log.Println(err)
		return 0
	}

	furryScore, err := rdb.HGet(ctx, "furry", userID).Int()
	if err != nil {
		log.Println(err)
		return 0
	}

	return furryScore
}
