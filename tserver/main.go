package main

import (
	"log"

	"github.com/rocymp/tserver/app/tcp"
)

func main() {
	log.Println("Start Device Manger Service")
	dm := tcp.InitDMHandler("127.0.0.1:18888", 3)

	go dm.Tick()
	dm.Run()
}
