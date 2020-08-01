package _4strategy

import "fmt"

//实现一个上下文类
type Context struct {
	Strategy
}

//抽象的策略
type Strategy interface {
	Do()
}

//实例化Strategy做第一个动作
type Strategy1 struct {

}

func (s *Strategy1)Do(){
	fmt.Println("我是在这里做第一个动作")
}

//实例化Strategy做第二个动作
type Strategy2 struct {

}

func (s *Strategy2)Do(){
	fmt.Println("我是在这里做第二个动作")
}