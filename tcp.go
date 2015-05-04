package main

import (
	"fmt"
	"log"
	"net"
)

type TcpConnection struct {
	*net.TCPConn
	ReadBuf  []byte
	WriteBuf []byte
}

type TcpServer struct {
	*net.TCPListener
	tcpAddr string
}

func NewTcpServer(address string) (*TcpServer, error) {
	tcpServer := &TcpServer{
		tcpAddr: address,
	}
	tcpListener, err := net.ListenTCP("tcp", tcpServer.tcpAddr)
	if err != nil {
		return nil, err
	}
	tcpServer.TCPListener = tcpListener
	return tcpServer, nil

}

func (this *TcpServer) Accept() *TcpConnection {
	tcpConn, err := this.TCPListener.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tcpConnection := &TcpConnection{TCPConn: tcpConn}
}

// 获取连接ip
func (this *TcpConnection) RemoteIp() string {
	remoteAddr := this.RemoteAddr().(*net.TCPAddr)
	return remoteAddr.IP.String()
}

// 获取连接port
func (this *TcpConnection) RemotePort() string {
	remoteAddr := this.RemoteAddr().(*net.TCPAddr)
	return strconv.Itoa(remoteAddr.Port)
}
