package _4strategy

import "fmt"

/*使用策略模式
实现一个日志记录模拟器
可以使用文件进行存储、数据库两种方式进行存取*/

//日志上下文
type LogContext struct {
	LogStra
}
//实现创建一个日志类
func NewLogContext(log LogStra) *LogContext{
	return &LogContext{log}
}

//创建一个日志模拟器的抽象
type LogStra interface {
	Info()
	Error()
}
//实例化一文件
type File struct {

}

func (f *File) Info(){
	fmt.Println("文件记录info")
}
func (f *File) Error(){
	fmt.Println("文件记录error")
}

//实例化数据库
type data struct {

}
func (d *data) Info(){
	fmt.Println("数据库记录info")
}
func (d *data) Error(){
	fmt.Println("数据库记录error")
}