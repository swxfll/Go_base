package main

import (
	"errors"
	"fmt"
)

// 捕获异常
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

// 自定义异常
func test2() (er error) {
	num1 := 10
	num2 := 0

	if num2 == 0 {
		//抛出自定义异常
		return errors.New("除数不能为0")
	}

	result := num1 / num2
	fmt.Println(result)
	return nil
}

func main() {
	// 1 捕获异常
	test()
	fmt.Println("hello")

	// 2 自定义异常
	err := test2()
	if err != nil {
		fmt.Println(err) // 输出 除数不能为0
		//panic(err)       // 停止当前Go程序(退出main函数),不继续往下执行
	}

	fmt.Println("main")
}
