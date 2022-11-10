package main

import "fmt"

func main() {
	var count int = 20

	if count < 30 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// 定义并判断
	if count2 := 20; count2 < 30 {
		fmt.Println("true")
	}

	// 多分支
	var score int = 60
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 { // 隐藏了 score < 90
		fmt.Println("B")
	} else if score >= 70 { // 隐藏了 score < 80
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
