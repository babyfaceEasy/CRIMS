package server

import (
	"flag"
	"os"

	"github.com/babyfaceeasy/crims/internal/seeds"
)

func HandleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			loadEnv()
			seeds.Execute(args[1:]...)
			os.Exit(0)
		}
	}
}
