package hbase

import (
	"OpenPlatform/testall/model"
	"OpenPlatform/testall/httptools"
	"fmt"
	"strings"
)

func SaveHBaseDataToHBase(traceDataArray []model.HBaseData,host string) {
	for _, traceData := range traceDataArray {
		strData, err := traceData.ToJSON(true)
		if err!=nil{
			panic(err)
		}
		url := fmt.Sprintf("http://%s/%s/row_key", host, traceData.Table)
		code, response, err := httptools.SendRequest(url, "PUT", strData, "application/json", "application/json")
		if err != nil {
			panic(err)
		}
		if code != 200 {
			fmt.Println(code)
			fmt.Println(response)
		}
	}
}

func DeleteAllTables(host string) {
	code,res,_:= httptools.SendRequest("http://"+host+"/","GET","","text/plain","text/plain")
	if code==200{
		arr:=strings.Split(res, "\n")
		for _,tbName:=range arr{
			_,_,_= httptools.SendRequest("http://"+host+"/"+tbName+"/schema","DELETE","","text/plain","text/plain")
		}
	}
	fmt.Println("删除所有表成功")
}

func DeleteTable(tableName string,host string){
	_,_,_= httptools.SendRequest("http://"+host+"/"+tableName+"/schema","DELETE","","text/plain","text/plain")
}

func CreateTable(tableName string,host string)error{
	html := `<TableSchema name="tableschema"><ColumnSchema name="id"/><ColumnSchema name="info"/><ColumnSchema name="data"/></TableSchema>`
	url := fmt.Sprintf("http://%s/%s/schema", host, tableName)
	_, _, err:= httptools.SendRequest(url, "PUT", html, "text/xml", "text/xml")
	fmt.Println("Create Table :'"+tableName+"'")
	return err
}
