package _1simpleFactory

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

func (p2 *Product2) PrintNa()  {
	fmt.Println("我是产品2的管家，现在你正在使用产品2")
}


type productType int

const (
	p1 productType = iota
	p2
)
//实现简单工厂
type productFactory struct{

}

func (pf productFactory) Create(productType productType) Product {
	if productType== p1{
		return &Product1{}
	}else if productType==p2 {
		return &Product2{}
	}else {
		return nil
	}
}