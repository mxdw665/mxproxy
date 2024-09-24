# mxproxy
# mxproxy 配置包文档

## 概述

`mxproxy` 是一个使用golang实现的tcp流量转发工具，提供了客户端和服务器模式。

## 配置文件启动变量

- `controladdr` : 控制服务器地址。
- `visitaddr` : 访问地址。
- `tunneladdr` : 隧道连接地址。
- `mode` : 应用运行模式。("server"或"client")
- `localserveraddr` : 本地服务器地址。
- `remoteserveraddr` : 远程服务器地址。
- `remotecontroladdr` : 远程控制地址。
- `checkip` : 是否检查 ip 地址。("true"或"false")

## 示例使用

在config.yaml修改配置变量

### server
```sh
./mxproxy_OS_ARCH
```

### client
```sh
./mxproxy_OS_ARCH
```

# 问题提交https://github.com/mxdw665/mxproxy/issues