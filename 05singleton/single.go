//package _5singleton
//
//var sc *Single
//
//type Single struct {
//	Cc string
//}
//
//func GetSile(cc string) *Single {
//	if sc==nil || sc.Cc!=cc{
//		sc = &Single{Cc:cc}
//	}
//	return sc
//}
/*
自己创建一个单例模式，为了防止创建多个对象可以进行加锁
 */
//package _5singleton
//
//import "sync"
//
//var sc *Single
//var sclack sync.Mutex
//type Single struct {
//	Cc string
//}
//
//func GetSile(cc string) *Single {
//	sclack.Lock()
//	defer sclack.Unlock()
//	if sc==nil || sc.Cc!=cc{
//		sc = &Single{Cc:cc}
//	}
//	return sc
//}
/*以上方式为懒汉模式，以下是饿汉
两者的区别就是：
懒汉模式：需要的时候进行调用创建并且返回所需要的东西
饿汉：会在包加载时进行创立，不需要手动添加，这样方便，但是启动程序的时候速度会慢一些*/
package _5singleton

var sc *Single

type Single struct {
	Cc string
}

func init(){
	sc = &Single{"8081"}
}

func GetSile() *Single {
	return sc
}