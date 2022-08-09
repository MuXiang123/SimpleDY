// Package rabbitMQ
// @author    : MuXiang123
// @time      : 2022/7/9 15:20
package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var MQURL = "amqp//simpledy:simplydy@127.0.0.1"

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

var rabbit *RabbitMQ

func InitRabbitMQ() {
	rabbit = &RabbitMQ{mqurl: MQURL}
	//接收一个url，返回一个tcp连接
	dial, err := amqp.Dial(rabbit.mqurl)
	rabbit.failOnErr(err, "创建连接失败")
	rabbit.conn = dial
}

//输出错误信息
func (rabbit *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		//连接失败 输出错误日志
		log.Fatalf("%s:%s\n", err, message)
		//恐慌
		panic(fmt.Sprintf("%s:%s\n", err, message))
	}
}

//关闭通道
func (rabbit *RabbitMQ) destroy() {
	rabbit.conn.Close()
}
