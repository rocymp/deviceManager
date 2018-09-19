package main

import (
	"log"

	"github.com/rocymp/deviceManager/tserver/app/tcp"
)

func main() {
	log.Println("Start Device Manger Service")
	dm := tcp.InitDMHandler("127.0.0.1:8888", 3)

	// go dm.Tick()
	dm.Run()
}
