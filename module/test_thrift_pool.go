package module

import (
	"OpenPlatform/testall/httptools"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

func TestPool(count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			Call()
			wg.Add(-1)
		}()
	}
	wg.Wait()
}

func Call() {
	for i := 0; i < 100; i++ {
		sTime:=time.Now()
		var url = "http://10.204.23.226:10039/trace/v1/history2?page_index=1&endtime=1583828519&starttime=1583828518&deviceid=23ea731e9cdfa190&page_size=10"
		response, err := httptools.SendRequest(url, "GET", "", "application/json", "application/json")
		if err != nil {
			panic(err)
		}

		str, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		if !strings.Contains(string(str), "08ebedff-07c1") {
			fmt.Println("出错了")
		}
		fmt.Println(time.Since(sTime).Milliseconds())
	}
}
