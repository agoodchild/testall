package main

import "OpenPlatform/testall/module"

//删除所有表
func main() {
	/*if len(os.Args)<2{
		fmt.Println("请输入要解密的字符串")
	}
	for idx,args:=range os.Args{
		if idx>0{
			bts,err:=base64.StdEncoding.DecodeString(args)
			if err!=nil{
				panic(err)
			}else{
				fmt.Println(string(bts))
			}

		}
	}*/
	//DeleteAllHBaseTable("127.0.0.1:9093")
	//CreateTestData()
	module.QueryData("test","127.0.0.1:9093")
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte("111")))
}
