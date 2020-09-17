package main // 声明 main 包，表明当前是一个可执行程序

import "fmt" // 导入内置 fmt 包

func test() {
	fmt.Printf("测试")
}

var name string
var age int

const (
	_  = iota
	a1 = iota
	a2 = iota
)

func main() { // main函数，是程序执行的入口

	fmt.Println("Hello World!") // 在终端打印 Hello World!
	test()                      //调用同文件中的函数
	fmt.Println("")
	name = "zhangshan"
	fmt.Printf("name:%s\n", name)
	fmt.Println(a1)
	fmt.Println(a2)
}
