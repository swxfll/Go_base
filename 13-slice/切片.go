package main

import "fmt"

func main() {
	/*
		切片(slice)是go中一种特有的数据类型
		数组的长度固定不可变,所以在Go代码中并不常见,相对的切片却是随处可见的,
		切片是一种建立在数组类型之上的抽象,所以切片是一个引用类型.
	*/

	// 1. 演示切片///////////////////////////////////
	var arr = [6]int{3, 6, 9, 1, 4, 7}
	// 切片构建在数组之上:
	// [0:3]切片 索引从0开始,到3结束(不包含3)
	//var slice []int = arr[0:3]
	slice := arr[0:3]
	fmt.Println(slice) // [3 6 9]

	// 切片元素的个数
	fmt.Println(len(slice)) // 3

	// 获取切片的容量: 容量可以动态变化
	fmt.Println(cap(slice)) // 6

	// 2. 切片内存分析///////////////////////////////////
	fmt.Printf("数组中下标为0的地址: %p", &arr[0])   // 0xc00000e3c0
	fmt.Printf("切片中下标为0的地址: %p", &slice[0]) // 0xc00000e3c0

	slice[0] = 999
	fmt.Println(slice) // [999 6 9]
	fmt.Println(arr)   // [999 6 9 1 4 7]

	// 3. 切片的定义///////////////////////////////////
	// make函数的三个参数: 1.切片类型 2.切片长度, 3 切片容量
	slice2 := make([]int, 4, 20)
	fmt.Println(slice2)               // [0 0 0 0]
	fmt.Println("切片的长度", len(slice2)) // 4
	fmt.Println("切片的容量", cap(slice2)) // 20
	slice2[0] = 66
	slice2[1] = 88
	fmt.Println(slice2) // [66 88 0 0]

	// 4 切片的遍历
	slice3 := make([]int, 4, 20)
	slice3[0] = 66
	slice3[1] = 88
	slice3[2] = 99
	slice3[3] = 100

	// 方式1: 普通for循环
	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}

	// 方式2 for-range 循环
	for i, v := range slice3 {
		fmt.Println(i, v)
	}

	// 5 切片可以继续切片
	// 定义数组
	var arr2 [6]int = [6]int{1, 4, 7, 2, 5, 8}
	slice4 := arr2[:]
	slice5 := slice4[:]
	slice5[0] = 999

	fmt.Println(arr2)   // [999 4 7 2 5 8]
	fmt.Println(slice5) // [999 4 7 2 5 8]

	// 6 切片可以动态增长
	var arr3 [6]int = [6]int{1, 4, 7, 2, 5, 8}
	slice6 := arr3[0:3]
	fmt.Println(len(slice6)) // 3

	slice6 = append(slice6, 88, 55)
	fmt.Println(slice6) // [1 4 7 88 55]

	// 通过append函数讲切片追加给切片
	slice7 := []int{99, 999}
	slice6 = append(slice6, slice7...)
	fmt.Println(slice6) // [1 4 7 88 55 99 999]

	// 切片的拷贝
	var a []int = []int{1, 2, 3}
	var b []int = make([]int, 10)
	copy(b, a)     // 讲a中的元素复制到b中
	fmt.Println(b) // [1 2 3 0 0 0 0 0 0 0]

}
