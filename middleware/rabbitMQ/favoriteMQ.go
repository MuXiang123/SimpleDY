// Package rabbitMQ
// @author    : MuXiang123
// @time      : 2022/7/9 15:19
package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

// FavoriteMQ rabbitMQ出来点赞操作
type FavoriteMQ struct {
	RabbitMQ                //连接实体
	channel   *amqp.Channel //通道
	queueName string        //队列名
	exchange  string        //交换机名
	key       string        //键
}

// NewFavoriteMQ 获取favoriteMQ对应的队列
func NewFavoriteMQ(queueName string) *FavoriteMQ {
	fmq := &FavoriteMQ{
		RabbitMQ:  *rabbit,
		queueName: queueName,
	}
	channel, err := fmq.conn.Channel()
	fmq.channel = channel
	fmq.failOnErr(err, "获取通道失败")
	return fmq
}

// Publish 将点赞发布到队列中
func (favorite FavoriteMQ) Publish(message string) {
	//声明一个队列传递给消费者
	_, err := favorite.channel.QueueDeclare(
		favorite.queueName,
		false, //持久化
		false, //自动删除
		false, //排他性
		false, //阻塞
		nil,   //其他属性
	)
	if err != nil {
		log.Fatalf("basic.publish: %v", err)
		panic(err)
	}
	err1 := favorite.channel.Publish(
		favorite.exchange,
		favorite.queueName,
		false, //不强制发布
		false, //不用立刻
		//设置放入的信息和类型
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err1 != nil {
		log.Fatalf("basic.publish: %v", err)
		panic(err)
	}
}

// Consume 消费者
func (favorite FavoriteMQ) Consume() {
	//声明一个队列
	_, err := favorite.channel.QueueDeclare(favorite.queueName, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	message, err := favorite.channel.Consume(
		favorite.queueName,
		"", //消费者名称
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	//判断是点赞还是取消赞，然后进行写数据的逻辑
	if strings.Compare(favorite.queueName, "like_add") == 0 {
		go favorite.consumerLikeAdd(message)
	} else if strings.Compare(favorite.queueName, "like_del") == 0 {
		go favorite.consumerLikeDel(message)
	}
}

//点赞的消费方式
func (favorite *FavoriteMQ) consumerFavoriteAdd(messages <-chan amqp.Delivery) {
	for msg := range messages {
		//解析字符串
		split := strings.Split(fmt.Sprintf("%s", msg.Body), " ")
		userId, _ := strconv.ParseInt(split[0], 10, 64)
		videoId, _ := strconv.ParseInt(split[1], 10, 64)

	}
}

//取消赞的消费方式
func (favorite *FavoriteMQ) consumerFavoriteDel(messages <-chan amqp.Delivery) {

}
