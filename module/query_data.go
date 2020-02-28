package module

import (
	"OpenPlatform/testall/httptools"
	"OpenPlatform/testall/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func ScanData(table string, host string) {
	var url = fmt.Sprintf("http://%s/%s/scanner", host, table)
	//var contentType = "text/xml"
	var contentType = "application/json"
	bts, err := ioutil.ReadFile("./filter/zfilter.xml")
	if err != nil {
		fmt.Println(err)
	}
	var filterFileName = string(bts)
	bts, err = ioutil.ReadFile("./filter/" + filterFileName)
	var data = string(bts)
	fmt.Println("过滤器文件:\n" + filterFileName)
	fmt.Println("过滤器为:\n" + data)
	response, err := httptools.SendRequest(url, "PUT", data, contentType, contentType)
	if err != nil {
		panic(err)
	}
	fmt.Println("创建Scanner响应结果:" + strconv.Itoa(response.StatusCode))
	if response.StatusCode != 201 {
	    bbb,err:=ioutil.ReadAll(response.Body)
	    if err!=nil{
	        panic(err)
        }
        fmt.Println(string(bbb))
		fmt.Println(response)
	} else {
		location := response.Header.Get("Location")
		fmt.Println("结果读取地址(Location)为:" + location)
		fmt.Println("---------------------")
		ShowResult(location)
	}
}
func ShowResult(location string){
	fmt.Println("查询结果为:")
	response, err := httptools.SendRequest(location, "GET", "", "application/json", "application/json")
	if err != nil {
		panic(err)
	}
	bts, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if string(bts)==""{
		fmt.Println("NO CONTENT")
		return
	}
	var hbData = model.HBaseData{}
	err = json.Unmarshal(bts, &hbData)
	hbData.DecodeBase64()
	if err != nil {
		panic(err)
		panic(string(bts))
	}
	for _, row := range hbData.Rows {
		fmt.Print( row.RowKey)
		for _,cell:=range row.Cells{
			fmt.Print("["+cell.Column+"->"+cell.Value+"]")
		}
		fmt.Print("\n")
	}
}

type Comparator struct{
	Type string `json:"type"`
	Value string `json:"value"`
}

type Filter struct{
	Type string `json:"type"`
	Op string `json:"op"`
	Comparator Comparator `json:"comparator"`
}

type Scanner struct{
	Batch int `json:"batch"`
	StartRow string `json:"startRow"`
	EndRow string `json:"endRow"`
	Filter string `json:"filter"`
}
func QueryData(){
	//startTime:=time.Now()
	//endTime:=time.Now()
}