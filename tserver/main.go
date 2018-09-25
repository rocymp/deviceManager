package main

import (
	"log"

	"github.com/rocymp/deviceManager/tserver/app/http"
	"github.com/rocymp/deviceManager/tserver/app/tcp"
)

func main() {
	log.Println("Start Device Manger Service")
	dm := tcp.InitDMHandler("0.0.0.0:18989", 3)

	// go dm.Tick()
	go dm.Run()

	http.Init(dm)
}

// SET CGO_ENABLED=0
// SET GOOS=linux
// SET GOARCH=amd64
