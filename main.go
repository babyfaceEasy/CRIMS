package main

import (
	"github.com/babyfaceeasy/crims/internal/server"
)

func main() {
	server.HandleArgs()

	if err := server.Initialize(); err != nil {
		panic(err)
	}

	server.Run()
}
