package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/seymourtang/project-layout/internal/injector"
)

func main() {
	init, cleanup, err := injector.Build()
	if err != nil {
		panic(err)
	}
	init.Run()
	cleanup()
}
