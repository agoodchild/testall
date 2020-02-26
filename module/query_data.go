package module

import (
    "OpenPlatform/testall/httptools"
    "OpenPlatform/testall/model"
    "encoding/json"
    "fmt"
    "io/ioutil"
)

func QueryData(table string,host string){
    var url=fmt.Sprintf("http://%s/%s/scanner",host,table)
    var contentType="application/json"
    var data=`{"batch":100}`
    fmt.Println(url)
    response,err:=httptools.SendRequest(url,"PUT",data,contentType,contentType)
    if err!=nil{
        panic(err)
    }
    fmt.Println(response.StatusCode)
    location:=response.Header.Get("Location")
    fmt.Println(location)
    fmt.Println("---------------------")
    response,err=httptools.SendRequest(location,"GET","","application/json","application/json")
    if err!=nil{
        panic(err)
    }
    fmt.Println(response.StatusCode)
    bts,err:=ioutil.ReadAll(response.Body)
    if err!=nil{
        panic(err)
    }
    var hbData =model.HBaseData{}

    err=json.Unmarshal(bts,&hbData)
    hbData.DecodeBase64()
    if err!=nil{
        panic(err)
    }
    fmt.Println(hbData)
}