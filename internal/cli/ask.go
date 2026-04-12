package cli

import (
	"strings"

	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
)

func Ask(args []string) {
	if len(args) == 0 {
		ui.Error("Kullanım: pilot ask <ne yapmak istiyorsun?>")
		return
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")
	ui.Loading("Düşünüyor...")

	result, err := ai.Ask(key, cheatsheet.SystemPrompt, "Şunu yapacak terminal komutu üret: "+query)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	ui.Panel("pilot ask", result, "cyan")
}