package module

import (
	"crypto/rand"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/big"
	"sync"
)

type Device struct {
	ServiceID  string ` xorm:"'serviceid'"`
	DeviceID   string ` xorm:"'deviceid'"`
	Title      string ` xorm:"'title'"`
	Type       int    ` xorm:"'type'"`
	Department string ` xorm:"'department'"`
	User       string ` xorm:"'user'"`
	Tel        string ` xorm:"'tel'"`
	Ext        string ` xorm:"'ext'"`
}

type Order struct {
	ID        string  `json:"id"  xorm:"'id'"`
	ServiceID string  `json:"serviceid" xorm:"'serviceid'"`
	OrderID   string  `json:"orderid" xorm:"'orderid'"`
	DeviceID  string  `json:"deviceid" xorm:"'deviceid'"`
	StartTime int64   `json:"starttime" xorm:"'starttime'"`
	EndTime   int64   `json:"endtime" xorm:"'endtime'"`
	Distance  float64 `json:"distance" xorm:"'distance'"`
}

var db *xorm.Engine

var openOnce sync.Once

func GetDB() *xorm.Engine {
	openOnce.Do(func() {
		dbEngine,err:= xorm.NewEngine("mysql", "tars:tars2015@tcp(49.232.64.229:6666)/lzh?charset=utf8")
		if err!=nil{
			panic(err)
		}
		db=dbEngine
	})
	return db
}

func TestMySQL() {
	engine, err := xorm.NewEngine("mysql", "tars:tars2015@tcp(9.134.193.247:3306)/trace?charset=utf8")
	if err != nil {
		panic(err)
	}
	tables, err := engine.DBMetas()
	if err != nil {
		panic(err)
	}
	fmt.Println(tables[0])
	order := Order{
		ID:        "id2",
		ServiceID: "service1",
		OrderID:   "order1",
		DeviceID:  "device1",
		StartTime: 1582247222,
		EndTime:   1582257222,
		Distance:  23.623,
	}
	count, err := engine.Table("trace_order").Insert(&order)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func InsertOrderTestData(){
	for i:=0;i<10000;i++{
		deviceIdInt, _:= rand.Int(rand.Reader, big.NewInt(1000))
		//取一个随机数作为deviceid
		rand.
		fmt.Println(n)
	}
}


