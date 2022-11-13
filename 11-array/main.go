package main

import (
	"fmt"
)

func main() {
	// 定义一个数组
	var scores [5]int
	scores[0] = 60
	scores[1] = 70
	scores[2] = 80
	scores[3] = 90

	fmt.Printf("数组的地址为: %p", &scores)         // 输出 数组的地址为: 0xc00000e3c0
	fmt.Printf("数组第一个元素的地址为: %p", &scores[0]) // 输出 0xc00000e3c0
	//与上一个地址相比,增加了8,因为数组的元素是int64类型, 64bit=8byte
	fmt.Printf("数组第二个元素的地址为: %p", &scores[1]) // 输出 0xc00000e3c8

	// 遍历方式1
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	// 遍历方式2 for-range
	for key, val := range scores {
		fmt.Println(key, val)
	}

	for _, val := range scores {
		fmt.Println(val)
	}

	fmt.Println("/////////////////////定义数组/////////////////////")
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [3]int{4, 5, 6}
	fmt.Println(arr2)

	var arr3 = [...]int{7, 8, 9}
	fmt.Println(arr3)

	// 指定索引位置的元素
	var arr4 = [...]int{0: 10, 2: 12}
	fmt.Println(arr4)

	fmt.Println("/////长度属于数组类型的一部分/////////////////////")
	var arr5 = [3]int{1, 2, 3}
	var arr6 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T", arr5) // 输出 [3]int
	fmt.Printf("%T", arr6) // 输出 [5]int

	fmt.Println("/////GO数组属于值类型,在默认情况下是值传递,因此会进行值拷贝//////")
	var arr7 = [3]int{1, 2, 3}
	test(arr7)
	fmt.Println(arr7) // 输出 [1 2 3]

	test1(&arr7)
	fmt.Println(arr7) // 输出 [999 2 3]
}

func test(arr [3]int) {
	arr[0] = 999
}

func test1(arr *[3]int) {
	arr[0] = 999
}
