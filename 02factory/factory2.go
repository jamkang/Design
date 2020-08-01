package _2factory

import "fmt"

// 创建一个抽象类的产品
type Product interface {
	SetName(name string)
	GetName() string
}

//实例化Product的产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name=name
}

func (p1 *Product1) GetName()string  {
	fmt.Println("工厂一的名为："+p1.name)
	return p1.name
}


//实例化Product的产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name=name
}

func (p2 *Product2) GetName()string  {
	fmt.Println("工厂二的名为："+p2.name)
	return p2.name
}

//工厂模式的抽象工厂
type Factory interface {
	Create() Product
}

//实现一个具体的工厂
type factory1 struct {

}

func (f1 *factory1) Create()Product {
	return &Product1{}
}

//实现第二个具体的工厂
type factory2 struct {

}

func (f2 *factory2) Create()Product {
	return &Product2{}
}