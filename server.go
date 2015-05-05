package main

import (
	"net"
	"time"
)

type Server struct {
	tcpConnection    *TcpConnection
	msgChan          chan *DocyData
	cmdChan          chan *DocyData
	rQuit            chan bool
	wQuit            chan bool
	heartBeatTimeout time.Duration
	remoteIp         string
	remotePort       int
}

func NewServer(tcpConnection *TcpConnection, msgLen int) *Server {
	remoteAddr := tcpConnection.RemoteAddr().(*net.TCPAddr)
	println("NewServer", remoteAddr.IP.String(), remoteAddr.Port)
	return &Server{
		msgChan:          make(chan *DocyData, msgLen),
		cmdChan:          make(chan *DocyData, msgLen),
		rQuit:            make(chan bool, 1),
		wQuit:            make(chan bool, 1),
		tcpConnection:    tcpConnection,
		heartBeatTimeout: time.Duration(20 * time.Second),
		remoteIp:         remoteAddr.IP.String(),
		remotePort:       remoteAddr.Port,
	}
}

// 读协程
func (this *Server) ServerRead() {
	for {
		docyData := NewDocyData()
		if err := docyData.Parser(this.tcpConnection, this.heartBeatTimeout); err != nil {
			println("ServerRead error:", err.Error())
			goto failed
		}
		this.cmdChan <- docyData
		println("ServerRead", docyData.GetVersion())
	}
failed:
	println("this.wQuit <- false")
	this.wQuit <- false
}

// 写协程
func (this *Server) Loop() {
	go this.ServerRead()
	for {
		select {
		case data := <-this.msgChan:
			err := this.tcpConnection.Write(data.ConvertToStream())
			if err != nil {
				Error("ServerWrite error:", err.Error())
				goto failed
			}
		case data := <-this.cmdChan:
			this.Process(data)
		case <-this.wQuit:
			println("Loop:this.wQuit")
			return
		}
	}
failed:
	this.rQuit <- false
	return
}

// 处理命令
func (this *Server) Process(docyData *DocyData) error {
	//协议处理函数
	return nil
}

func Listen() {
	tcpServer, err := NewTcpServer(ConfigInfo.String("listen", "5678"))
	if err != nil {
		panic(err.Error())
	}
	for {
		tcpConnection, err := tcpServer.Accept()
		if err != nil {
			Error("main:Accept", err.Error())
			continue
		}
		server := NewServer(tcpConnection, 50)
		go server.Loop()
		println("Listen")
	}
}
