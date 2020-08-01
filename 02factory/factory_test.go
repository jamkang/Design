package _2factory

import "testing"

func TestFactory1_Create(t *testing.T) {
	factory1 :=&factory1{}
	f1 :=factory1.Create()
	f1.SetName("小七七")
	f1.GetName()
}