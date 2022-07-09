// Package redis Package config
// @author    : MuXiang123
// @time      : 2022/6/21 22:25
package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbLikeUserId *redis.Client
var RdbLikeVideoId *redis.Client

func InitRedis() {
	RdbLikeUserId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "tiktok",
		DB:       0, //  选择将点赞视频id信息存入 DB0
	})

	RdbLikeVideoId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "tiktok",
		DB:       1, //  选择将点赞用户id信息存入 DB1.
	})
}
