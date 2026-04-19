package cli

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/safety"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
)

const chainSystemPrompt = `You are pilot, a terminal assistant.
The user will describe multiple steps they want to perform.
You MUST combine them into a single shell command using && or | operators.

STRICT OUTPUT FORMAT:
` + "```" + `
command1 && command2 && command3
` + "```" + `
📌 What it does: one sentence per step, joined with →

Rules:
- Use && to run commands sequentially (stop on error)
- Use | to pipe output between commands
- Use ; only when subsequent commands must run regardless
- NEVER write prose, explanations, or multiple code blocks
- Add ⚠️ prefix in the 📌 line if any step is destructive
- If steps cannot be chained, explain why in one sentence`

// Chain generates a multi-step chained command from natural language.
func Chain(args []string) {
	if len(args) == 0 {
		ui.Error("Usage: pilot chain [--dry] <step 1, step 2, ...>")
		color.Yellow("  Example: pilot chain \"create branch dev, add all files, commit 'init'\"")
		return
	}

	// --dry flag
	dry := false
	if args[0] == "--dry" {
		dry = true
		args = args[1:]
		if len(args) == 0 {
			ui.Error("Usage: pilot chain --dry <steps>")
			return
		}
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")

	sp := ui.NewSpinner("Building command chain...")
	sp.Start()
	result, err := ai.Ask(key, cheatsheet.SystemPrompt+"\n\n"+chainSystemPrompt,
		"Chain these steps into one command: "+query)
	sp.Stop()

	if err != nil {
		ui.Error(err.Error())
		return
	}

	cmd := extractCommand(result)
	if cmd == "" {
		ui.Error("Could not extract a command from the response.")
		return
	}

	// Show full result panel
	ui.Panel("pilot chain", result, "cyan")

	if dry {
		color.Yellow("\n  --dry mode: command not executed.\n")
		return
	}

	// Safety check on the full chain
	check := safety.Check(cmd)
	switch check.Level {
	case safety.Danger:
		ui.Panel("⚠️  DANGEROUS COMMAND", check.Reason, "red")
	case safety.Warning:
		ui.Warning(check.Reason)
	}

	res := ui.Select("Run this chain?", []ui.Option{
		{Label: "Yes, run all steps", Value: "yes"},
		{Label: "No, cancel",         Value: "no"},
		{Label: "Exit",               Value: "exit"},
	})

	switch res.Value {
	case "yes":
		execCommand(cmd)
		history.Save("chain", query, cmd)
	case "exit":
		os.Exit(0)
	default:
		color.White("Cancelled.")
	}
}