package main

import (
	"log"
)

func main() {
	tcpServer, err := NewTcpServer(ConfigInfo.String("listen", "5678"))
	if err != nil {
		panic(err.Error())
	}
	for {
		tcpConnection, err := tcpServer.Accept()
		if err != nil {
			log.Fatal("main:Accept", err.Error())
			continue
		}
		server := NewServer(tcpConnection, 50)
		go server.Loop()
	}
}
