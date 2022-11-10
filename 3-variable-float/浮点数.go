package main

import "fmt"

func main() {
	var num1 float32 = 3.14
	fmt.Println(num1)

	var num2 float32 = -3.14
	fmt.Println(num2)

	var num3 float32 = 314e+2
	fmt.Println(num3) // 31400

	var num4 float32 = 314e-2
	fmt.Println(num4) // 3.14

	var num5 float64 = 314e-2
	fmt.Println(num5) // 3.14

	// float可能会有精度的损失,浮点数默认是float64
	var num6 float32 = 256.0000000916
	fmt.Println(num6) // 256

	var num7 float64 = 256.0000000916
	fmt.Println(num7) // 256.0000000916
}
