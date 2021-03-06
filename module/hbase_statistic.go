package module

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hbase"
	"net"
	"strconv"
	"sync"
	"time"
)

var thriftClient *hbase.THBaseServiceClient

var clientInitOnce sync.Once

func GetClient()*hbase.THBaseServiceClient {
	clientInitOnce.Do(func() {
		thrift.NewTJSONProtocolFactory()
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		thrift.NewTSerializerPool(nil)
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
		maxTime:=string(tResults[len(tResults)-1].Row)
		minTimeInt,err:=strconv.ParseInt(minTime,10,64)
		maxTimeInt,err:=strconv.ParseInt(maxTime,10,64)
		fmt.Println(tableName+","+minTime+","+maxTime+","+time.Unix(minTimeInt,0).String()+","+time.Unix(maxTimeInt,0).String())
	}
}