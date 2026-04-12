package cli

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/muhofy/pilot/internal/ai"
	"github.com/muhofy/pilot/internal/safety"
	"github.com/muhofy/pilot/internal/ui"
	"github.com/muhofy/pilot/pkg/cheatsheet"
	"github.com/fatih/color"
)

func Run(args []string) {
	if len(args) == 0 {
		ui.Error("Kullanım: pilot run <ne yapmak istiyorsun?>")
		return
	}

	key, err := ai.GetAPIKey()
	if err != nil {
		ui.Error(err.Error())
		return
	}

	query := strings.Join(args, " ")
	ui.Loading("Komut üretiliyor...")

	result, err := ai.Ask(key, cheatsheet.SystemPrompt,
		"Şunu yapacak terminal komutunu SADECE ver (açıklama yok): "+query)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	cmd := extractCommand(result)
	ui.Panel("Üretilen Komut", cmd, "green")

	// Safety check
	check := safety.Check(cmd)
	switch check.Level {
	case safety.Danger:
		ui.Panel("⚠️  TEHLİKELİ KOMUT", check.Reason, "red")
		color.Red("Bu komut geri alınamaz! Devam etmek istediğinden emin misin?")
	case safety.Warning:
		ui.Warning(check.Reason)
	}

	if confirm("Bu komutu çalıştırayım mı?") {
		execCommand(cmd)
	} else {
		color.White("İptal edildi.")
	}
}

func extractCommand(text string) string {
	lines := strings.Split(text, "\n")
	var cmd []string
	inBlock := false
	for _, l := range lines {
		if strings.HasPrefix(l, "```") {
			inBlock = !inBlock
			continue
		}
		if inBlock {
			cmd = append(cmd, l)
		}
	}
	if len(cmd) > 0 {
		return strings.TrimSpace(strings.Join(cmd, "\n"))
	}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" && !strings.HasPrefix(l, "📌") {
			return l
		}
	}
	return strings.TrimSpace(text)
}

func confirm(prompt string) bool {
	color.Yellow(prompt + " [e/h]: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ans := strings.ToLower(strings.TrimSpace(sc.Text()))
	return ans == "e" || ans == "evet" || ans == "y" || ans == "yes"
}

func execCommand(cmd string) {
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", cmd)
	} else {
		c = exec.Command("sh", "-c", cmd)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}