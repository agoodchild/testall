package module

import (
	"OpenPlatform/testall/hbase"
	. "OpenPlatform/testall/model"
	"fmt"
	"strconv"
)

func CreateTestData() {
	table:="test1"
	host:="127.0.0.1:9093"
	//删除表
	//hbase.DeleteTable(table,host)
	//创建表
	_=hbase.CreateTable(table,host)
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
			{
				RowKey:"1582147222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1582047222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581947222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581847222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581747222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581647222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581547222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581447222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581347222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581247222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581147222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1581047222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1580947222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1580847222",
				Cells: []HBaseCell{
					{Column:"id:id",Value:"555"},
					{Column:"info:starttime",Value:"1582247222"},
					{Column:"info:endtime",Value:"1582333622"},
					{Column:"info:jsondata",Value:"data555"},
				},
			},
			{
				RowKey:"1580747222",
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

func CreateBigTestData(table string,host string ,count int){
	hbd:= HBaseData{
		Table:table,
	}
	var traceData=`{"orderid":"%s","deviceid":"device_key_value","api_key":"api_key_value","starttime":1582602213,"endtime":1582602213,"distance":0,"duration":12,"opt":0,"time_res":[{"time_start":"17:00","time_end":"20:00","zone_start":"2020-02-24 17:00:00","zone_end":"2020-02-24 20:00:00","start":"2020-02-24 17:48:27","end":"2020-02-24 17:48:39","distance":0,"duration":12}],"track":[{"loctime":1582537715,"bind_lng":116.805898,"bind_lat":40.038987,"bind_bearing":40.038987,"altitude":40.038987,"radius":40.038987,"speed":40.038987,"bearing":40.038987}],"track_opt":[{"loctime":1582537715,"bind_lng":116.805898,"bind_lat":40.038987,"bind_bearing":40.038987,"altitude":40.038987,"radius":40.038987,"speed":40.038987,"bearing":40.038987}]}`
	max:=1582633622
	var rows []HBaseRow
	for i:=0;i<count;i++{
		key:=strconv.Itoa(max-i)
		td:=fmt.Sprintf(traceData,strconv.Itoa(i))
		row:=HBaseRow{
			RowKey:key,
			Cells: []HBaseCell{
				{Column:"id:id",Value:key},
				{Column:"info:starttime",Value:"1582633622"},
				{Column:"info:endtime",Value:"1582720022"},
				{Column:"info:jsondata",Value:td},
			},
		}
		rows=append(rows,row)
	}
	hbd.Rows=rows
	hbase.SaveHBaseDataToHBase([]HBaseData{hbd},host)
	fmt.Println("创建测试数据成功")
}

