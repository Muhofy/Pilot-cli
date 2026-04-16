package cli

import (
	"fmt"

	"github.com/muhofy/pilot/internal/ui"
)

// Completion prints the shell completion script for the requested shell.
func Completion(args []string) {
	if len(args) == 0 {
		ui.Error("Usage: pilot completion <bash|zsh|fish>")
		return
	}

	switch args[0] {
	case "bash":
		fmt.Print(bashCompletion)
	case "zsh":
		fmt.Print(zshCompletion)
	case "fish":
		fmt.Print(fishCompletion)
	default:
		ui.Error("Unsupported shell: " + args[0])
		fmt.Println("  Supported: bash, zsh, fish")
	}
}

// ── Bash ─────────────────────────────────────────────────────

const bashCompletion = `# pilot bash completion
# Add to ~/.bashrc:
#   eval "$(pilot completion bash)"

_pilot_completion() {
  local cur prev words
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"
  words="${COMP_WORDS[*]}"

  local subcommands="ask explain run history config setup completion"

  case "$prev" in
    pilot)
      COMPREPLY=( $(compgen -W "${subcommands}" -- "${cur}") )
      return 0
      ;;
    history)
      COMPREPLY=( $(compgen -W "search clear" -- "${cur}") )
      return 0
      ;;
    config)
      COMPREPLY=( $(compgen -W "set show" -- "${cur}") )
      return 0
      ;;
    set)
      COMPREPLY=( $(compgen -W "lang model" -- "${cur}") )
      return 0
      ;;
    lang)
      COMPREPLY=( $(compgen -W "en tr" -- "${cur}") )
      return 0
      ;;
    model)
      local models="deepseek/deepseek-chat-v3.1:free meta-llama/llama-4-maverick:free qwen/qwen3-235b-a22b:free google/gemma-3-27b-it:free"
      COMPREPLY=( $(compgen -W "${models}" -- "${cur}") )
      return 0
      ;;
    completion)
      COMPREPLY=( $(compgen -W "bash zsh fish" -- "${cur}") )
      return 0
      ;;
  esac
}

complete -F _pilot_completion pilot
`

// ── Zsh ──────────────────────────────────────────────────────

const zshCompletion = `#compdef pilot
# pilot zsh completion
# Add to ~/.zshrc:
#   eval "$(pilot completion zsh)"

_pilot() {
  local state

  _arguments \
    '1: :->subcommand' \
    '*: :->args'

  case $state in
    subcommand)
      local subcommands=(
        'ask:Generate a terminal command from natural language'
        'explain:Explain what a command does'
        'run:Generate and execute a command'
        'history:Show or search command history'
        'config:View or update configuration'
        'setup:API key setup guide'
        'completion:Print shell completion script'
      )
      _describe 'subcommand' subcommands
      ;;

    args)
      case ${words[2]} in
        history)
          local history_cmds=(
            'search:Search history by keyword'
            'clear:Clear all history'
          )
          _describe 'history command' history_cmds
          ;;

        config)
          case ${words[3]} in
            set)
              case ${words[4]} in
                lang)
                  local langs=('en:English' 'tr:Turkish')
                  _describe 'language' langs
                  ;;
                model)
                  local models=(
                    'deepseek/deepseek-chat-v3.1:free'
                    'meta-llama/llama-4-maverick:free'
                    'qwen/qwen3-235b-a22b:free'
                    'google/gemma-3-27b-it:free'
                  )
                  _values 'model' $models
                  ;;
                *)
                  local keys=('lang:UI language' 'model:AI model')
                  _describe 'config key' keys
                  ;;
              esac
              ;;
            *)
              local config_cmds=('set:Set a config value' 'show:Show current config')
              _describe 'config command' config_cmds
              ;;
          esac
          ;;

        completion)
          local shells=('bash:Bash completion' 'zsh:Zsh completion' 'fish:Fish completion')
          _describe 'shell' shells
          ;;
      esac
      ;;
  esac
}

_pilot
`

// ── Fish ─────────────────────────────────────────────────────

const fishCompletion = `# pilot fish completion
# Add to ~/.config/fish/config.fish:
#   pilot completion fish | source

# Disable file completion for pilot
complete -c pilot -f

# Subcommands
complete -c pilot -n '__fish_use_subcommand' -a ask        -d 'Generate a command from natural language'
complete -c pilot -n '__fish_use_subcommand' -a explain    -d 'Explain what a command does'
complete -c pilot -n '__fish_use_subcommand' -a run        -d 'Generate and execute a command'
complete -c pilot -n '__fish_use_subcommand' -a history    -d 'Show or search command history'
complete -c pilot -n '__fish_use_subcommand' -a config     -d 'View or update configuration'
complete -c pilot -n '__fish_use_subcommand' -a setup      -d 'API key setup guide'
complete -c pilot -n '__fish_use_subcommand' -a completion -d 'Print shell completion script'

# history subcommands
complete -c pilot -n '__fish_seen_subcommand_from history' -a search -d 'Search history by keyword'
complete -c pilot -n '__fish_seen_subcommand_from history' -a clear  -d 'Clear all history'

# config subcommands
complete -c pilot -n '__fish_seen_subcommand_from config' -a set  -d 'Set a config value'
complete -c pilot -n '__fish_seen_subcommand_from config' -a show -d 'Show current config'

# config set keys
complete -c pilot -n '__fish_seen_subcommand_from set' -a lang  -d 'UI language'
complete -c pilot -n '__fish_seen_subcommand_from set' -a model -d 'AI model'

# lang values
complete -c pilot -n '__fish_seen_subcommand_from lang' -a en -d 'English'
complete -c pilot -n '__fish_seen_subcommand_from lang' -a tr -d 'Turkish'

# model values
complete -c pilot -n '__fish_seen_subcommand_from model' -a 'deepseek/deepseek-chat-v3.1:free'    -d 'DeepSeek v3.1'
complete -c pilot -n '__fish_seen_subcommand_from model' -a 'meta-llama/llama-4-maverick:free'    -d 'Llama 4 Maverick'
complete -c pilot -n '__fish_seen_subcommand_from model' -a 'qwen/qwen3-235b-a22b:free'           -d 'Qwen3 235B'
complete -c pilot -n '__fish_seen_subcommand_from model' -a 'google/gemma-3-27b-it:free'          -d 'Gemma 3 27B'

# completion shells
complete -c pilot -n '__fish_seen_subcommand_from completion' -a bash -d 'Bash'
complete -c pilot -n '__fish_seen_subcommand_from completion' -a zsh  -d 'Zsh'
complete -c pilot -n '__fish_seen_subcommand_from completion' -a fish -d 'Fish'
`