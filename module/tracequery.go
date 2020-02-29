package module

import (
	"OpenPlatform/testall/httptools"
	"OpenPlatform/testall/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

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
func Base64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}
func LogTime(ti time.Time) {
	fmt.Println(time.Since(ti).Milliseconds())
}
func Log(msg interface{}) {
	fmt.Println(msg)
}


func QueryData(table string, host string, startTime int64, endTime int64, pageIndex int, pageSize int) {
	defer LogTime(time.Now())
	sTime := strconv.FormatInt(startTime, 10)
	eTime := strconv.FormatInt(endTime, 10)

	var scanner = Scanner{Batch: 10000024, StartRow: Base64(sTime), EndRow: Base64(eTime)}
	filter:=Filter{Type:"FilterList",Op:"MUST_PASS_ALL",Filters:[]Filter{{Type: "FirstKeyOnlyFilter"}, {Type: "PageFilter", Value: pageIndex*pageSize}}}
	scanner.Filter = filter.ToJSON()
	jsonString := scanner.ToJSON()
	Log(jsonString)

	var url = fmt.Sprintf("http://%s/%s/scanner", host, table)
	var contentType = "application/json"
	response, err := httptools.SendRequest(url, "PUT", jsonString, contentType, contentType)
	if err != nil {
		panic(err)
	}
	fmt.Println("创建Scanner响应结果:" + strconv.Itoa(response.StatusCode))
	if response.StatusCode != 201 {
		bbb, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bbb))
		fmt.Println(response)
	} else {
		location := response.Header.Get("Location")
		fmt.Println("结果读取地址(Location)为:" + location)
		fmt.Println("---------------------")
		response, err = httptools.SendRequest(location, "GET", "", "application/json", "application/json")
		if err != nil {
			panic(err)
		}
		bts, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(bts))

		var hbData = model.HBaseData{}
		err = json.Unmarshal(bts, &hbData)
		//hbData.DecodeBase64()
		if err != nil {
			panic(err)
		}
		fmt.Println(len(hbData.Rows))
		/*for _, row := range hbData.Rows {
			fmt.Print(row.RowKey)
			for _, cell := range row.Cells {
				fmt.Print("[" + cell.Column + "->" + cell.Value + "]")
			}
			fmt.Print("\n")
		}*/

		//开始找最大最小值
	}
}
