package cli

import "github.com/fatih/color"

func Usage() {
	color.Cyan(`
  pilot — Terminal komut asistanı

  Kullanım:
    pilot ask     <istek>          → doğal dil → komut üret
    pilot explain <komut>          → komutu açıkla
    pilot run     <istek>          → üret + onay al + çalıştır
    pilot history                  → geçmiş sorgular
    pilot history search <kelime>  → geçmişte ara
    pilot history clear            → geçmişi temizle
    pilot setup                    → API key kurulum rehberi
`)
}

func Setup() {
	color.Cyan(`
  pilot — Kurulum Rehberi
  ───────────────────────
  1. https://openrouter.ai/keys adresine git
  2. Sign up / Login
  3. "Create Key" butonuna tıkla
  4. Terminale yapıştır:

     export OPENROUTER_API_KEY=your_key_here

  5. Kalıcı yapmak için ~/.zshrc veya ~/.bashrc dosyana ekle:

     echo 'export OPENROUTER_API_KEY=your_key' >> ~/.zshrc
     source ~/.zshrc
`)
}