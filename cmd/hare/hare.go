package main

import (
	"github.com/captainpete/hare"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var hare = &hare.Hare{}
	err := hare.Start()
	if err != nil {
		log.Println(err)
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	hare.Stop()
}
