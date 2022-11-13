package main

import "fmt"

/*
Go中的方法作用在指定的数据类型上,和指定的数据类型绑定,
因此自定义类型,都可以有方法,而不仅仅是struct, 比如int,float32都可以有方法
*/
type integerll int

func (n1 integerll) test() {
	n1 = 999
	fmt.Println(n1)
}

func (n1 *integerll) test2() {
	*n1 = 666
	fmt.Println(*n1)
}

func main() {
	var num1 integerll = 20
	num1.test()       // 输出 999
	fmt.Println(num1) // 20

	num1.test2()      // 输出 666
	fmt.Println(num1) // 666
}
