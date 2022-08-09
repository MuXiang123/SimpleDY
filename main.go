package main

import (
	"SimpleDY/initial"
	"SimpleDY/middleware/redis"
	"SimpleDY/router"
)

func main() {
	initial.LoadConfig()
	initial.InitMysql() //初始化mysql
	redis.InitRedis()   //初始化redis
	r := router.InitRouter()
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
