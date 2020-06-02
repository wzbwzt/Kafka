package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"logagent/config"
	"logagent/kafka"
	"logagent/taillog"
	"time"
)

//logAgent 入口程序

var (
	iniFile =new(config.AppConf)
)

func run(){
	//1读取文件
	for{
		select {
		//2发送到kafka
		case line:=<-taillog.ReadChan():
			kafka.SendToKafka(iniFile.Topic,line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main(){
	//0.加载配置文件
	err := ini.MapTo(iniFile, "./config/logagent.ini")
	if err != nil {
		return
	}

	//1.初始化kafka连接
	err = kafka.Init([]string{iniFile.Address})
	if err != nil {
		fmt.Println("init kafka failed ,err:",err)
		return
	}
	fmt.Println("init kafka success!!")
	//2.打开日志文件准备收集日志
	err = taillog.Init(iniFile.Path)
	if err != nil {
		fmt.Println("init tail failed ,err:",err)
		return
	}
	fmt.Println("init tail success!!")
	run()
}
