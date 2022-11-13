package main

import "fmt"

func main() {
	// int16 是16位,占2个字节
	var arr [2][3]int16
	fmt.Println(arr) // [[0 0 0] [0 0 0]]

	fmt.Printf("%p", &arr) // 0xc00001a0c0
	fmt.Println()
	fmt.Printf("%p", &arr[0]) // 0xc00001a0c0
	fmt.Println()
	fmt.Printf("%p", &arr[0][0]) // 0xc00001a0c0
	fmt.Println()
	fmt.Printf("%p", &arr[1]) // 0xc00001a0c6
	fmt.Println()
	fmt.Printf("%p", &arr[1][0]) // 0xc00001a0c6
	fmt.Println()
	fmt.Printf("%p", &arr[1][1]) // 0xc00001a0c8

	// 二维数组初始化
	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1) // [[1 2 3] [4 5 6]]

	var arr2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)

	var arr3 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)

	var arr4 = [...][3]int{{0: 1}, {1: 5}}
	fmt.Println(arr4) // [[1 0 0] [0 5 0]]

	// 二维数组遍历
	var arr5 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// 普通for循环
	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Print(arr5[i][j], " ") // 1 2 3 4 5 6 7 8 9
		}
	}

	// for range方式
	for key, _ := range arr5 {
		for k, v := range arr5[key] {
			fmt.Print(arr5[key][k], " ") // 1 2 3 4 5 6 7 8 9
			fmt.Print(v, " ")            // 1 2 3 4 5 6 7 8 9
		}
	}
}
