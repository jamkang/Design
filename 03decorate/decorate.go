package _3decorate

import "fmt"

//定义一个抽象组件
type Component interface {
	Operate()
}

//实现Component的一个组件

type Component1 struct {

}
func (c1 *Component1) Operate(){
	fmt.Println("实例化了一个Component组件")
}

//开始添加一个功能，定义一个装饰抽象
type Decorate interface {
	Component
	Do()
}

//实现Decorate装饰
type Decorate1 struct {
	Component
}

func (d1 *Decorate1) Do()  {
	fmt.Println("开始进行Decorate装饰")
}
func (d1 *Decorate1) Operate() {
	fmt.Println("这个是开始进行Decorate里面")
	d1.Do()
}