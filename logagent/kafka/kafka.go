package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//专门往kafka写日志的模块

var (
	client sarama.SyncProducer
)

// Init 初始化生产者
func Init(addrs []string)(err error){
	config:=sarama.NewConfig()
	config.Producer.RequiredAcks=sarama.WaitForAll //ACK反馈机制，all(需要leader 和follow都确认了)
	config.Producer.Partitioner=sarama.NewRandomPartitioner//指定写往哪个分区 新选出一个partition
	config.Producer.Return.Successes=true//成功交付的信息将在success channel返回

	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Printf("producer close the err is %v",err)
		return
	}
	//defer client.Close()  //不需要关闭 日志每时都在产生，所以不需要关闭
	return
}

func SendToKafka(topic,data string)(err error){
	//构造一个消息
	msg:=&sarama.ProducerMessage{}
	msg.Topic=topic
	msg.Value=sarama.StringEncoder(data)
	//发送消息
	message, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed err:",err)
		return err
	}
	fmt.Printf("pid:%v;offset:%v\n",message,offset)
	return
}