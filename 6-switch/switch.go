package main

import "fmt"

func main() {
	var score int = 101
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
	case 8:
		fmt.Println("C")
	default:
		fmt.Println("默认")
	}
	// 输出 A

	// 写法2 case后面携带多个值
	var a int32 = 100
	switch a {
	case 10, 100, 1000:
		fmt.Println("命中10,100,1000")
	}
	// 输出 命中10,100,1000

	// 3 怪异写法 当做if使用
	var b = 2
	switch {
	case b == 1:
		fmt.Println("1")
	case b == 2:
		fmt.Println("2")
	}
	// 输出 2

	// 4 怪异写法
	switch c := 7; {
	case c > 6:
		fmt.Println("大于6")
	case c <= 6:
		fmt.Println("不大于6")
	}
	// 输出 大于6

	// 5 fallthrough 穿透
	var d = 6
	switch d {
	case 6:
		fmt.Println("6")
		// 往下穿透一层
		fallthrough
	case 7:
		fmt.Println("7")
	}
	// 输出 6 7
}
