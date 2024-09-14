package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    
    "proxy/client"
    "proxy/config"
    "proxy/server"
)

func main() {
    fmt.Println("欢迎使用mxproxy，软件开发中，有bug欢迎去反馈，github: https://github.com/mxdw665/mxproxy")
    config.InitFlags()
    flag.Parse()
    
    if config.Checkip == "true" {
        checkip()
    }
    
    start()
}

func start() {
    switch config.Mode {
    case "server":
        server.Main()
    case "client":
        client.Main()
    default:
        fmt.Println("未知模式")
    }
}

func checkip() {
    response, err := http.Get("https://4.ipw.cn")
    if err != nil {
        fmt.Println("获取 IP 地址失败:", err)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("读取响应体失败:", err)
        return
    }
    myip := string(body)
    fmt.Printf("您的IP为 %s\n", myip)
}