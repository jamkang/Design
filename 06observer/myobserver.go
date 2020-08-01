package main

import "fmt"

/*我自己经过查看写出的观察者模式
一个小实例---广播*/

//主题抽象
type target interface {
	register(observer)
	unregister(observer)
	sent()
}

//观察者抽象
type observer interface {
	GetData()(int,string)
	get(string)
}


//-------------实例化一个主题
type weather struct {
	obs		[]observer
	value 	int
	data 	string
}

func NewWeather(data string) target {
	return &weather{
		data:data,
	}
}
//注册观察者
func (w *weather) register(ob observer){
	w.obs=append(w.obs,ob)
	w.value=len(w.obs)
	fmt.Println("注册weather成功")
	fmt.Printf("weath里面有注册数：%d个\n",w.value)
}
//注销注册者
func (w *weather) unregister(ob observer){
	for i,v:=range w.obs {
		if v==ob{
			w.obs=append(w.obs[:i],w.obs[i+1:]...)
			break
		}
	}
	fmt.Println("注销weather成功")
}
//发送消息
func (w *weather) sent(){
	for _,v:=range w.obs{
		v.get(w.data+"发送的")
	}
}





//----------实例化第二个主题
type Hotdata struct {
	obs		[]observer
	value 	int
	data 	string
}

func NewHotdata(data string) target {
	return &Hotdata{
		data:data,
	}
}
//注册观察者
func (w *Hotdata) register(ob observer){
	w.obs=append(w.obs,ob)
	w.value=len(w.obs)
	fmt.Println("注册Hotdata成功")
	fmt.Printf("weath里面有注册数：%d个\n",w.value)
}
//注销注册者
func (w *Hotdata) unregister(ob observer){
	for i,v:=range w.obs {
		if v==ob{
			w.obs=append(w.obs[:i],w.obs[i+1:]...)
			break
		}
	}
	fmt.Println("注销Hotdata成功")
}
//发送消息
func (w *Hotdata) sent(){
	for _,v:=range w.obs{
		v.get(w.data+"发送的")
	}
}





//----------实例化观察者
type observer1 struct {
	id		int
	name 	string
}

func (o *observer1) GetData()(int,string){
	return o.id,o.name
}

func (o *observer1) get(data string){
	fmt.Println(o.name+"接收到了"+data)
}