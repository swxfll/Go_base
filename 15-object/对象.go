package main

import "fmt"

type Teacher struct {
	// 变量名大写,外界可以访问这个属性
	Name   string
	Age    int
	School string
}

func main() {
	// 创建老师结构体的实例
	var t1 Teacher
	fmt.Println(t1) // { 0 }

	t1.Name = "test"
	t1.Age = 20
	t1.School = "school"
	fmt.Println(t1) // {test 20 school}

	// 方式2
	var t2 Teacher = Teacher{"张三", 25, "清华"}
	fmt.Println(t2) // {张三 25 清华}

	// 方式3
	var t3 *Teacher = new(Teacher)
	(*t3).Name = "李四"
	(*t3).Age = 999
	// go提供了简化的赋值方式
	t3.School = "清华" //go编译器底层对t3.School转化为(*t3).School
	fmt.Println(*t3) // {李四 999 清华}

	// 方式4
	var t4 *Teacher = &Teacher{"王五", 555, "清华"}
	(*t4).Age = 999
	fmt.Println(*t4) // {王五 999 清华}}
}
