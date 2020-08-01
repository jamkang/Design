package _3decorate

import "testing"

func TestComponent1_Operate(t *testing.T) {
	c1 :=&Component1{}
	c1.Operate()
}

func TestDecorate1_Do(t *testing.T) {
	d1 :=Decorate1{}
	d1.Operate()
}