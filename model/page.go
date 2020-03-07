package model

type ErrInfo struct {
    Status int `json:"status"`
    Msg    string  `json:"message"`
}

type QueryResult struct{
    ErrInfo
    Result interface{} `json:"result"`
}

type PageData struct {
    PageIndex int `json:"pageindex"`
    PageSize  int `json:"pagesize"`
    Count     int         `json:"count"`
    Data      interface{} `json:"data"`
}

