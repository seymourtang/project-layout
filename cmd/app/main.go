package main

import (
	"os"
	"os/signal"

	_ "go.uber.org/automaxprocs"

	"github.com/seymourtang/project-layout/internal/injector"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	init, cleanup, err := injector.Build()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	init.Run()
	<-c
}
