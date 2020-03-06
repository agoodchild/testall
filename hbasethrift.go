package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hbase"
	"net"
	"os"
)

const (
	HOST       = "127.0.0.1"
	PORT       = "9090"
	TESTRECORD = 10
)

//删除所有表
func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	bool, err := client.Exists(context.Background(), []byte("trackinfo_time_id99"), &hbase.TGet{Row: []byte("1582595459")})
	fmt.Println(bool)

	tres,err:=client.Get(context.Background(),[]byte("trackinfo_time_id99"), &hbase.TGet{Row: []byte("1582595459")})
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(tres.Row))
	var batchsize =int32(1000)
	scan:=hbase.TScan{
		StartRow:           []byte("1582595459"),
		StopRow:           []byte("1582626697"),
		Columns:            nil,
		Caching:            nil,
		MaxVersions:        1,
		TimeRange:          nil,
		FilterString:       nil,
		BatchSize:          &batchsize,
		Attributes:         nil,
		Authorizations:     nil,
		Reversed:           nil,
		CacheBlocks:        nil,
		ColFamTimeRangeMap: nil,
		ReadType:           nil,
		Limit:              nil,
		Consistency:        nil,
		TargetReplicaId:    nil,
		FilterBytes:        nil,
	}
	scan=hbase.TScan{
		BatchSize:          &batchsize,
		Limit:              &batchsize,
	}
	scannerid,err:=client.OpenScanner(context.Background(),[]byte("trackinfo_time_id99"),&scan)
	if err!=nil{
		panic(err)
	}
	resarr,err:=client.GetScannerResults(context.Background(),[]byte("trackinfo_time_id99"),&scan,100)
/*	resarr,err=client.GetScannerRows(context.Background(),scannerid,100)
	if err!=nil{
		panic(err)
	}*/
	fmt.Println(scannerid)
	fmt.Println(len(resarr))
	for _,res:=range resarr{
		fmt.Println(string(res.Row))
	}
}
