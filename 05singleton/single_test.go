package _5singleton

import (
	"testing"
	"fmt"
)

func TestGetSile(t *testing.T) {
	//测试go自带的单例
	getsi:=GetGosc()
	fmt.Println(getsi.Cc)
	//测试自己写得单例
	getsi2 :=GetSile("8081")
	fmt.Println(getsi2.Cc)
}

