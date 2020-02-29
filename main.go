package main

import (
	"OpenPlatform/testall/module"
	"fmt"
	"os"
)

//删除所有表
func main() {
	//不输入参数的情况下，走查询
	if len(os.Args) < 2 {
		fmt.Println("使用方法：")
		fmt.Println("scan：根据filter.xml中的内容扫描数据")
		fmt.Println("createdata：生成20条测试数据")
		fmt.Println("big：生成10万条测试数据")
		fmt.Println("query：分页查询")
	} else {
		operate := os.Args[1]

		if operate == "scan" {
			module.ScanData("test1", "127.0.0.1:9093")
		}
		if operate == "createdata" {
			module.CreateTestData()
		}
		if operate == "big" {
			module.CreateBigTestData("test1", "127.0.0.1:9093", 100000)
		}
		if operate == "query" {
			pageSize := 50000
			startTime :=int64(1582533622)  // endTime - int64(pageSize)
			endTime :=startTime+ int64(pageSize)
			module.QueryData("test1", "127.0.0.1:9093", startTime, endTime, 1, pageSize)
		}
	}
}
