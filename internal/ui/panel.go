package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func termWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w < 40 {
		return 60
	}
	if w > 120 {
		return 120
	}
	return w - 2
}

func border(w int, left, mid, right string) string {
	return left + strings.Repeat(mid, w-2) + right
}

func colorPrint(style, text string) {
	switch style {
	case "cyan":
		color.Cyan("%s", text)
	case "yellow":
		color.Yellow("%s", text)
	case "green":
		color.Green("%s", text)
	case "red":
		color.Red("%s", text)
	default:
		fmt.Println(text)
	}
}

func Panel(title, content, style string) {
	w := termWidth()

	top := border(w, "┌", "─", "┐")
	titleLine := "│ " + title + strings.Repeat(" ", max(0, w-4-len(title))) + " │"
	mid := border(w, "├", "─", "┤")
	bot := border(w, "└", "─", "┘")

	colorPrint(style, top)
	colorPrint(style, titleLine)
	colorPrint(style, mid)
	fmt.Println(content)
	colorPrint(style, bot)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}