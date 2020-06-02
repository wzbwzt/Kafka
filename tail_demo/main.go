package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

//读日志 tailf第三方库的使用demo
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,//重新打开  切换文件时，创建新文件
		Follow:    true,//跟随文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},//从文件哪个地方开始读
		MustExist: false,//文件不存在不报错
		Poll:      true,//
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	var (
		line *tail.Line
		ok bool
	)
	for  {
		line,ok=<-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename is %s\n",tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:",line.Text)
	}
}
