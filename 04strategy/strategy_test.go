package _4strategy

import "testing"

func TestStrategy1_Do(t *testing.T) {
	//创建一个上下文类
	con :=&Context{}

	//实现第一个动作
	str1 :=&Strategy1{}
	con.Strategy=str1
	con.Do()

	//实现第二个动作
	str2 :=&Strategy2{}
	con.Strategy=str2
	con.Do()
}
//测试实战的
func TestFile_Error(t *testing.T) {
	//使用文件进行记录
	file:=&File{}
	con :=NewLogContext(file)

	con.Error()
	con.Info()

	//使用数据库进行记录
	data:=&data{}
	con1 :=NewLogContext(data)

	con1.Error()
	con1.Info()
}