package database

import (
	"fmt"
	"job-portal-api/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ConnectionToRedis(cfg config.Config) *redis.Client {
	str := cfg.RedisDB.Db
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return &redis.Client{}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisDB.Addr,
		Password:  cfg.RedisDB.Password,
		DB:       num,                                     
	})
	return rdb
}
