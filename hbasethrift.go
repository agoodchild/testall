package main

import (
	. "OpenPlatform/testall/model"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hbase"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	HOST       = "9.22.15.25"
	PORT       = "9090"
	TESTRECORD = 10
)

func HbaseThrift2() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
/*	IsRowExists(client)
	GetRowAndPrint(client)*/
	ScanData(client)
}

func HbaseThrift3() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/*	IsRowExists(client)
		GetRowAndPrint(client)*/
	ScanData3(client)
}

func ScanData(client *hbase.THBaseServiceClient) {
	Query(client,"trackinfo_time_id99",int64(1572714381),int64(1582626697),1,10)
}
func ScanData3(client *hbase.THBaseServiceClient) {
	Query3(client,"trackinfo_time_id99",int64(1572714381),int64(1582626697),1,10)
}

func Query(client *hbase.THBaseServiceClient, table string, startTime int64, endTime int64, pageIndex int, pageSize int) QueryResult{
	defer CountTime(time.Now())
	result := QueryResult{}
	pageData := PageData{PageIndex: pageIndex, PageSize: pageSize}

	startRow:=strconv.FormatInt(startTime,10)
	endRow:=strconv.FormatInt(endTime+1,10)
	var tColumnId = hbase.TColumn{Family: []byte("id"),Qualifier: []byte("id"),}
	var tColumns= []*hbase.TColumn{ &tColumnId}
	var batchSize = int32(10240)
	var tScan = hbase.TScan{ StartRow:[]byte(startRow), StopRow:[]byte(endRow),Columns:tColumns,BatchSize: &batchSize,MaxVersions: 0,}
	scannerID,err:=client.OpenScanner(context.Background(),[]byte(table),&tScan)
	if err!=nil{
		panic(err)
	}
	rows,err:=client.GetScannerRows(context.Background(),scannerID,10240)
	if err!=nil{
		panic(err)
	}
	var recordCount =len(rows)//总记录数
	if (pageIndex-1)*pageSize >= recordCount {
		pageData.Count = recordCount
		result.Result,result.Msg = pageData,"页码超出范围"
		return result
	}
	startIndex,endIndex := (pageIndex - 1) * pageSize,Min(pageIndex*pageSize-1, recordCount-1)
	startRow,endRow=string(rows[startIndex].Row),string(rows[endIndex].Row)
	endRowInt, err := strconv.Atoi(endRow)
	endRow =strconv.Itoa(endRowInt+1)

	tColumnData:=hbase.TColumn{Family:[]byte("info"),Qualifier: []byte("jsondata"),}
	tColumns=append(tColumns,&tColumnData)
	tScan = hbase.TScan{ StartRow:[]byte(startRow), StopRow:[]byte(endRow),Columns:tColumns,BatchSize: &batchSize,MaxVersions: 0,}
	scannerID,err=client.OpenScanner(context.Background(),[]byte(table),&tScan)
	if err!=nil{
		panic(err)
	}
	fmt.Println(scannerID)
	rows,err=client.GetScannerRows(context.Background(),scannerID,10240)
	if err!=nil{
		panic(err)
	}
	fmt.Println(len(rows))
	for _,row:=range rows{
		fmt.Println(string(row.Row))
	}
	return result
}

func Query3(client *hbase.THBaseServiceClient, table string, startTime int64, endTime int64, pageIndex int, pageSize int) QueryResult{
	defer CountTime(time.Now())
	result := QueryResult{}
	pageData := PageData{PageIndex: pageIndex, PageSize: pageSize}

	startRow:=strconv.FormatInt(startTime,10)
	endRow:=strconv.FormatInt(endTime+1,10)
	var tColumnId = hbase.TColumn{Family: []byte("id"),Qualifier: []byte("id"),}
	var tColumns= []*hbase.TColumn{ &tColumnId}
	var batchSize = int32(10240)
	var tScan = hbase.TScan{ StartRow:[]byte(startRow), StopRow:[]byte(endRow),Columns:tColumns,BatchSize: &batchSize,MaxVersions: 0,}
	resultArr,err:= client.GetScannerResults(context.Background(),[]byte(table),&tScan,10240)

	var recordCount =len(resultArr)//总记录数
	if (pageIndex-1)*pageSize >= recordCount {
		pageData.Count = recordCount
		result.Result,result.Msg = pageData,"页码超出范围"
		return result
	}
	startIndex,endIndex := (pageIndex - 1) * pageSize,Min(pageIndex*pageSize-1, recordCount-1)
	startRow,endRow=string(resultArr[startIndex].Row),string(resultArr[endIndex].Row)
	endRowInt, err := strconv.Atoi(endRow)
	endRow =strconv.Itoa(endRowInt+1)

	tColumnData:=hbase.TColumn{Family:[]byte("info"),Qualifier: []byte("jsondata"),}
	tColumns=append(tColumns,&tColumnData)
	tScan = hbase.TScan{ StartRow:[]byte(startRow), StopRow:[]byte(endRow),Columns:tColumns,BatchSize: &batchSize,MaxVersions: 0,}

	resultArr,err=client.GetScannerResults(context.Background(),[]byte(table),&tScan,10240)
	if err!=nil{
		panic(err)
	}
	fmt.Println(len(resultArr))
	for _,row:=range resultArr{
		fmt.Println(string(row.Row))
	}
	return result
}

func Min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func IsRowExists(client *hbase.THBaseServiceClient) {
	defer CountTime(time.Now())
	exists, err := client.Exists(context.Background(), []byte("trackinfo_time_id99"), &hbase.TGet{Row: []byte("1582595459")})
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)
}

func GetRowAndPrint(client *hbase.THBaseServiceClient) {
	defer CountTime(time.Now())
	tres, err := client.Get(context.Background(), []byte("trackinfo_time_id99"), &hbase.TGet{Row: []byte("1582595459")})
	if err != nil {
		panic(err)
	}
	for _, columnValue := range tres.ColumnValues {
		fmt.Println(string(columnValue.Qualifier) + ":" + string(columnValue.Value))
	}
}

func CountTime(t1 time.Time) {
	fmt.Println("耗时:" + strconv.FormatInt(time.Since(t1).Milliseconds(), 10))
}
