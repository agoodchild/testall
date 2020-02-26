package module

import (
    "OpenPlatform/testall/httptools"
    "fmt"
)

func QueryData(table string,host string){
    var url=fmt.Sprintf("http://%s/%s/scanner",host,table)
    var contentType="application/json"
    var data=`{"batch":100}`
    response,err:=httptools.SendRequest(url,"PUT",data,contentType,contentType)
    if err!=nil{
        panic(err)
    }
    location:=response.Header.Get("Location")
    fmt.Println(location)
}