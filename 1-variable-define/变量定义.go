package main //声明文件所在的包,每个go文件必须有归属的包

//引入程序中需要用的包,为了使用包下的函数
import "fmt" //引入程序中需要用的包,为了使用包下的函数

// 全局变量,定义在函数外面
var n100 int = 100
var n200 = 200

// 全局变量简写,
var (
	n300 int = 300
	n400     = 400
)

func main() { // 程序的入口
	//打印全局变量
	fmt.Println(n100, n200, n300, n400) // 100 200 300 400

	// 定义在{}中的变量叫局部变量

	// 指定变量的类型,并赋值
	var num int = 18
	fmt.Println("num: ", num) // num:  18

	// 指定变量的类型,但不赋值
	var num2 int
	fmt.Println("num2: ", num2) // num2:  0

	// 没有指定变量的类型,根据=后面的值自动推断变量的类型
	var num3 = 10.1
	fmt.Println("num3: ", num3) // num3:  10.1

	// 省略 var关键字
	num4 := 999
	fmt.Println("num4: ", num4) // num4:  999

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3) // 0 0 0

	var n4, n5, n6 = 10, "LL", 20
	fmt.Println(n4, n5, n6) // 10 LL 20

	n7, n8 := 7.7, 8.8
	fmt.Println(n7, n8) // 7.7 8.8
}
