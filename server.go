package main

import (
	"time"
)

type Server struct {
	fd               int64
	msgChan          chan interface{}
	cmdChan          chan interface{}
	quit             chan bool
	user             string
	heartBeatTimeout time.Duration
}
