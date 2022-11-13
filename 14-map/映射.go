package main

import (
	"fmt"
)

func main() {
	// 方式1
	var a map[int]string
	a = make(map[int]string, 10)
	a[2009] = "张三"
	a[2010] = "李四"
	fmt.Println(a) // map[2009:张三 2010:李四]

	// 方式2
	b := make(map[int]string)
	b[2009] = "张三"
	b[2010] = "李四"
	fmt.Println(b) // map[2009:张三 2010:李四]

	// 方式3
	c := map[int]string{
		2009: "张三",
		2010: "李四",
	}
	fmt.Println(c) // map[2009:张三 2010:李四]

	// 对map的操作
	// 删除
	delete(c, 2009)
	fmt.Println(c) // map[2010:李四]

	// 查找
	v, flag := c[2010]
	fmt.Println(v)    // 李四
	fmt.Println(flag) // true

	// 获取map的长度
	d := map[int]string{
		2009: "张三",
		2010: "李四",
		2011: "王五",
	}
	fmt.Println(len(d)) // 3

	// map只支持for-range遍历
	for k, v := range d {
		fmt.Println(k, v)
	}

	// 加深难度
	e := make(map[string]map[int]string)
	e["班级1"] = make(map[int]string, 3)
	e["班级1"][2009] = "张三"
	e["班级1"][2010] = "李四"

	e["班级2"] = make(map[int]string, 3)
	e["班级2"][2009] = "王五"
	e["班级2"][2010] = "赵六"

	fmt.Println(e) // map[班级1:map[2009:张三 2010:李四] 班级2:map[2009:王五 2010:赵六]]

	for k, _ := range e {
		for key, value := range e[k] {
			fmt.Println(key, value)
			/*
				2009 张三
				2010 李四
				2010 赵六
				2009 王五
			*/
		}
	}

}
