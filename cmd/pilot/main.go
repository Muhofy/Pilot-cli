package main

import (
	"os"

	"github.com/muhofy/pilot/internal/cli"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		cli.Usage()
		return
	}

	sub := os.Args[1]
	args := os.Args[2:]

	switch sub {
	case "ask":
		cli.Ask(args)
	case "explain":
		cli.Explain(args)
	case "run":
		cli.Run(args)
	case "history":
		cli.History(args)
	case "setup":
		cli.Setup()
	default:
		color.Red("Unknown command: %s\n", sub)
		cli.Usage()
	}
}