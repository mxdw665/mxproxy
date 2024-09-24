package client

import (
   "bufio"
   "io"
   "log"
   "net"

   "proxy/network"
   "proxy/config"
)

func Main() {
   tcpConn, err := network.CreateTCPConn(config.GetConfig().Client.RemoteControlAddr)
   if err != nil {
      log.Println("[连接失败]" + config.GetConfig().Client.RemoteControlAddr + err.Error())
      return
   }
   log.Println("[已连接]" + config.GetConfig().Client.RemoteControlAddr)

   reader := bufio.NewReader(tcpConn)
   for {
      s, err := reader.ReadString('\n')
      if err != nil || err == io.EOF {
         break
      }

      // 当有新连接信号出现时，新建一个tcp连接
      if s == network.NewConnection+"\n" {
         go connectLocalAndRemote()
      }
   }

   log.Println("[已断开]" + config.GetConfig().Client.RemoteControlAddr)
}

func connectLocalAndRemote() {
   local := connectLocal()
   remote := connectRemote()

   if local != nil && remote != nil {
      network.Join2Conn(local, remote)
   } else {
      if local != nil {
         _ = local.Close()
      }
      if remote != nil {
         _ = remote.Close()
      }
   }
}

func connectLocal() *net.TCPConn {
   conn, err := network.CreateTCPConn(config.GetConfig().Client.LocalServerAddr)
   if err != nil {
      log.Println("[连接本地服务失败]" + err.Error())
   }
   return conn
}

func connectRemote() *net.TCPConn {
   conn, err := network.CreateTCPConn(config.GetConfig().Client.RemoteServerAddr)
   if err != nil {
      log.Println("[连接远端服务失败]" + err.Error())
   }
   return conn
}