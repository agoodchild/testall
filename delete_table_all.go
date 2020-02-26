package main

import (
	"OpenPlatform/testall/hbase"
)

//删除所有表
func DeleteAllHBaseTable(host string) {
	hbase.DeleteAllTables(host)
}
