package main

import "fmt"

//相对于变量，常量是恒定不变的值，
// 多用于定义程序运行期间不会改变的那些值。
// 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

const pi = 3.1415
const e = 2.7182

const (
	a = 1
	b = 2
)

//const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 1
	n2 //n2=1
	n3 //n3=1
)

func main() {
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)

}
