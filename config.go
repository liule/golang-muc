package main

import (
	"flag"
	"github.com/fogcreek/mini"
)

var C string
var ConfigInfo = NewConfig()

func init() {

}

func NewConfig() *mini.Config {
	flag.StringVar(&C, "f", "", "config file")
	flag.Parse()
	if C == "" {
		panic("usage: server -f config.ini")
	}
	if conf, err := mini.LoadConfiguration(C); err != nil {
		panic(err.Error())
	} else {
		return conf
	}
}
