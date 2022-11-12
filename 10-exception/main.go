package main

import "fmt"

func test() {
	// 使用defer+recover来捕获错误: defer后加上匿名函数的调用
	defer func() {
		// 调用 recover 内置函数, 可以捕获错误
		err := recover()
		// 如果没有捕获错误,返回值为零值:nil
		if err != nil {
			fmt.Println("错误已捕获")
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	test()
	fmt.Println("hello")
}
