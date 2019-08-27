package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

//rabbitmq的exchange为TransExchangeName，模式为direct，绑定到队列TransQueueName，Routing key为TransRouteKey

const (
	RabbitMqUrl       = "amqp://admin:Admin@114.67.227.157:5672";
	TransExchangeName = "msg.ex.oss"        //交换机名
	TransQueueName    = "msg.queue.oss"     //队列名
	TransErrQueueName = "msg.err.queue.oss" //异常处理队列
	TransRouteKey     = "oss"               //消息路由key
)

type TransferData struct {
	Msg  string
	Code int
}

var rabbitMqConn *amqp.Connection
var channel *amqp.Channel
var done chan bool

func ProcessTransfer(msg []byte) bool {

	//解析msg

	var recvData TransferData

	err := json.Unmarshal(msg, &recvData)

	if err != nil {
		return false
	}

	log.Printf("recvmsg code:%d message:%s", recvData.Code, recvData.Msg)

	return true
}

//RabbitMqConsumeDispatcher 消费消息
func RabbitMqConsumeDispatcher(queueName, consumeName string, callback func(msg []byte) bool) {

	// 通过channel获得消息信道

	msgChan, err := channel.Consume(
		queueName,
		consumeName,
		true,  //自动确认
		false, // 非唯一的消费者
		false, // rabbitMQ只能设置为false
		false, // noWait, false表示会阻塞直到有消息过来
		nil)

	if err != nil {
		log.Printf("get consume chan failed %s", err.Error())
		return
	} else {
		log.Printf("get consume chan succeed")
	}

	done = make(chan bool)

	go func() {

		//循环读取channel数据
		for msg := range msgChan {

			if processResult := callback(msg.Body); !processResult {
				log.Printf("parser message failed: %v", string(msg.Body))
			}

		}
	}()

	//等待退出
	<-done

	//关闭mq
	_ = channel.Close()

}

// StopConsume : 停止监听队列
func StopConsume() {
	done <- true
}

func initChannel() bool {

	//判断channel是否创建过了

	if channel != nil {
		return true
	}

	//获取rabbitmq连接
	conn, err := amqp.Dial(RabbitMqUrl)

	if err != nil {
		log.Printf(err.Error())
		return false;
	}

	//打开一个channel

	channel, err = conn.Channel()

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func RabbitMqPublic(exchange, routeKey string, msg []byte) bool {

	if !initChannel() {
		return false
	}

	//发布消息

	err := channel.Publish(
		exchange, routeKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		log.Println(err.Error())
		return false
	} else {
		log.Printf("message send succeed")
	}

	return true

}

func main() {

	sendData, err := json.Marshal(TransferData{
		Msg:  "test message",
		Code: 200,
	})

	if err != nil {
		log.Printf("marshal failed ")
		return
	}

	RabbitMqPublic(TransExchangeName, TransRouteKey, sendData)

	RabbitMqConsumeDispatcher(TransQueueName, "consume-ch-1" /*消费者名字，随便定义*/, ProcessTransfer)

}
