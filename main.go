package main

import "OpenPlatform/testall/module"

//删除所有表
func main() {
	//DeleteAllHBaseTable("127.0.0.1:9093")
	//CreateTestData()
	module.QueryData("127.0.0.1:9093","test")
}
