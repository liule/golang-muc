package main

import ()

func init() {
	SetConsole(true)
	SetRollingFile("./log", "server.log", 10, 500, MB)
	SetRollingDaily("./log", "server.log")
	SetLevel(DEBUG)
}

func main() {
	go Listen()
	select {}
}
