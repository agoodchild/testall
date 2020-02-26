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
