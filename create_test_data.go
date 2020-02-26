package main

import (
	"OpenPlatform/testall/hbase"
	. "OpenPlatform/testall/model"
	"fmt"
)

func CreateTestData() {
	table:="test1"
	host:="9.134.193.247:9093"
	//删除表
	//hbase.DeleteTable(table,host)
	//创建表
	_=hbase.CreateTable(table,host)
	//添加数据
	hbd:= HBaseData{
		Table:table,
		Rows: []HBaseRow{
			{
				RowKey:"1582633622",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"111111111"},
					{Column:"info:starttime",Value:"1582633622"},
					{Column:"info:endtime",Value:"1582720022"},
					{Column:"info:jsondata",Value:"2020/2/25 20:27:2-2020/2/26 20:27:2"},
				},
			},
			{
				RowKey:"1582547222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"22222222222222"},
					{Column:"info:starttime",Value:"1582547222"},
					{Column:"info:endtime",Value:"1582633622"},
					{Column:"info:jsondata",Value:"2020/2/24 20:27:2-2020/2/25 20:27:2"},
				},
			},
			{
				RowKey:"1582447222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"333333333333"},
					{Column:"info:starttime",Value:"1582447222"},
					{Column:"info:endtime",Value:"1582533622"},
					{Column:"info:jsondata",Value:"2020/2/23 20:27:2-2020/2/24 20:27:2"},
				},
			},
			{
				RowKey:"1582347222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"444444444444"},
					{Column:"info:starttime",Value:"1582347222"},
					{Column:"info:endtime",Value:"1582433622"},
					{Column:"info:jsondata",Value:"2020/2/22 20:27:2-2020/2/23 20:27:2"},
				},
			},
			{
				RowKey:"1582247222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"444444444444"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"2020/2/21 20:27:2-2020/2/22 20:27:2"},
				},
			},
		},
	}
	hbase.SaveHBaseDataToHBase([]HBaseData{hbd},host)
	fmt.Println("创建测试数据成功")
}

