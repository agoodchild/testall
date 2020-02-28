package module

import (
	"OpenPlatform/testall/hbase"
	. "OpenPlatform/testall/model"
	"fmt"
)

func CreateTestData() {
	table:="test1"
	host:="127.0.0.1:9093"
	//删除表
	//hbase.DeleteTable(table,host)
	//创建表
//	_=hbase.CreateTable(table,host)
	//添加数据
	hbd:= HBaseData{
		Table:table,
		Rows: []HBaseRow{
			{
				RowKey:"111",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"111"},
					{Column:"info:starttime",Value:"1582633622"},
					{Column:"info:endtime",Value:"1582720022"},
					{Column:"info:jsondata",Value:"data111"},
				},
			},
			{
				RowKey:"222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"222"},
					{Column:"info:starttime",Value:"1582547222"},
					{Column:"info:endtime",Value:"1582633622"},
					{Column:"info:jsondata",Value:"data222"},
				},
			},
			{
				RowKey:"333",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"333"},
					{Column:"info:starttime",Value:"1582447222"},
					{Column:"info:endtime",Value:"1582533622"},
					{Column:"info:jsondata",Value:"data333"},
				},
			},
			{
				RowKey:"444",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"444"},
					{Column:"info:starttime",Value:"1582347222"},
					{Column:"info:endtime",Value:"1582433622"},
					{Column:"info:jsondata",Value:"data444"},
				},
			},
			{
				RowKey:"555",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
		},
	}
	hbd= HBaseData{
		Table:table,
		Rows: []HBaseRow{
			{
				RowKey:"1582633622",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"111"},
					{Column:"info:starttime",Value:"1582633622"},
					{Column:"info:endtime",Value:"1582720022"},
					{Column:"info:jsondata",Value:"data111"},
				},
			},
			{
				RowKey:"1582547222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"222"},
					{Column:"info:starttime",Value:"1582547222"},
					{Column:"info:endtime",Value:"1582633622"},
					{Column:"info:jsondata",Value:"data222"},
				},
			},
			{
				RowKey:"1582447222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"333"},
					{Column:"info:starttime",Value:"1582447222"},
					{Column:"info:endtime",Value:"1582533622"},
					{Column:"info:jsondata",Value:"data333"},
				},
			},
			{
				RowKey:"1582347222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"444"},
					{Column:"info:starttime",Value:"1582347222"},
					{Column:"info:endtime",Value:"1582433622"},
					{Column:"info:jsondata",Value:"data444"},
				},
			},
			{
				RowKey:"1582247222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
		},
	}

	hbase.SaveHBaseDataToHBase([]HBaseData{hbd},host)
	fmt.Println("创建测试数据成功")
}

