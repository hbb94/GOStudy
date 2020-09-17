package main

import "fmt"

func main() {

	//1、定义int类型

	// var num int = 10
	// num = 22
	// fmt.Printf("num=%v 类型:%T", num, num)

	//2、int8的范围演示

	// var num int8 = 98
	// fmt.Printf("num=%v 类型:%T", num, num)

	// (-128 到 127)  错误
	// var num int8 = 130 //constant 130 overflows int8
	// fmt.Printf("num=%v 类型:%T", num, num)

	// var num int16 = 130 //ok
	// fmt.Printf("num=%v 类型:%T", num, num)

	//3、uint8的范围(0-255)

	// var n1 uint8 = -2 //constant -2 overflows uint8
	// fmt.Printf("n1=%v 类型:%T\n", n1, n1)

	// var n2 uint8 = 200 //constant -2 overflows uint8
	// fmt.Printf("n2=%v 类型:%T\n", n2, n2)

	//4、int8 int16 ... 占用的存储空间大小  unsafe.Sizeof 查看不同长度的整型 在内存里面的存储空间

	// var a int8 = 15
	// fmt.Printf("num=%v 类型:%T\n", a, a)
	// fmt.Println(unsafe.Sizeof(a)) // 1个字节

	// var a int32 = 15
	// fmt.Printf("num=%v 类型:%T\n", a, a)
	// fmt.Println(unsafe.Sizeof(a)) // 4个字节

	//5、int类型转换

	// var a1 int32 = 10
	// var a2 int64 = 21

	// fmt.Println(int64(a1) + a2) //把a1转换成64位
	// fmt.Println(a1 + int32(a2)) //把a2转换成32位

	//高位向底位转换的时候要注意

	// var n1 int16 = 130
	// fmt.Println(int8(n1)) //-126  有问题了

	// var n2 int16 = 110
	// fmt.Println(int8(n2)) //110

	// fmt.Println(int64(n2)) //110

	//6、数字字面量语法  %d 表示10进制输出 %b表示二进制输出 %o 八进制输出 %x 表示16进制输出(了解)

	// num := 30
	// fmt.Printf("num=%v 类型：%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num)) //表示64位的计算机 int就是int64   占用8个字节

	num := 12

	fmt.Printf("num=%v\n", num) //%v 原样输出

	fmt.Printf("num=%d\n", num) //%d 表示10进制输出

	fmt.Printf("num=%b\n", num) //%b 表示二进制输出

	fmt.Printf("num=%o\n", num) //%o 表示八进制输出

	fmt.Printf("num=%x\n", num) //%x 表示16进制输出

}
