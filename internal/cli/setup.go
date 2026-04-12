package cli

import "github.com/fatih/color"

// Usage prints the help message.
func Usage() {
	color.Cyan(`
  pilot — Your terminal co-pilot

  Usage:
    pilot ask     <request>         → natural language → generate command
    pilot explain <command>         → explain a command
    pilot run     <request>         → generate + confirm + execute
    pilot history                   → show recent queries
    pilot history search <keyword>  → search history
    pilot history clear             → clear all history
    pilot setup                     → API key setup guide
`)
}

// Setup prints the API key setup instructions.
func Setup() {
	color.Cyan(`
  pilot — Setup Guide
  ───────────────────
  1. Go to https://openrouter.ai/keys
  2. Sign up or log in
  3. Click "Create Key"
  4. Export your key:

     export OPENROUTER_API_KEY=your_key_here

  5. To persist across sessions, add to ~/.zshrc or ~/.bashrc:

     echo 'export OPENROUTER_API_KEY=your_key' >> ~/.zshrc
     source ~/.zshrc
`)
}