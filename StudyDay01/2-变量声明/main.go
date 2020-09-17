package main

import "fmt"

//变量声明
var name string = "huangbiao" //也可在声明变量的时候为其指定初始值
var age int
var isOK bool

//批量声明变量
var (
	a string
	b int
	c bool
)

//类型推导声明变量
var d = "zhangshan"
var e = 12

//f := "lishi" 短变量声明只能在函数里面

func foo() (int, string) {
	return 10, "wangwu"
}

func main() {

	age = 15

	//短变量声明
	f := "lishi"
	g := 13
	fmt.Printf("f=%s,g=%d\n", f, g) //局部变量声明后需要使用，否则将会报错

	//匿名变量声明,匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	h, _ := foo()
	_, i := foo()
	fmt.Printf("h=%d i=%s", h, i)
}

// 注意：
// 1. 函数外的每个语句都必须以关键字开始（var、const、func等）
// 2. :=不能使用在函数外。
// 3. _多用于占位，表示忽略值
