package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/history"
	"github.com/muhofy/pilot/internal/ui"
)

func History(args []string) {
	if len(args) == 0 {
		listHistory()
		return
	}

	switch args[0] {
	case "search":
		if len(args) < 2 {
			ui.Error("Kullanım: pilot history search <kelime>")
			return
		}
		searchHistory(strings.Join(args[1:], " "))
	case "clear":
		clearHistory()
	default:
		ui.Error("Kullanım: pilot history [search <kelime>] [clear]")
	}
}

func listHistory() {
	entries, err := history.List(20)
	if err != nil {
		ui.Error("Geçmiş okunamadı: " + err.Error())
		return
	}
	if len(entries) == 0 {
		color.White("Henüz geçmiş yok.")
		return
	}

	color.Cyan("━━━ Son %d Sorgu ━━━\n", len(entries))
	for i, e := range entries {
		icon := iconFor(e.Type)
		timeStr := e.Time.Format("02 Jan 15:04")
		color.White("%2d. %s [%s] %s\n", i+1, icon, timeStr, e.Query)
	}
	fmt.Println()
	color.Yellow("pilot history search <kelime>  → ara")
	color.Yellow("pilot history clear            → temizle")
}

func searchHistory(keyword string) {
	entries, err := history.Search(keyword)
	if err != nil {
		ui.Error("Arama hatası: " + err.Error())
		return
	}
	if len(entries) == 0 {
		color.White("'%s' için sonuç bulunamadı.", keyword)
		return
	}

	color.Cyan("━━━ '%s' için %d sonuç ━━━\n", keyword, len(entries))
	for i, e := range entries {
		icon := iconFor(e.Type)
		timeStr := e.Time.Format("02 Jan 15:04")
		color.White("%2d. %s [%s] %s\n", i+1, icon, timeStr, e.Query)
	}
}

func clearHistory() {
	if !confirm("Tüm geçmiş silinecek, emin misin?") {
		color.White("İptal edildi.")
		return
	}
	if err := history.Clear(); err != nil {
		ui.Error("Temizleme hatası: " + err.Error())
		return
	}
	ui.Success("Geçmiş temizlendi.")
}

func iconFor(typ string) string {
	switch typ {
	case "ask":
		return "🤔"
	case "explain":
		return "📖"
	case "run":
		return "▶️ "
	default:
		return "•"
	}
}