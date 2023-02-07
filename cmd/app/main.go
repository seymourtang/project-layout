package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/seymourtang/project-layout/internal/injector"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	init, cleanup, err := injector.Build()
	if err != nil {
		log.Fatal(err)
	}
	init.Run()
	<-c
	cleanup()
}
