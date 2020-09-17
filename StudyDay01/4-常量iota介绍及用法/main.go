package main

import "fmt"

// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
// 使用iota能简化定义，在定义枚举时很有用。
func main() {
	const (
		n1 = iota //0
		n2        //1  n2=iota=1
		n3        //2  n3=iota+1=2
		n4        //3  n4=iota+1=3
	)
	fmt.Println("n1=", n1)
	fmt.Println("n2=", n2)
	fmt.Println("n3=", n3)
	fmt.Println("n4=", n4)

	//使用_跳过某些值
	const (
		n5 = iota //0
		n6        //1
		_         //2
		n8        //3
	)
	fmt.Println("n5=", n5)
	fmt.Println("n6=", n6)
	fmt.Println("n8=", n8)

	//iota声明中间插队
	const (
		a1 = iota //0
		a2 = 100  //100 iota+1
		a3 = iota //2   iota+1
		a4        //3
	)
	const a5 = iota //0
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)
	fmt.Println("a3=", a3)
	fmt.Println("a4=", a4)
	fmt.Println("a5=", a5)

	// 定义数量级 （这里的<<表示左移操作，
	// 			1<<10表示将1的二进制表示向左移10位，
	// 			也就是由1变成了10000000000，也就是十进制的1024。
	// 			同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，
	// 			也就是十进制的8。）
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println("KB=", KB)
	fmt.Println("MB=", MB)
	fmt.Println("GB=", GB)
	fmt.Println("TB=", TB)
	fmt.Println("PB=", PB)

	//多个iota定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
}
