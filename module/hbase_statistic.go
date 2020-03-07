package module

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hbase"
	"net"
	"strconv"
	"sync"
)

var thriftClient *hbase.THBaseServiceClient

var clientInitOnce sync.Once

func GetClient()*hbase.THBaseServiceClient {
	clientInitOnce.Do(func() {
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		transport, err := thrift.NewTSocket(net.JoinHostPort("9.22.15.25","9090"))
		if err != nil {
			fmt.Println(err)
		}
		client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
		if err := transport.Open(); err != nil {
			fmt.Println(err)
		}
		thriftClient=client
	})
	return thriftClient
}


func Statistic(){
	client:=GetClient()
	tColumn:=hbase.TColumn{
		Family:    []byte("info"),
		Qualifier: []byte("starttime"),
	}
	batchSize:=int32(10000)
	tScan:=hbase.TScan{
		Columns:            []*hbase.TColumn{&tColumn},
		BatchSize:          &batchSize,
	}
	id:=1
	for ;id<100;id++{
		tableName:="trackinfo_time_id"+strconv.Itoa(id)
		tResults,err:=client.GetScannerResults(context.TODO(),[]byte(tableName),&tScan,10000)
		if err!=nil{
			panic(err)
		}
		minTime:=string(tResults[0].Row)
		maxTime:=string(tResults[len(tResults)].Row)
		fmt.Println(tableName+","+minTime+","+maxTime)
	}
}