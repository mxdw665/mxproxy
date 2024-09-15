# mxproxy
# mxproxy 配置包文档

## 概述

`mxproxy` 是一个使用golang实现的tcp流量转发工具，提供了客户端和服务器模式。

## 启动变量

- `controladdr` (string): 控制服务器地址。
- `visitaddr` (string): 访问地址。
- `tunneladdr` (string): 隧道连接地址。
- `mode` (string): 应用运行模式。默认值是 "server"。
- `remoteip` (string): 远程 ip 地址（当前代码中未使用）。
- `localserveraddr` (string): 本地服务器地址。默认值是 "127.0.0.1:8080"。
- `remoteserveraddr` (string): 远程服务器地址。默认值是 "127.0.0.1:8008"。
- `remotecontroladdr` (string): 远程控制地址。默认值是 "127.0.0.1:8009"。
- `checkip` (string): 是否检查 ip 地址。默认值是 "true"

## 示例使用

运行程序时，可以使用以下命令行参数进行自定义配置：

### 'server'
```sh
./mxproxy_OS_ARCH -mode="server" -controladdr="0.0.0.0:8009" -tunneladdr="0.0.0.0:8008" -visitaddr="0.0.0.0:8007"
```

### 'client'
```sh
./mxproxy_OS_ARCH -mode="client" -remoteserveraddr="127.0.0.1:8008" -remotecontroladdr="127.0.0.1:8009" -localserveraddr="127.0.0.1:8080"
```