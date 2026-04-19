package cli

import "github.com/fatih/color"

// Usage prints the help message.
func Usage() {
	color.Cyan(`
  pilot — Your terminal co-pilot

  Usage:
    pilot ask        <request>         → natural language → generate command
    pilot explain    <command>         → explain a command
    pilot run        <request>         → generate + confirm + execute
    pilot chain      <steps>           → chain multiple steps into one command
    pilot chain      --dry <steps>     → preview chain without executing
    pilot update                       → check for updates and install latest
    pilot history                      → show recent queries
    pilot history    search <keyword>  → search history
    pilot history    clear             → clear all history
    pilot config     set model <model> → set preferred AI model
    pilot config     set lang  <lang>  → set UI language (en, tr)
    pilot config     show              → display current config
    pilot completion <bash|zsh|fish>   → print shell completion script
    pilot setup                        → API key setup guide
`)
}

// Setup prints the API key setup instructions.
func Setup() {
	color.Cyan(`
  pilot — Setup
  ─────────────
  1. Get a free API key: https://openrouter.ai/keys

  2. Export it:
     export OPENROUTER_API_KEY=your_key_here

  3. Persist it:
     echo 'export OPENROUTER_API_KEY=your_key' >> ~/.zshrc
     source ~/.zshrc
`)
}