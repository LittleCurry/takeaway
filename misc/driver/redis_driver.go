package driver

import (
	"gopkg.in/redis.v5"
	"fmt"
)

const REDIS_KEY_PREFIX = "AccessToken"
//const REDIS_CODE_PREFIX = "VerifyCode"
const REDIS_REQUEST_PER_DAY_OF_DEVICE = "Device"

var (
	client *redis.Client
	RedisClient *redis.Client
)

func RedisInit(redisAddr string, redisDb int) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   redisDb,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("RedisInit_err:", err)
	}
}

func Redis() *redis.Client {
	return client
}
