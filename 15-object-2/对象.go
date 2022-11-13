package main

import (
	"fmt"
)

// 1. /////////////////////////////////////////
type Student struct {
	Age int
}

type Person struct {
	Age int
}

// 2. /////////////////////////////////////////
// Stu 使用type重新定义(相当于起别名),go认为是新的数据类型,但是互相间可以强转
type Stu Student

// 3. /////////////////////////////////////////
// 定义结构体
type Person2 struct {
	Name string
}

// 给Person2结构体绑定方法, (p Person2)体现方法test和结构体Person2绑定关系
func (p Person2) test() {
	p.Name = "李四"
	fmt.Println(p.Name)
}

func (p *Person2) change() {
	p.Name = "王五"
	fmt.Println(p.Name)
}

func main() {
	// 1. /////////////////////////////////////////
	var s Student = Student{10}
	var p Person = Person{20}
	s = Student(p)
	fmt.Println(s) // {20}
	fmt.Println(p) // {20}

	// 2. /////////////////////////////////////////
	var s1 Student = Student{19}
	var s2 Stu = Stu{99}
	s2 = Stu(s1)
	fmt.Println(s2) // {19}

	// 3. /////////////////////////////////////////
	// 创建结构体对象
	var p2 Person2
	p2.Name = "张三"
	// 结构体对象传入test方法中, 是值传递
	p2.test()            // 输出 李四
	fmt.Println(p2.Name) // 输出 张三

	fmt.Println("///////////////////")

	// 简写 p2.test2(), 底层会自动转为 (&p2).test2()
	(&p2).change()       // 输出 王五
	fmt.Println(p2.Name) // 输出 王五
}
