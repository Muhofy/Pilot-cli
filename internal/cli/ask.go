package cli

import (
	"strings"

	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
)

// Ask generates a terminal command from a natural language query.
func Ask(args []string) {
	if len(args) == 0 {
		ui.Error("Usage: pilot ask <what do you want to do?>")
		return
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")
	ui.Loading("Thinking...")

	result, err := ai.Ask(key, cheatsheet.SystemPrompt, "Generate a terminal command for: "+query)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	ui.Panel("pilot ask", result, "cyan")
	history.Save("ask", query, result)
}