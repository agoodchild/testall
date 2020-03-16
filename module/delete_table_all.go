package module

import (
	"OpenPlatform/testall/hbase"
)

//删除所有表
func DeleteAllHBaseTable(host string) {
	hbase.DeleteAllTables(host,"trackinfo_time_id")
	hbase.DeleteAllTables(host,"trackinfo_order_id")
}
