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

func (data HBaseData) DecodeBase64() {

	for idxRow, _ := range data.Rows {
		bts, err := base64.StdEncoding.DecodeString(data.Rows[idxRow].RowKey)
		if err != nil {
			panic(err)
		}
		data.Rows[idxRow].RowKey = string(bts)
		for idxCell, _ := range data.Rows[idxRow].Cells {
			bts, err = base64.StdEncoding.DecodeString(data.Rows[idxRow].Cells[idxCell].Column)
			if err != nil {
				panic(err)
			}
			data.Rows[idxRow].Cells[idxCell].Column = string(bts)

			bts, err = base64.StdEncoding.DecodeString(data.Rows[idxRow].Cells[idxCell].Value)
			if err != nil {
				panic(err)
			}
			data.Rows[idxRow].Cells[idxCell].Value = string(bts)
		}
	}
}

func (row HBaseRow) GetCell(column string) *HBaseCell {
	for _, cell := range row.Cells {
		if cell.Column == column {
			return &cell
		}
	}
	return nil
}
