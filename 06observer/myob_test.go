package main

import "testing"

//实例广播的测试
func TestConcreteObserver1_Update(t *testing.T) {
	wth :=NewWeather("今天天气下大雨")
	hot :=NewHotdata("今天停电")
	ob1 :=&observer1{
		name:"小二",
		id:1,
	}
	wth.register(ob1)
	ob2 :=&observer1{
		name:"小三",
		id:1,
	}
	wth.register(ob2)
	hot.register(ob2)
	wth.sent()
	hot.sent()
}
