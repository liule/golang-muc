package main

import (
	"time"
)

type Server struct {
	tcpConnection    *TcpConnection
	fd               int64
	msgChan          chan *DocyData
	cmdChan          chan *DocyData
	quit             chan bool
	user             string
	heartBeatTimeout time.Duration
}

func NewServer(msgLen int) *Server {
	return &Server{
		msgChan: make(chan *DocyData, msgLen),
		quit:    make(chan bool, 1),
	}
}
