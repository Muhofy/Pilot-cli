package cli

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/safety"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
)

// Run generates a command, shows a safety check, and executes it after confirmation.
func Run(args []string) {
	if len(args) == 0 {
		ui.Error("Usage: pilot run <what do you want to do?>")
		return
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")
	ui.Loading("Generating command...")

	result, err := ai.Ask(key, cheatsheet.SystemPrompt,
		"Generate ONLY the terminal command (no explanation): "+query)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	cmd := extractCommand(result)
	ui.Panel("Generated Command", cmd, "green")

	// Safety check
	check := safety.Check(cmd)
	switch check.Level {
	case safety.Danger:
		ui.Panel("⚠️  DANGEROUS COMMAND", check.Reason, "red")
	case safety.Warning:
		ui.Warning(check.Reason)
	}

	// Arrow-key confirmation
	res := ui.Select("Run this command?", []ui.Option{
		{Label: "Yes",  Value: "yes"},
		{Label: "No",   Value: "no"},
		{Label: "Exit", Value: "exit"},
	})

	switch res.Value {
	case "yes":
		execCommand(cmd)
		history.Save("run", query, cmd)
	case "no":
		color.White("Cancelled.")
	case "exit":
		color.White("Exiting.")
		os.Exit(0)
	default:
		color.White("Cancelled.")
	}
}

// extractCommand pulls the shell command out of a markdown code block.
func extractCommand(text string) string {
	lines := strings.Split(text, "\n")
	var cmd []string
	inBlock := false
	for _, l := range lines {
		if strings.HasPrefix(l, "```") {
			inBlock = !inBlock
			continue
		}
		if inBlock {
			cmd = append(cmd, l)
		}
	}
	if len(cmd) > 0 {
		return strings.TrimSpace(strings.Join(cmd, "\n"))
	}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" && !strings.HasPrefix(l, "📌") {
			return l
		}
	}
	return strings.TrimSpace(text)
}

// execCommand runs a shell command in the current terminal session.
func execCommand(cmd string) {
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", cmd)
	} else {
		c = exec.Command("sh", "-c", cmd)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}