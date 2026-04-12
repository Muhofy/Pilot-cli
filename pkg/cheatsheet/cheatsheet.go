package cheatsheet

import (
	_ "embed"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
)

//go:embed cheatsheet.csv
var defaultCSV []byte

// Entry represents a single cheatsheet row.
type Entry struct {
	Category    string
	Command     string
	Description string
}

// Load reads the cheatsheet from ~/.pilot/cheatsheet.csv if it exists,
// otherwise falls back to the embedded default.
func Load() []Entry {
	if custom := loadCustom(); len(custom) > 0 {
		return custom
	}
	return parse(defaultCSV)
}

// loadCustom attempts to read a user-defined cheatsheet from ~/.pilot/cheatsheet.csv.
func loadCustom() []Entry {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	path := filepath.Join(home, ".pilot", "cheatsheet.csv")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return parse(data)
}

// parse decodes CSV bytes into a slice of Entry.
func parse(data []byte) []Entry {
	r := csv.NewReader(strings.NewReader(string(data)))
	r.Comment = '#'
	records, err := r.ReadAll()
	if err != nil {
		return nil
	}

	var entries []Entry
	for _, row := range records {
		// Skip header row and malformed rows
		if len(row) < 3 || row[0] == "category" {
			continue
		}
		entries = append(entries, Entry{
			Category:    strings.TrimSpace(row[0]),
			Command:     strings.TrimSpace(row[1]),
			Description: strings.TrimSpace(row[2]),
		})
	}
	return entries
}

// BuildPrompt converts the cheatsheet entries into a formatted string for the AI system prompt.
func BuildPrompt() string {
	entries := Load()
	var sb strings.Builder

	current := ""
	for _, e := range entries {
		if e.Category != current {
			current = e.Category
			sb.WriteString("\n# " + strings.ToUpper(current) + "\n")
		}
		sb.WriteString(e.Command + "  →  " + e.Description + "\n")
	}
	return sb.String()
}

// SystemPrompt is the full AI system prompt including the cheatsheet.
var SystemPrompt = `You are pilot, a terminal assistant. You ONLY generate or explain terminal/git/docker commands.

STRICT RULES:
1. When generating a command, use ONLY this format:
` + "```" + `
command here
` + "```" + `
📌 What it does: one sentence

2. When explaining a command, use ONLY this format:
🔍 This command: one sentence
📦 Parts:
  - part1: explanation
  - part2: explanation

3. Add ⚠️ for dangerous commands
4. Chain multiple operations with && or |
5. NEVER write stories or unnecessary text
6. If you don't know the command, say: "I don't know this command."

Cheatsheet:
` + BuildPrompt()