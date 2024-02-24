package main

import (
	"log"

	"github.com/kynrai/nhshackday-24/server"
)

func main() {
	conf := server.NewConfig()
	srv, err := server.NewServer(conf)
	if err != nil {
		log.Fatal(err)
	}
	srv.Start()
}
