package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/ui"
)

// History handles the history subcommands: list, search, clear.
func History(args []string) {
	if len(args) == 0 {
		listHistory()
		return
	}

	switch args[0] {
	case "search":
		if len(args) < 2 {
			ui.Error("Usage: pilot history search <keyword>")
			return
		}
		searchHistory(strings.Join(args[1:], " "))
	case "clear":
		clearHistory()
	default:
		ui.Error("Usage: pilot history [search <keyword>] [clear]")
	}
}

func listHistory() {
	entries, err := history.List(20)
	if err != nil {
		ui.Error("Could not read history: " + err.Error())
		return
	}
	if len(entries) == 0 {
		color.White("No history yet.")
		return
	}

	color.Cyan("━━━ Last %d queries ━━━\n", len(entries))
	for i, e := range entries {
		icon := iconFor(e.Type)
		timeStr := e.Time.Format("02 Jan 15:04")
		color.White("%2d. %s [%s] %s\n", i+1, icon, timeStr, e.Query)
	}
	fmt.Println()
	color.Yellow("pilot history search <keyword>  → search")
	color.Yellow("pilot history clear             → clear all")
}

func searchHistory(keyword string) {
	entries, err := history.Search(keyword)
	if err != nil {
		ui.Error("Search error: " + err.Error())
		return
	}
	if len(entries) == 0 {
		color.White("No results for '%s'.", keyword)
		return
	}

	color.Cyan("━━━ %d result(s) for '%s' ━━━\n", len(entries), keyword)
	for i, e := range entries {
		icon := iconFor(e.Type)
		timeStr := e.Time.Format("02 Jan 15:04")
		color.White("%2d. %s [%s] %s\n", i+1, icon, timeStr, e.Query)
	}
}

func clearHistory() {
	res := ui.Select("All history will be deleted. Are you sure?", []ui.Option{
		{Label: "Yes, clear all", Value: "yes"},
		{Label: "No, cancel",     Value: "no"},
	})

	if res.Value != "yes" {
		color.White("Cancelled.")
		return
	}
	if err := history.Clear(); err != nil {
		ui.Error("Clear error: " + err.Error())
		return
	}
	ui.Success("History cleared.")
}

// iconFor returns an emoji icon for the given command type.
func iconFor(typ string) string {
	switch typ {
	case "ask":
		return "🤔"
	case "explain":
		return "📖"
	case "run":
		return "▶️ "
	case "chain":
		return "⛓️ "
	default:
		return "•"
	}
}