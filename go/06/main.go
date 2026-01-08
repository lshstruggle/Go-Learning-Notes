package main

import "fmt"

func MyPrint(a interface{}) {
	if b, ok := a.(string); ok {
		println("这是一个字符串", b)
	}
}

func MyPrint2(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Println("这是一个字符串", a)
	case int:
		fmt.Println("这是一个整数", a)
	default:
		fmt.Println("不知道什么类型")
	}
}

func main() {
	MyPrint2("hello")
}
