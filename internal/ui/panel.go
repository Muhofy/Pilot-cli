package ui

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const panelWidth = 60

func printBorder(style, pos string) {
	line := strings.Repeat("─", panelWidth)
	chars := map[string]map[string]string{
		"top":    {"cyan": "┌%s┐", "yellow": "┌%s┐", "green": "┌%s┐", "red": "┌%s┐"},
		"mid":    {"cyan": "├%s┤", "yellow": "├%s┤", "green": "├%s┤", "red": "├%s┤"},
		"bottom": {"cyan": "└%s┘", "yellow": "└%s┘", "green": "└%s┘", "red": "└%s┘"},
	}
	format := chars[pos][style]
	switch style {
	case "cyan":
		color.Cyan(format, line)
	case "yellow":
		color.Yellow(format, line)
	case "green":
		color.Green(format, line)
	case "red":
		color.Red(format, line)
	}
}

func Panel(title, content, style string) {
	printBorder(style, "top")
	switch style {
	case "cyan":
		color.Cyan("│ %-*s │", panelWidth-2, title)
	case "yellow":
		color.Yellow("│ %-*s │", panelWidth-2, title)
	case "green":
		color.Green("│ %-*s │", panelWidth-2, title)
	case "red":
		color.Red("│ %-*s │", panelWidth-2, title)
	}
	printBorder(style, "mid")
	fmt.Println(content)
	printBorder(style, "bottom")
}

func Loading(msg string) {
	color.White("⏳ %s", msg)
}

func Error(msg string) {
	color.Red("❌ %s\n", msg)
}

func Success(msg string) {
	color.Green("✅ %s\n", msg)
}

func Warning(msg string) {
	color.Yellow("⚠️  %s\n", msg)
}