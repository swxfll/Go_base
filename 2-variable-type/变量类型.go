package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 127
	fmt.Println(num1) // 127

	var num2 uint8 = 255
	fmt.Println(num2) // 255

	var num3 = 28
	fmt.Printf("num3的类型是: %T", num3) // num3的类型是: int
	fmt.Println(unsafe.Sizeof(num3)) // 8
}
