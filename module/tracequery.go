package module

import (
	"OpenPlatform/testall/httptools"
	"OpenPlatform/testall/model"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type PageResult struct {
	PageIndex int `json:"pageindex"`
	PageSize  int `json:"pagesize"`
	Count     int         `json:"count"`
	Data      interface{} `json:"data"`
}
type Scanner struct {
	Batch    int    `json:"batch,omitempty"`
	StartRow string `json:"startRow,omitempty"` //包含
	EndRow   string `json:"endRow,omitempty"`   //不包含
	Filter   string `json:"filter,omitempty"`
	Column   string `json:"column,omitempty"`
}
type Filter struct {
	Type       string      `json:"type,omitempty"`
	Op         string      `json:"op,omitempty"`
	Comparator *Comparator `json:"comparator,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Filters    []Filter    `json:"filters,omitempty"`
}
type Comparator struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
func (filter Filter) ToJSON() string {
	bts, err := json.Marshal(filter)
	if err != nil {
		panic(err)
	}
	return string(bts)
}
func (scanner Scanner) ToJSON() string {
	bts, err := json.Marshal(scanner)
	if err != nil {
		panic(err)
	}
	return string(bts)
}
func(res PageResult)ToJSON()string{
	bts, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(bts)
}
func Base64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}
func Debase64(str string)string{
	bts,err:=base64.StdEncoding.DecodeString(str)
	if err!=nil{
		panic(err)
	}
	return string(bts)
}
func LogTime(ti time.Time) {
	fmt.Println(time.Since(ti).Milliseconds())
}
func Log(msg interface{}) {
	fmt.Println(msg)
}
func GetLocationFromResponse(response *http.Response)(string,error){
	if response.StatusCode != 201 {
		return "", errors.New("create failed")
	} else {
		location := response.Header.Get("Location")
		return location ,nil
	}
}
func QueryData(table string, host string, startTime int64, endTime int64, pageIndex int, pageSize int) {
	defer LogTime(time.Now())
	result:=PageResult{PageIndex:pageIndex,PageSize:pageSize}
	var scanner =CreateScanner(startTime,endTime,pageIndex,pageSize)
	var url = fmt.Sprintf("http://%s/%s/scanner", host, table)
	response, err:= httptools.SendRequest(url, "PUT", scanner.ToJSON(), "application/json", "application/json")
	location,err:=GetLocationFromResponse(response)
	response, err = httptools.SendRequest(location, "GET", "", "application/json", "application/json")
	bts, err := ioutil.ReadAll(response.Body)
	var hbData = model.HBaseData{}
	err = json.Unmarshal(bts, &hbData)
	
	//--------获取到了全部的结果，计算分页的起始和结束位置------------------------------------------------
	var recordCount= len(hbData.Rows) //总记录数
	if (pageIndex-1)*pageSize>=recordCount{
		fmt.Println("页码超出范围")
		result.Count=recordCount
		fmt.Println(result.ToJSON())
		return
	}
	startIndex:=(pageIndex-1)*pageSize
	endIndex:=Min(startIndex+pageSize-1,recordCount-1)
	startKey:=hbData.Rows[startIndex].RowKey
	endKey:=Debase64(hbData.Rows[endIndex].RowKey)
	endKeyInt,err:=strconv.Atoi(endKey)
	endKeyInt+=1
	endKey=Base64(strconv.Itoa(endKeyInt))
	//计算出了startKey和endKey，接下来查询数据
	lastScanner:=Scanner{StartRow:startKey,EndRow:endKey,Batch:1000000}
	response, err = httptools.SendRequest(url, "PUT", lastScanner.ToJSON(), "application/json", "application/json")
	location,err=GetLocationFromResponse(response)
	fmt.Println(location)
	response, err = httptools.SendRequest(location, "GET", "", "application/json", "application/json")
	bts, err = ioutil.ReadAll(response.Body)
	var hbDataLast = model.HBaseData{}



	err = json.Unmarshal(bts, &hbDataLast)
	hbDataLast.DecodeBase64()
	var tss [] model.TraceMessage
	for _,row:=range hbDataLast.Rows{
		cell:= row.GetCell("info:jsondata")
		var ts model.TraceMessage
		json.Unmarshal([]byte(cell.Value),&ts)
		tss=append(tss,ts)
	}
	pagerRes:=PageResult{Count:recordCount,PageIndex:pageIndex,PageSize:pageSize,Data:tss}
	bts,_=json.Marshal(pagerRes)
	fmt.Println(string(bts))
	Log(err)
}
func Min(a int ,b int)int{
	if a<=b{
		return a
	}else{
		return b
	}
}
func CreateScanner(startTime int64,endTime int64,pageIndex int,pageSize int)Scanner{
	sTime := strconv.FormatInt(startTime, 10)
	eTime := strconv.FormatInt(endTime+1, 10)
	var scanner = Scanner{Batch: 1000000, StartRow: Base64(sTime), EndRow: Base64(eTime)}
	filter:=Filter{Type:"FilterList",Op:"MUST_PASS_ALL",Filters:[]Filter{{Type: "FirstKeyOnlyFilter"}, {Type: "PageFilter", Value: pageIndex*pageSize}}}
	filter=Filter{Type: "FirstKeyOnlyFilter"}
	scanner.Filter = filter.ToJSON()
	return scanner
}