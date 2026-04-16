package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// Option represents a single selectable item.
type Option struct {
	Label  string
	Value  string
}

// SelectResult is the outcome of a Select call.
type SelectResult struct {
	Index  int
	Value  string
	Label  string
}

var (
	colorSelected = color.New(color.FgCyan, color.Bold)
	colorNormal   = color.New(color.FgWhite)
	colorQuestion = color.New(color.FgYellow, color.Bold)
)

// Select renders an arrow-key interactive prompt and returns the chosen option.
// Returns index -1 if the user presses Ctrl+C or Escape.
func Select(question string, options []Option) SelectResult {
	// Switch terminal to raw mode
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		// Fallback to plain confirm if raw mode unavailable
		return fallbackSelect(question, options)
	}
	defer term.Restore(fd, oldState)

	cursor := 0
	n := len(options)

	// Hide cursor
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	render := func() {
		// Clear previously rendered lines
		// Move up n+1 lines and clear each
		if cursor >= 0 {
			fmt.Printf("\033[%dA", n+1)
		}
		for i := 0; i < n+1; i++ {
			fmt.Print("\033[2K\r")
			if i < n {
				fmt.Println()
			}
		}
		fmt.Printf("\033[%dA", n)

		// Question line
		fmt.Print("\033[2K\r")
		colorQuestion.Printf("? %s\n", question)

		// Options
		for i, opt := range options {
			fmt.Print("\033[2K\r")
			if i == cursor {
				colorSelected.Printf("  ❯ %s\n", opt.Label)
			} else {
				colorNormal.Printf("    %s\n", opt.Label)
			}
		}
	}

	// Initial render (no clear on first draw)
	colorQuestion.Printf("? %s\n", question)
	for i, opt := range options {
		if i == cursor {
			colorSelected.Printf("  ❯ %s\n", opt.Label)
		} else {
			colorNormal.Printf("    %s\n", opt.Label)
		}
	}

	buf := make([]byte, 3)
	for {
		os.Stdin.Read(buf)

		switch {
		// Arrow Up / k
		case buf[0] == 27 && buf[1] == 91 && buf[2] == 65, buf[0] == 'k':
			if cursor > 0 {
				cursor--
			} else {
				cursor = n - 1 // wrap to bottom
			}
			render()

		// Arrow Down / j
		case buf[0] == 27 && buf[1] == 91 && buf[2] == 66, buf[0] == 'j':
			if cursor < n-1 {
				cursor++
			} else {
				cursor = 0 // wrap to top
			}
			render()

		// Enter
		case buf[0] == 13, buf[0] == 10:
			fmt.Println()
			return SelectResult{
				Index: cursor,
				Value: options[cursor].Value,
				Label: options[cursor].Label,
			}

		// Ctrl+C or Escape
		case buf[0] == 3, buf[0] == 27 && buf[1] == 0:
			fmt.Println()
			return SelectResult{Index: -1}
		}

		// Reset buffer
		buf = make([]byte, 3)
	}
}

// fallbackSelect is a simple numbered fallback when raw mode is unavailable.
func fallbackSelect(question string, options []Option) SelectResult {
	colorQuestion.Printf("? %s\n", question)
	for i, opt := range options {
		colorNormal.Printf("  %d) %s\n", i+1, opt.Label)
	}
	var input int
	fmt.Scan(&input)
	if input < 1 || input > len(options) {
		return SelectResult{Index: -1}
	}
	return SelectResult{
		Index: input - 1,
		Value: options[input-1].Value,
		Label: options[input-1].Label,
	}
}