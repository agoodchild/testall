package main

import (
	. "OpenPlatform/testall/model"
	. "OpenPlatform/testall/httptools"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func HbaseREST() {
	QueryData("trackinfo_time_id99","9.22.15.25:9093",int64(1572714381),int64(1582626697),1,10)
}

func Base64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func CreateScanner(startTime int64, endTime int64, pageIndex int, pageSize int) Scanner {
	sTime := strconv.FormatInt(startTime, 10)
	eTime := strconv.FormatInt(endTime+1, 10)
	var scanner = Scanner{Batch: 1000000, StartRow: Base64(sTime), EndRow: Base64(eTime)}
	filter := Filter{Type: "FilterList", Op: "MUST_PASS_ALL", Filters: []Filter{{Type: "FirstKeyOnlyFilter"}, {Type: "PageFilter", Value: pageIndex * pageSize}}}
	filter = Filter{Type: "FirstKeyOnlyFilter"}
	scanner.Filter = filter.ToJSON()
	return scanner
}

func Debase64(str string) string {
	bts, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(bts)
}

func GetLocationFromResponse(response *http.Response) (string, error) {
	if response.StatusCode != 201 {
		return "", errors.New("create failed")
	} else {
		location := response.Header.Get("Location")
		return location, nil
	}
}

func QueryData(table string, host string, startTime int64, endTime int64, pageIndex int, pageSize int) QueryResult {
	defer CountTime(time.Now())
	pageData := PageData{PageIndex: pageIndex, PageSize: pageSize}
	result := QueryResult{}
	var scanner = CreateScanner(startTime, endTime, pageIndex, pageSize)
	var url = fmt.Sprintf("http://%s/%s/scanner", host, table)

	response, err := SendRequest(url, "PUT", scanner.ToJSON(), "application/json", "application/json")
	if err != nil {

	}

	location, err := GetLocationFromResponse(response)
	if err != nil {
		panic(err)
	}

	response, err = SendRequest(location, "GET", "", "application/json", "application/json")
	if err != nil {
		panic(err)
	}
	bts, err := ioutil.ReadAll(response.Body)



	var hbData = HBaseData{}
	err = json.Unmarshal(bts, &hbData)
	if err != nil {
		panic(err)
	}


	//--------获取到了全部的结果，计算分页的起始和结束位置------------------------------------------------
	var recordCount = len(hbData.Rows) //总记录数
	if (pageIndex-1)*pageSize >= recordCount {
		pageData.Count = recordCount
		result.Result, result.Msg = pageData, "页码超出范围"
		return result
	}

	//计算待查询页的startKey和endKey
	startIndex := (pageIndex - 1) * pageSize
	endIndex := Min(startIndex+pageSize-1, recordCount-1)
	startKey := hbData.Rows[startIndex].RowKey
	endKey := Debase64(hbData.Rows[endIndex].RowKey)
	endKeyInt, err := strconv.Atoi(endKey)
	endKeyInt += 1
	endKey = Base64(strconv.Itoa(endKeyInt))

	//接下来查询数据
	lastScanner := Scanner{StartRow: startKey, EndRow: endKey, Batch: 1000000}

	response, err = SendRequest(url, "PUT", lastScanner.ToJSON(), "application/json", "application/json")
	if err != nil {
		panic(err)
	}


	location, err = GetLocationFromResponse(response)
	if err != nil {
		panic(err)
	}


	response, err = SendRequest(location, "GET", "", "application/json", "application/json")
	if err != nil {
		panic(err)
	}

	bts, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var hbDataLast = HBaseData{}
	err = json.Unmarshal(bts, &hbDataLast)
	if err != nil {
		panic(err)
	}
	hbDataLast.DecodeBase64()


	var traceArray []TraceMessage
	for _, row := range hbDataLast.Rows {
		fmt.Println(row.RowKey)
		cell := row.GetCell("data:orderdata")
		var orderTrace TraceMessage
		if err = json.Unmarshal([]byte(cell.Value), &orderTrace); err != nil {
			panic(err)
		}
		traceArray = append(traceArray, orderTrace)
	}

	pageData.Count, pageData.Data = recordCount, traceArray
	result.Result = pageData
	return result
}