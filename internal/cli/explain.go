package cli

import (
	"strings"

	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
)

// Explain describes what a given terminal command does.
func Explain(args []string) {
	if len(args) == 0 {
		ui.Error("Usage: pilot explain <command>")
		return
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")

	sp := ui.NewSpinner("Thinking...")
	sp.Start()
	result, err := ai.Ask(key, cheatsheet.SystemPrompt, "Explain this command: "+query)
	sp.Stop()

	if err != nil {
		ui.Error(err.Error())
		return
	}

	ui.Panel("pilot explain", result, "yellow")
	history.Save("explain", query, result)
}