//go:build windows

package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Select on Windows uses numbered menu (raw terminal not supported).
func Select(question string, options []Option) SelectResult {
	colorQuestion := color.New(color.FgYellow, color.Bold)
	colorSelected := color.New(color.FgCyan, color.Bold)

	colorQuestion.Printf("\n  ? %s\n", question)
	for i, opt := range options {
		colorSelected.Printf("  %d) %s\n", i+1, opt.Label)
	}
	fmt.Print("\n  Enter number: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := strings.TrimSpace(sc.Text())

	for i, opt := range options {
		if input == fmt.Sprintf("%d", i+1) {
			fmt.Println()
			return SelectResult{Index: i, Value: opt.Value, Label: opt.Label}
		}
	}

	fmt.Println()
	return SelectResult{Index: -1}
}