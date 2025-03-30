package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(5 * time.Hour)
	// decrement the commie score every 5 hours
	go func() {
		DecrementCommieScore()
		for {
			select {
			case <-ticker.C:
				DecrementCommieScore()
			}
		}
	}()
}

// DecrementCommieScore is a function that decrements the commie score of all users
func DecrementCommieScore() {
	log.Println("Decrementing commie values")
	users := rdb.HGetAll(ctx, "commie").Val()
	for user, score := range users {
		i, err := strconv.ParseInt(score, 10, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		if i > 0 {
			err = rdb.HIncrBy(ctx, "commie", user, -1).Err()
			if err != nil {
				log.Println(err)
			}
		}
	}
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

func UpdateCommieScore(userID string, score int) int {

	err := rdb.HIncrBy(ctx, "commie", userID, int64(score)).Err()
	if err != nil {
		log.Println(err)
		return 0
	}

	furryScore, err := rdb.HGet(ctx, "commie", userID).Int()
	if err != nil {
		log.Println(err)
		return 0
	}

	return furryScore
}

func GetAllFurryScores() map[string]string {
	furryScores, err := rdb.HGetAll(ctx, "furry").Result()
	if err != nil {
		log.Println(err)
		return map[string]string{}
	}

	return furryScores
}

func GetAllCommieScores() map[string]string {
	furryScores, err := rdb.HGetAll(ctx, "commie").Result()
	if err != nil {
		log.Println(err)
		return map[string]string{}
	}

	return furryScores
}
