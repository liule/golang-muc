package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type TcpConnection struct {
	*net.TCPConn
	ReadBuf  []byte //读取的byte数组
	WriteBuf []byte //写入的byte数组
	timeout  time.Duration
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
	tcpConnection.head = make([]byte, 8)
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

// 一次读取的字节数,循环读取
func (this *TcpConnection) Read(length int) ([]byte, error) {
	if len(this.ReadBuf) > length {
		err := this.RealRead(this.ReadBuf, length)
		return this.ReadBuf, err
	}
	data := make([]byte, length)
	err := this.RealRead(data, length)
	return data, err
}

func (this *TcpConnection) RealRead(data []byte, length int) error {
	// 设置读写超时时间
	this.TCPConn.SetReadDeadline(time.Now().Add(this.timeout))
	n := 0
}

func (this *TcpConnection) Write(stream []byte) error {

}
