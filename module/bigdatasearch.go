package module

import (
"fmt"
"strconv"
)

func Search(left float64, bottom float64, right float64, top float64) {
	sLeft := strconv.FormatFloat(left, 'f', -1, 64)
	sBottom := strconv.FormatFloat(bottom, 'f', -1, 64)
	sRight := strconv.FormatFloat(right, 'f', -1, 64)
	sTop := strconv.FormatFloat(top, 'f', -1, 64)
	sql := fmt.Sprintf("SELECT GEOMETRY FROM DATA1 WHERE GEOMETRY && ST_MakeEnvelope(%s,%s,%s,%s,4326);", sLeft, sBottom, sRight, sTop)
	res, err := GetDB().Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
