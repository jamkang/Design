package _5singleton

import "sync"

//go中自带了单例模式--sync.Once
//sync.Once的Do方法可以实现在程序运行过程中只运行一次其中的回调

var once sync.Once
var gosc *goSingle
type goSingle struct {
	Cc string
}

func GetGosc() *goSingle{
	once.Do(
		func() {
			gosc =&goSingle{"8080"}
		})
	return gosc
}