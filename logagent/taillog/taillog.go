package taillog

import (
	"github.com/hpcloud/tail"
)

//专门从日志文件中收集日志文件的模块

var (
	tailObj *tail.Tail
)

//Init 打开日志文件初始化
func Init(fileName string)(err error){
	config := tail.Config{
		ReOpen:    true,//重新打开  切换文件时，创建新文件
		Follow:    true,//跟随文件
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},//从文件哪个地方开始读
		MustExist: false,//文件不存在不报错
		Poll:      true,//
	}
	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		return
	}
	return
}

//ReadChan 循环读取文件
func ReadChan() <-chan *tail.Line{
	return tailObj.Lines
}