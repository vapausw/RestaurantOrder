package main

import "fmt"

var data MysqlData

func main() {
	//data.Connect()     // 连接数据库
	//defer data.Close() // 用完将数据库关闭
	//token := GenerateSecureToken()
	//fmt.Printf("Generated Token: %s", token)
	id := CreateID()
	fmt.Printf("Generated ID: %s", id)
}
