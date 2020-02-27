package model

import (
	"encoding/base64"
	"encoding/json"
)

type HBaseData struct {
	Table string     `json:"-"`
	Rows  []HBaseRow `json:"Row"`
}

type HBaseCell struct {
	Column string `json:"column"`
	Value  string `json:"$"`
}

type HBaseRow struct {
	RowKey string      `json:"key"`
	Cells  []HBaseCell `json:"Cell"`
}

func (data HBaseData) ToJSON(encrypt bool) (string, error) {
	if encrypt {
		for idxRow, _ := range data.Rows {
			data.Rows[idxRow].RowKey = base64.StdEncoding.EncodeToString([]byte(data.Rows[idxRow].RowKey))
			for idxCell, _ := range data.Rows[idxRow].Cells {
				data.Rows[idxRow].Cells[idxCell].Column = base64.StdEncoding.EncodeToString([]byte(data.Rows[idxRow].Cells[idxCell].Column))
				data.Rows[idxRow].Cells[idxCell].Value = base64.StdEncoding.EncodeToString([]byte(data.Rows[idxRow].Cells[idxCell].Value))
			}
		}
	}
	bts, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

func (data *HBaseData) DecodeBase64() {

	for idxRow, _ := range data.Rows {
		bts, err := base64.StdEncoding.DecodeString(data.Rows[idxRow].RowKey)
		if err != nil {
			panic(err)
		}
		data.Rows[idxRow].RowKey=string(bts)
		for idxCell, _ := range data.Rows[idxRow].Cells {
			bts, err = base64.StdEncoding.DecodeString(data.Rows[idxRow].Cells[idxCell].Column)
			if err != nil {
				panic(err)
			}
			data.Rows[idxRow].Cells[idxCell].Column=string(bts)

			bts, err = base64.StdEncoding.DecodeString(data.Rows[idxRow].Cells[idxCell].Value)
			if err != nil {
				panic(err)
			}
			data.Rows[idxRow].Cells[idxCell].Value=string(bts)
		}
	}
}

func (row HBaseRow)GetCell(column string)*HBaseCell{
	for _,cell:=range row.Cells{
		if cell.Column==column{
			return &cell
		}
	}
	return nil
}