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

func CreateBigTestData(host string ){
	//deviceid,distance,endtime,orderid,starttime
	var traceData=`{"apikey":"OSPBZ-MOE3X-4A64S-ZZNXQ-QNQB2-4VBG7","deviceid":"%s","distance":%s,"duration":13,"endtime":%s,"orderid":"%s","starttime":%s,"time_res":[{"distance":0,"duration":13,"end":"2020-03-03 08:56:00","start":"2020-03-03 08:55:47","time_end":"10:00","time_start":"07:00","zone_end":"2020-03-03 10:00:00","zone_start":"2020-03-03 07:00:00"}],"track":[{"altitude":5,"bearing":20,"bind_bearing":30,"bind_lat":40.039,"bind_lng":116.806,"latitude":40.039,"loctime":1583196956,"longitude":116.806,"radius":15,"speed":15}],"track_opt":[{"altitude":5,"bearing":30,"bind_bearing":0,"bind_lat":0,"bind_lng":0,"latitude":0,"loctime":1583196947,"longitude":0,"radius":15,"speed":15}]}`
	trackData:=`[{"altitude":5,"bearing":20,"bind_bearing":30,"bind_lat":40.039,"bind_lng":116.806,"latitude":40.039,"loctime":1583196956,"longitude":116.806,"radius":15,"speed":15}]`
	trackOptData:=`[{"altitude":5,"bearing":30,"bind_bearing":0,"bind_lat":0,"bind_lng":0,"latitude":0,"loctime":1583196947,"longitude":0,"radius":15,"speed":15}]`
	orderIndex:=1
	//循环100个设备，每个设备100个轨迹数据
	for i:=0;i<100;i++{
		deviceID:="id"+strconv.Itoa(i+1)
		hbase.CreateTable("trackinfo_order_"+deviceID,host)
		hbase.CreateTable("trackinfo_time_"+deviceID,host)
		hbdOrderKey:=HBaseData{Table:"trackinfo_order_"+deviceID}
		hbdTimeKey:=HBaseData{Table:"trackinfo_time_"+deviceID}
		//每个设备100条数据
		for j:=0;j<100;j++{
			orderID:="oid"+strconv.Itoa(orderIndex)
			distanceFloat := float64(RangeRand(10, 1000)) / 10.0
			startTime := RangeRand(1572533622, 1582633622)
			endTime := startTime + 10000

			jsonData:=fmt.Sprintf(traceData,deviceID,strconv.FormatFloat(distanceFloat,'f', -1, 64),strconv.FormatInt(endTime,10),orderID,strconv.FormatInt(startTime,10))

			hbRowOrderKey:=HBaseRow{
				RowKey: orderID,
				Cells: []HBaseCell{
					{Column:"id:orderid",Value:orderID},
					{Column:"id:deviceid",Value:deviceID},
					{Column:"info:starttime",Value: strconv.FormatInt(startTime,10)},
					{Column:"info:endtime",Value: strconv.FormatInt(endTime,10),},
					{Column:"data:orderdata",Value:jsonData},
					{Column:"data:trackorigin",Value:trackData},
					{Column:"data:trackoptimized",Value:trackOptData},
				},
			}
			hbRowTimeKey:=HBaseRow{
				RowKey: strconv.FormatInt(startTime,10),
				Cells: []HBaseCell{
					{Column:"id:orderid",Value:orderID},
					{Column:"id:deviceid",Value:deviceID},
					{Column:"info:starttime",Value: strconv.FormatInt(startTime,10)},
					{Column:"info:endtime",Value: strconv.FormatInt(endTime,10),},
					{Column:"data:orderdata",Value:jsonData},
					{Column:"data:trackorigin",Value:trackData},
					{Column:"data:trackoptimized",Value:trackOptData},
				},
			}
			hbdOrderKey.Rows=append(hbdOrderKey.Rows,hbRowOrderKey)
			hbdTimeKey.Rows=append(hbdTimeKey.Rows,hbRowTimeKey)
			orderIndex+=1
		}
		hbase.SaveHBaseDataToHBase([]HBaseData{hbdOrderKey,hbdTimeKey},host)
	}
	fmt.Println("创建测试数据成功")
}


