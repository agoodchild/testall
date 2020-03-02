package module

import (
	"crypto/rand"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math"
	"math/big"
	"strconv"
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
		dbEngine, err := xorm.NewEngine("mysql", "root:@tcp(106.54.101.237:3306)/trace?charset=utf8")
		if err != nil {
			panic(err)
		}
		db = dbEngine
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

func InsertOrderTestData() {
	var orders []Order
	session := GetDB().Table("t_order")
	for i := 5000; i < 10000; i++ {
		deviceIdInt, _ := rand.Int(rand.Reader, big.NewInt(1000))
		//取一个随机数作为deviceid
		startTime := RangeRand(1582533622, 1582633622)
		endTime := startTime + 10000
		orderID := i + 1
		distanceInt := RangeRand(10, 1000)
		distanceFloat := float64(distanceInt) / 10.0
		order := Order{
			ID:        strconv.Itoa(orderID),
			ServiceID: "1",
			OrderID:   strconv.Itoa(orderID),
			DeviceID:  strconv.FormatInt(deviceIdInt.Int64(), 10),
			StartTime: startTime,
			EndTime:   endTime,
			Distance:  distanceFloat,
		}
		orders = append(orders, order)
		/*if len(orders)>=2000{
			session.Insert(orders)
			clearSlice(&orders)
		}*/
		/*_,err:=GetDB().Table("t_order").Insert(order)
		if err!=nil{
			panic(err)
		}*/
		fmt.Println(i)
	}
	count, err := session.Insert(orders)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	fmt.Println("完成")
}

type DeviceOrderObj struct {
	ServiceID  string  ` xorm:"'serviceid'"`
	DeviceID   string  ` xorm:"'deviceid'"`
	Title      string  ` xorm:"'title'"`
	Type       int     ` xorm:"'type'"`
	Department string  ` xorm:"'department'"`
	User       string  ` xorm:"'user'"`
	Tel        string  ` xorm:"'tel'"`
	Ext        string  ` xorm:"'ext'"`
	Distance   float64 ` xorm:"'distance'"`
	StartTime  int64   ` xorm:"'starttime'"`
	EndTime    int64   ` xorm:"'endtime'"`
}

func Q() {
	startTime := 1582533622
	endTime := 1582543622
	serviceID := "1"
	pageIndex := 1
	pageSize := 10
	objArr:=PageQuery(serviceID,startTime,endTime,pageIndex,pageSize)
	for _,obj:=range objArr{
		fmt.Println(obj.DeviceID+":"+strconv.FormatFloat(obj.Distance,'f', -1, 64))
	}
}
func PageQuery(serviceID string, startTime int, endTime int, pageIndex int, pageSize int) []DeviceOrderObj {
	sql := fmt.Sprintf(`SELECT b.*,a.distance,a.starttime,a.endtime
			FROM (SELECT deviceid, SUM(distance) AS distance,min(starttime) AS starttime,max(starttime) AS endtime
			FROM t_order where starttime>%s and starttime<%s and serviceid=%s GROUP BY deviceid
			ORDER BY distance DESC LIMIT %s,%s) a
			LEFT JOIN (SELECT * from t_device )b  on a.deviceid=b.deviceid`, strconv.Itoa(startTime), strconv.Itoa(endTime),
		serviceID, strconv.Itoa((pageIndex-1)*pageSize), strconv.Itoa(pageSize))
	var objArray []DeviceOrderObj
	ERR := GetDB().SQL(sql).Find(&objArray)
	if ERR != nil {
		panic(ERR)
	}
	return objArray
}
func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func clearSlice(s *[]Order) {
	*s = append([]Order{})
}
