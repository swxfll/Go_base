package main

import "fmt"

func main() {
	// +加号:
	// 1.整数 2.相加操作 3.字符串拼接
	var n1 int = +10
	fmt.Println(n1) // 10

	var n2 int = 4 + 7
	fmt.Println(n2) // 11

	var s1 string = "abc" + "def"
	fmt.Println(s1) // abcdef

	// /除号:
	// 两个int类型数据运算, 结果一定为整数类型
	fmt.Println(10 / 3) // 3
	// 浮点类型参与运算,结果为浮点类型
	fmt.Println(10.0 / 3) // 3.3333333333333335

	// % 取模
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	// ++自增操作
	// go语言里, ++ -- 操作只能单独使用,不能参与到运算符中
	// go语言里, ++ -- 只能在变量的后面, 不能写在变量的前面 --a, ++a是错误的写法
	var a int = 10
	a++
	fmt.Println(a) // 11

	// = 赋值运算符
	var num1 int = (10 + 20) * 2
	fmt.Println(num1) // 60

	var num2 int = 10
	num2 += 20
	fmt.Println(num2) // 30

	// 练习: 交换两个数的值并输出结果
	var x int = 6
	var y int = 9
	fmt.Printf("x = %v,y = %v", x, y) // x = 6,y = 9

	var t int
	t = x
	x = y
	y = t
	fmt.Printf("x = %v,y = %v", x, y) // x = 9,y = 6

	// 关系运算符
	fmt.Println(5 == 9) // false
	fmt.Println(5 != 9) // true
	fmt.Println(5 > 9)  // false

	// 逻辑运算符, 短路与 &&
	fmt.Println(true && true)   // true
	fmt.Println(false && false) // false
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false

	// 短路与 ||
	fmt.Println(true || true)   // true
	fmt.Println(false || false) // false
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true

	// 取反 !
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

	// 位运算度

	// 其它运算法
	var age int = 18
	fmt.Println("age对对应的存储空间的地址为: ", &age) // age对应的存储空间的地址为:  0xc0000a61a0

	// 定义指针变量
	var ptr *int = &age
	println(ptr)  // 0xc0000a61a0
	println(*ptr) // 18
}
