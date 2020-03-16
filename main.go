package main

import (
	"OpenPlatform/testall/log"
	"OpenPlatform/testall/module"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func IsValidDevieID(deviceid string) bool {
	if deviceid == "" {
		return false
	}
	pattern := `^[0-9a-zA-Z_]+[0-9a-zA-Z_]$`
	result, err := regexp.MatchString(pattern, deviceid)
	if err != nil {
		panic(err)
	}
	return !result
}

//删除所有表
func main() {
	//不输入参数的情况下，走查询
	if len(os.Args) < 2 {
		//module.CreateBigTestData("9.22.15.25:9093")
		//module.InsertOrderTestData()
		//module.Q()
		t1 := time.Unix(111, 0)
		t := time.Unix(222, 0)
		fmt.Println(t1)
		fmt.Println(t)
		fmt.Println("使用方法：")
		fmt.Println("scan：根据filter.xml中的内容扫描数据")
		fmt.Println("createdata：生成20条测试数据")
		fmt.Println("big：生成10万条测试数据")
		fmt.Println("query：分页查询")
	} else {
		operate := os.Args[1]
		if operate == "statistic" {
			module.Statistic()
		}
		if operate == "rest" {
			HbaseREST()
		}
		if operate == "thrift" {
			HbaseThrift2()
		}
		if operate == "thrift3" {
			HbaseThrift3()
		}
		if operate == "scan" {
			module.ScanData("test1", "127.0.0.1:9093")
		}
		if operate == "createdata" {
			module.CreateTestData()
		}
		if operate == "big" {
			//	"9.134.193.247:9093"
			module.CreateBigTestData("9.22.15.25:9093")
			//	module.CreateBigTestData("trackinfo_deviceid1", "127.0.0.1:9093", 100000)
		}
		if operate == "query" {
			pageIndex := 1
			pageSize := 10
			if len(os.Args) > 2 {
				pageIndex, _ = strconv.Atoi(os.Args[2])
			}
			if len(os.Args) > 3 {
				pageSize, _ = strconv.Atoi(os.Args[3])
			}
			startTime := int64(1582623622) // endTime - int64(pageSize)
			endTime := int64(1582633622)
			module.QueryData("test1", "127.0.0.1:9093", startTime, endTime, pageIndex, pageSize)
		}
		if operate == "testtime" {
			t := time.Unix(222, 0)
			fmt.Println(t)
		}
		if operate == "mysql" {
			module.TestMySQL()
		}
		if operate == "pool" {
			if len(os.Args) > 2 {
				count := os.Args[2]
				icount, _ := strconv.Atoi(count)
				module.TestPool(icount)
			} else {
				module.TestPool(20)
			}
		}
		if operate == "log" {
			log.GetLogInstance().Init()
			i, err := strconv.Atoi("12w")
			if err != nil {
				log.Error(err)
			}
			log.Info(i)
			log.Debug("debug")
			log.Warn("warn")
		}
		if operate=="deletetables"{
			module.DeleteAllHBaseTable("9.22.15.25:9093")
		}
	}
}
