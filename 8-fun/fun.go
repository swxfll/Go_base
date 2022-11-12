package main

import (
	"fmt"
)

// 1.自定义函数
// 驼峰命令
// 首字母大写该函数可以被本包文件和其它包文件使用(类似public)
// 首字母小写只能被本包文件使用,其它包文件不能使用(类似private)?
func cal(num1 int, num2 int) int {
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

// 2.多返回值
func cal2(num1 int, num2 int) (int, int) {
	s1 := num1 + num2
	s2 := num1 - num2
	return s1, s2
}

// 3.可变参数
// 可以传入任意数量的int类型的数据
func test(args ...int) {
	// 将可变参数当做切片处理
	// 遍历可变参数
	for i := 0; i < len(args); i++ {
		fmt.Print(args[i])
	}
}

// 4.
// 基本数据类型和数据默认都是值传递,即进行值拷贝.在函数内修改,不会影响到原来的值
// 以值传递方式的数据类型,如果希望在函数内的变量能修改函数外的变量,可以传入变量的地址&,函数内以指针的方式操作变量
func pointer(num *int) {
	*num += 10
}

// 5
// 在GO中,函数也是一种数据类型,可以赋值给一个变量,则该变量就是一个函数类型的变量了.通过该变量可以对函数调用
func variableFun(num int) {
	fmt.Println(num)
}

// 6
// 函数既然是一种数据类型,因此在Go中,函数可以作为形参,并且调用(把函数本身当做一种数据类型)
func argsIsFun(num1 int, num2 int, testFunc func(int)) {
	fmt.Println("argsIsFun")
}

// 7
// 给函数起别名
type myfun func(int)

func argsIsFun2(num1 int, num2 int, testFunc myfun) {
	fmt.Println("argsIsFun")
}

// 8
// 支持对函数返回值命名
func modReturnName(num1 int, num2 int) (rs2 int, rs1 int) {
	rs1 = num1 + num2
	rs2 = num1 - num2
	return
}

func main() {
	//1==========================================
	// 功能 10 + 20
	sum := cal(10, 20)
	fmt.Println(sum) // 输出30

	//2==========================================
	r1, r2 := cal2(10, 20)
	fmt.Println(r1) // 30
	fmt.Println(r2) // -10

	// 多返回值,忽略其中的一个返回值
	r3, _ := cal2(10, 20)
	fmt.Println(r3) // 30

	//3==========================================
	//可变参数
	test(1, 2, 3) // 输出123

	//4==========================================
	var num int = 10
	fmt.Println(num) // 输出 10
	pointer(&num)
	fmt.Println(num) // 输出 20

	//5==========================================
	a := variableFun
	fmt.Println(a) // 输出 0xde860
	a(10)          // 输出 10

	//6==========================================
	argsIsFun(1, 2, variableFun) // 输出 argsIsFun

	//7==========================================
	// 为了简化数据类型定义,Go支持自定义数据类型
	// 基本语法: type 自定义数据类型名 数据类型
	// 可以理解为: 相当于起了一个别名
	type myInt int

	var num1 myInt = 30
	fmt.Println(num1) // 输出 30

	var num2 int = 10
	//num2 = num1  //虽然都是整形,但是go在编译时,还是认为myInt和int不是同一种数据类型
	num2 = int(num1)
	fmt.Println(num2) // 输出 30

	//函数别名
	argsIsFun2(1, 2, variableFun)

	//8==========================================
	fmt.Println(modReturnName(1, 2))
}
