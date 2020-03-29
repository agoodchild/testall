package dbaccess

import (
	"OpenPlatform/CloudMapService/log"
	. "OpenPlatform/CloudMapService/model"
	"context"
	"fmt"
	"github.com/go-pg/pg"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

//使用github.com/go-pg/pg
var db *pg.DB

var openOnce sync.Once

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	log.DBAccess(fmt.Sprintf(q.FormattedQuery()))
	return nil
}

var host = "9.134.193.247"
var user = "postgres"
var password = "jl3406JLB"
var database = "geocloud"

//获取数据库连接
func GetDB() *pg.DB {
	openOnce.Do(func() {
		db = pg.Connect(&pg.Options{
			Addr:    host,
			User:     user,
			Password: password,
			Database: database,
		})
		//db.AddQueryHook(dbLogger{})
	})
	return db
}

//验证Map是否存在
func IsMapExists(mapId string) bool {
	var gMap = GeoMap{ID: mapId}
	exists, err := GetDB().Model(&gMap).WherePK().Exists()
	if err != nil {
		log.Error(err)
		return false
	}
	return exists
}

//验证Layer是否存在
func IsLayerExists(layerId string) bool {
	var gLayer = GeoLayer{ID: layerId}
	exists, err := GetDB().Model(&gLayer).WherePK().Exists()
	if err != nil {
		log.Error(err)
		return false
	}
	return exists
}

//判断表是否存在
func IsTableExists(tableId string) bool {
	if !IsValidTableId(tableId) {
		return false
	}
	var tables []GeoTable
	exists, err := GetDB().Model(&tables).Where("id =?", tableId).Exists()
	if err != nil {
		log.Error(err)
		return false
	}
	return exists
}

//验证TableID
func IsValidTableId(tableID string) bool {
	if tableID == "" {
		return false
	}
	pattern := `^[0-9a-zA-Z]+`
	result, _ := regexp.MatchString(pattern, tableID)
	return result
}

//根据tableID获取table
func GetTable(tableID string) *GeoTable {
	var table = GeoTable{ID: tableID}
	err := GetDB().Model(&table).WherePK().Relation("Fields").Select()
	if err != nil {
		log.Error(err)
		return nil
	}
	return &table
}

//根据地图ID获取地图
func GetMap(mapID string) *GeoMap {
	var gMap = GeoMap{ID: mapID}
	err := GetDB().Model(&gMap).WherePK().Select()
	if err != nil {
		log.Error(err)
		return nil
	}
	return &gMap
}

//查询某个表的自定义字段
func GetColumns(tableID string) []GeoField {
	var fields []GeoField
	err := GetDB().Model(&fields).Where("table_id=?", tableID).Select()
	if err != nil {
		log.Error(err)
		return nil
	}
	return fields
}

//判断是否系统保留字段
func IsSystemField(fieldName string) bool {
	fieldName = strings.ToLower(fieldName)
	if fieldName == "id" || fieldName == "ud_id" || fieldName == "title" || fieldName == "address" ||
		fieldName == "province" || fieldName == "city" || fieldName == "district" || fieldName == "adcode" ||
		fieldName == "tel" || fieldName == "location" || fieldName == "geometry" || fieldName == "task_id" ||
		fieldName == "lng" || fieldName == "lat" {
		return true
	}
	return false
}

func GetValidMapName(name string) string {
	var count int = 0
	sqlCount := "select count(1) from geo_maps where map_name='" + name + "'"
	_, err := GetDB().Query(&count, sqlCount)
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return name
	}
	var maxNumber int = 0
	maxSql := fmt.Sprintf(`select max(cast(substring(map_name from '%s_#"[0-9]*#"' for '#') as integer ))   from geo_maps  where map_name SIMILAR TO '%s_[0-9]*'`, name, name)
	_, err = GetDB().QueryOne(&maxNumber, maxSql)
	if err != nil {
		panic(err)
	}
	return name + "_" + strconv.Itoa(maxNumber+1)
}
