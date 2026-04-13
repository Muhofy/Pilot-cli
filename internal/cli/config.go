package cli

import (
	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/config"
	"github.com/muhofy/pilot/internal/ui"
)

var supportedLangs = map[string]bool{
	"en": true,
	"tr": true,
}

// Config handles the config subcommands: set, show.
func Config(args []string) {
	if len(args) == 0 {
		showConfig()
		return
	}

	switch args[0] {
	case "set":
		if len(args) < 3 {
			ui.Error("Usage: pilot config set <key> <value>")
			color.Yellow("  Keys: lang, model")
			color.Yellow("  e.g.: pilot config set lang tr")
			color.Yellow("  e.g.: pilot config set model deepseek/deepseek-chat-v3.1:free")
			return
		}
		key, value := args[1], args[2]
		if key == "lang" && !supportedLangs[value] {
			ui.Error("Unsupported language: " + value)
			color.Yellow("  Supported: en, tr")
			return
		}
		if err := config.Set(key, value); err != nil {
			ui.Error("Could not save config: " + err.Error())
			return
		}
		ui.Success("Config updated: " + key + " = " + value)

	case "show":
		showConfig()

	default:
		ui.Error("Usage: pilot config [set <key> <value>] [show]")
	}
}

func showConfig() {
	cfg := config.Load()
	lang := cfg.Lang
	if lang == "" {
		lang = "(auto-detect)"
	}
	model := cfg.Model
	if model == "" {
		model = "(auto-fallback)"
	}
	color.Cyan("━━━ pilot config ━━━\n")
	color.White("  lang   : %s\n", lang)
	color.White("  model  : %s\n", model)
	color.White("  config : ~/.pilot/config.json\n")
}