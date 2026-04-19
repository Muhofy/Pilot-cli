package cli

import (
	"os"
	"os/exec"
	"regexp"
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

	sp := ui.NewSpinner("Thinking...")
	sp.Start()
	result, err := ai.Ask(key, cheatsheet.SystemPrompt,
		"Generate ONLY the terminal command (no explanation): "+query)
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

	ui.Panel("Generated Command", cmd, "green")

	check := safety.Check(cmd)
	switch check.Level {
	case safety.Danger:
		ui.Panel("⚠️  DANGEROUS COMMAND", check.Reason, "red")
	case safety.Warning:
		ui.Warning(check.Reason)
	}

	res := ui.Select("Run this command?", []ui.Option{
		{Label: "Yes",  Value: "yes"},
		{Label: "No",   Value: "no"},
		{Label: "Exit", Value: "exit"},
	})

	switch res.Value {
	case "yes":
		execCommand(cmd)
		history.Save("run", query, cmd)
	case "exit":
		os.Exit(0)
	default:
		color.White("Cancelled.")
	}
}

// extractCommand pulls the shell command out of a markdown code block.
// Fixes: previously picked up text outside ``` blocks.
func extractCommand(text string) string {
	// Strategy 1: extract content inside first ``` block
	re := regexp.MustCompile("(?s)```(?:bash|sh|shell|zsh)?\n?(.*?)```")
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		cmd := strings.TrimSpace(matches[1])
		if cmd != "" {
			return cmd
		}
	}

	// Strategy 2: single backtick inline code
	reInline := regexp.MustCompile("`([^`]+)`")
	inlineMatches := reInline.FindStringSubmatch(text)
	if len(inlineMatches) > 1 {
		cmd := strings.TrimSpace(inlineMatches[1])
		if cmd != "" {
			return cmd
		}
	}

	// Strategy 3: first non-empty line that looks like a command
	// (does not start with emoji, #, or plain prose words)
	reCmd := regexp.MustCompile(`^[a-zA-Z$./~_-]`)
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if reCmd.MatchString(line) {
			return line
		}
	}

	return strings.TrimSpace(text)
}

// execCommand runs a shell command in the current terminal session.
func execCommand(cmd string) {
	var c *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// Try PowerShell first, fall back to cmd
		if _, err := exec.LookPath("powershell"); err == nil {
			c = exec.Command("powershell", "-NoProfile", "-Command", cmd)
		} else {
			c = exec.Command("cmd", "/C", cmd)
		}
	default:
		// Use $SHELL if available, fall back to sh
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "sh"
		}
		c = exec.Command(shell, "-c", cmd)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}