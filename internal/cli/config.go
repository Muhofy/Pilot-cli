package cli

import (
	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/config"
	"github.com/muhofy/pilot/internal/ui"
)

var supportedLangs = map[string]bool{
	"en": true,
	"tr": true,
	"de": true,
	"es": true,
	"fr": true,
	"zh": true,
}

// knownModels is the list of known free OpenRouter models.
var knownModels = []string{
	"deepseek/deepseek-chat-v3.1:free",
	"meta-llama/llama-4-maverick:free",
	"qwen/qwen3-235b-a22b:free",
	"google/gemma-3-27b-it:free",
	"openrouter/free",
}

// Config handles the config subcommands: set, show.
func Config(args []string) {
	if len(args) == 0 {
		showConfig()
		return
	}

	switch args[0] {
	case "set":
		if len(args) < 2 {
			ui.Error("Usage: pilot config set <key> [value]")
			color.Yellow("  Keys: lang, model")
			return
		}
		handleSet(args[1], args[2:])

	case "show":
		showConfig()

	default:
		ui.Error("Usage: pilot config [set <key> [value]] [show]")
	}
}

// handleSet routes to the appropriate setter based on key.
func handleSet(key string, rest []string) {
	switch key {
	case "lang":
		setLang(rest)
	case "model":
		setModel(rest)
	default:
		ui.Error("Unknown config key: " + key)
		color.Yellow("  Supported keys: lang, model")
	}
}

// setLang sets UI language via arrow-key select or direct value.
func setLang(args []string) {
	var value string

	if len(args) == 0 {
		// Interactive select
		res := ui.Select("Select language", []ui.Option{
			{Label: "English", Value: "en"},
			{Label: "Turkish / Türkçe", Value: "tr"},
		})
		if res.Index == -1 {
			color.White("Cancelled.")
			return
		}
		value = res.Value
	} else {
		value = args[0]
	}

	if !supportedLangs[value] {
		ui.Error("Unsupported language: " + value)
		color.Yellow("  Supported: en, tr")
		return
	}

	if err := config.Set("lang", value); err != nil {
		ui.Error("Could not save config: " + err.Error())
		return
	}
	ui.Success("Language set to: " + value)
}

// setModel sets AI model via arrow-key select or direct value.
func setModel(args []string) {
	var value string

	if len(args) == 0 {
		// Interactive select
		options := make([]ui.Option, len(knownModels))
		for i, m := range knownModels {
			options[i] = ui.Option{Label: m, Value: m}
		}

		res := ui.Select("Select AI model", options)
		if res.Index == -1 {
			color.White("Cancelled.")
			return
		}
		value = res.Value
	} else {
		value = args[0]
		// Warn if model is not in known list
		if !isKnownModel(value) {
			color.Yellow("⚠️  Unknown model: %s", value)
			color.Yellow("   This model may not work. Known models:")
			for _, m := range knownModels {
				color.Yellow("   - %s", m)
			}
			// Ask confirmation via arrow-key
			res := ui.Select("Use it anyway?", []ui.Option{
				{Label: "Yes, use it", Value: "yes"},
				{Label: "No, cancel",  Value: "no"},
			})
			if res.Value != "yes" {
				color.White("Cancelled.")
				return
			}
		}
	}

	if err := config.Set("model", value); err != nil {
		ui.Error("Could not save config: " + err.Error())
		return
	}
	ui.Success("Model set to: " + value)
}

// isKnownModel checks if a model string is in the known list.
func isKnownModel(model string) bool {
	for _, m := range knownModels {
		if m == model {
			return true
		}
	}
	return false
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
	color.Yellow("\n  pilot config set model  → select AI model")
	color.Yellow("  pilot config set lang   → select language")
}