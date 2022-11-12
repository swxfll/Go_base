package dbutils

import "fmt"

// GetConn 首字母小写不能被其它包访问, 首字母小写可以被其它包访问
func GetConn() {
	fmt.Println("执行了dbutils包下的getConn函数")
}
