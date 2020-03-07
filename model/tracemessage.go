package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type TraceMessage struct {
	OrderID   string        `json:"orderid"`
	Deviceid  string        `json:"deviceid"`
	APIKey    string        `json:"api_key"`
	StartTime int64         `json:"starttime"`
	EndTime   int64         `json:"endtime"`
	Distance  float64           `json:"distance"`
	Duration  int           `json:"duration"`
	Opt       int           `json:"opt"`
	TimeRes   []TimeResDsc  `json:"time_res"`
	Track     []TrackDsc    `json:"track"`
	TrackOpt  []TrackOptDsc `json:"track_opt"`
}

type TrackDsc struct {
	LocTime     int     `json:"loctime"`
	BindLng     float64 `json:"bind_lng"`
	BindLat     float64 `json:"bind_lat"`
	BindBearing float64 `json:"bind_bearing"`
	Altitude    float64 `json:"altitude"`
	Radius      float64 `json:"radius"`
	Speed       float64 `json:"speed"`
	Bearing     float64 `json:"bearing"`
}

type TrackOptDsc struct {
	LocTime     int     `json:"loctime"`
	BindLng     float64 `json:"bind_lng"`
	BindLat     float64 `json:"bind_lat"`
	BindBearing float64 `json:"bind_bearing"`
	Altitude    float64 `json:"altitude"`
	Radius      float64 `json:"radius"`
	Speed       float64 `json:"speed"`
	Bearing     float64 `json:"bearing"`
}

type TimeResDsc struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	ZoneStart string `json:"zone_start"`
	ZoneEnd   string `json:"zone_end"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Distance  float64    `json:"distance"`
	Duration  int    `json:"duration"`
}

func (tm TraceMessage) ToHBaseDataArray() []HBaseData {
	orderStartTime := time.Unix(tm.StartTime, 0)
	dateStr := orderStartTime.Format("20060102")
	deviceID := tm.Deviceid
	orderID:=tm.OrderID
	jsonAll,_:=json.Marshal(tm)
	jsonTrack,_:=json.Marshal(tm.Track)
	jsonTrackOpt,_:=json.Marshal(tm.TrackOpt)

	tableNameAll := "trackinfo_all_" + dateStr
	tableNameDevice := "trackinfo_" + deviceID

	var hbdAll = HBaseData{Table: tableNameAll}
	var hbdDevice = HBaseData{Table: tableNameDevice}

	var rowAll = HBaseRow{RowKey: orderID, Cells: []HBaseCell{
		{Column: "id:orderid", Value: orderID},
		{Column: "id:deviceid", Value: deviceID},
		{Column: "info:starttime", Value:strconv.FormatInt(tm.StartTime,10)},
		{Column: "info:endtime", Value:strconv.FormatInt(tm.EndTime,10)},
		{Column: "data:orderdata", Value: string(jsonAll)},
		{Column: "data:trackorigin", Value: string(jsonTrack)},
		{Column: "data:trackoptimized", Value: string(jsonTrackOpt)},
	}}
	var rowDevice = HBaseRow{RowKey: orderID, Cells: []HBaseCell{
		{Column: "id:orderid", Value: orderID},
		{Column: "id:deviceid", Value: deviceID},
		{Column: "info:starttime", Value:strconv.FormatInt(tm.StartTime,10)},
		{Column: "info:endtime", Value:strconv.FormatInt(tm.EndTime,10)},
		{Column: "data:orderdata", Value: string(jsonAll)},
		{Column: "data:trackorigin", Value: string(jsonTrack)},
		{Column: "data:trackoptimized", Value: string(jsonTrackOpt)},
	}}
	hbdAll.Rows = append(hbdAll.Rows, rowAll)
	hbdDevice.Rows = append(hbdDevice.Rows, rowDevice)
	return []HBaseData{hbdAll, hbdDevice}
}
