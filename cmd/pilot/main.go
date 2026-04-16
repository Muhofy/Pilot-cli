package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/cli"
	"github.com/muhofy/pilot/internal/config"
	"github.com/muhofy/pilot/internal/locale"
)

func main() {
	cfg := config.Load()
	locale.Init(cfg.Lang)
	ai.SetPreferredModel(cfg.Model)

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
	case "config":
		cli.Config(args)
	case "setup":
		cli.Setup()
	case "completion":
		cli.Completion(args)
	default:
		color.Red(locale.T("err_unknown_cmd"), sub)
		cli.Usage()
	}
}