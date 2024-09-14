package config

import (
    "flag"
)

var ControlAddr string
var VisitAddr string
var TunnelAddr string
var Mode string
var RemoteIP string
var LocalServerAddr string
var RemoteServerAddr string
var RemoteControlAddr string
var Checkip string

func InitFlags() {

    flag.StringVar(&ControlAddr, "controladdr", "", "定义controlAddr")
    flag.StringVar(&TunnelAddr, "tunneladdr", "", "定义tunnelAddr")
    flag.StringVar(&VisitAddr, "visitaddr", "", "定义visitAddr")
    flag.StringVar(&Mode, "mode", "server", "定义mode")
    flag.StringVar(&LocalServerAddr, "localserveraddr", "127.0.0.1:8080", "定义localserveraddr")
    flag.StringVar(&RemoteControlAddr, "remotecontroladdr", "127.0.0.1:8009", "定义remotecontroladdr")
    flag.StringVar(&RemoteServerAddr, "remoteserveraddr", "127.0.0.1:8008", "定义remoteserveraddr")
    flag.StringVar(&Checkip, "checkip", "ture", "定义checkip")
    
}