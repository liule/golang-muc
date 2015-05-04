package main

import (
	"fmt"
	"github.com/fogcreek/mini"
	"log"
	"net"
)

func main() {
	tcpServer, err := NewTcpServer(ConfigInfo.String("listen"))
	if err != nil {
		panic(err.Error())
	}
	for {
		tcpConnection, err := tcpServer.Accept()
		if err != nil {
			log.Fatal("main:Accept", err.Error())
			continue
		}
	}
}
