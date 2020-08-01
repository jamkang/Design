package _1simpleFactory

import (
	"testing"
)

func TestDesign(t *testing.T)  {
	t.Run("简单工厂模式",testProduct)
}





//测量简单工厂函数
func testProduct(t *testing.T) {
	//使用普通的方法
	pro1 := &Product1{}
	pro1.SetName("小二")
	pro1.GetName()
	pro2 := &Product2{}
	pro2.SetName("小三")
	pro2.GetName()

	//使用简单工厂进行调用
	product :=productFactory{}
	p2:=product.Create(p2)
	p2.SetName("小四")
	p2.GetName()

	p1:=product.Create(p1)
	p1.SetName("小五")
	p1.GetName()
	//使用简单工厂的好处就是，只要找到创造产品的类和参数，就能创造出不同产品
}