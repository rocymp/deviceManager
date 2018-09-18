package main

import (
	"log"
	"time"

	"github.com/rocymp/zero"
)

func main() {
	csSlice := make([]zero.SocketClient, 0)
	for i := 1; i <= 1000; i++ {
		cs := zero.NewSocketClient("127.0.0.1:18888", 3)
		if cs == nil {
			log.Printf("Index[%d] connect failed\n", i)
			continue
		}

		log.Printf("Index[%d] Online\n", i)
		cs.Online()
		csSlice = append(csSlice, *cs)
	}

	for _, cs := range csSlice {
		cs.SendMessage(24, []byte("hello world"))
	}

	time.Sleep(30 * time.Second)

	for _, cs := range csSlice {
		cs.Stop()
	}
}
