package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)


func main() {
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll //ACK反馈机制，all(需要leader 和follow都确认了)
	config.Producer.Partitioner=sarama.NewRandomPartitioner//指定写往哪个分区 新选出一个partition
	config.Producer.Return.Successes=true//成功交付的信息将在success channel返回
	//构造一个消息
 	msg:=&sarama.ProducerMessage{}
 	msg.Topic="web_log"
 	msg.Value=sarama.StringEncoder("this is a test log")
 	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("producer close the err is %v",err)
		return
	}
	defer client.Close()
	//发送消息
	message, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed err:",err)
		return
	}
	fmt.Printf("pid:%v;offset:%v\n",message,offset)

}
