package main

import (
	"log"
	"time"
)

type Server struct {
	tcpConnection    *TcpConnection
	msgChan          chan *DocyData
	cmdChan          chan *DocyData
	rQuit            chan bool
	wQuit            chan bool
	heartBeatTimeout time.Duration
}

func NewServer(tcpConnection *TcpConnection, msgLen int) *Server {
	return &Server{
		msgChan:       make(chan *DocyData, msgLen),
		rQuit:         make(chan bool, 1),
		wQuit:         make(chan bool, 1),
		tcpConnection: tcpConnection,
	}
}

// 读协程
func (this *Server) ServerRead() {
	for {
		docyData := NewDocyData()
		if err := docyData.Parser(this.tcpConnection, this.heartBeatTimeout); err != nil {
			log.Fatal("ServerRead error:", err.Error())
			goto failed
		}
		this.cmdChan <- docyData
	}
failed:
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
				log.Fatal("ServerWrite error:", err.Error())
				goto failed
			}
		case data := <-this.cmdChan:
			this.Process(data)
		case <-this.wQuit:
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
