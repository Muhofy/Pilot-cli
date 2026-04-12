# Pilot — TODO & Roadmap

> Terminal komut asistanı. Doğal dilde sor, komutu üret veya açıkla.

---

## Status Legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Done |
| 🔄 | In Progress |
| 🔲 | Planned |
| 💡 | Idea / Research |
| ❌ | Cancelled |

---

## ✅ v0.1.0 — MVP (Done)

- [x] Go ile CLI mimarisi kuruldu
- [x] OpenRouter free tier entegrasyonu (`openrouter/free`)
- [x] `pilot ask` — doğal dil → komut üret
- [x] `pilot explain` — komutu açıkla
- [x] `pilot run` — üret + onay al + çalıştır
- [x] `pilot setup` — API key kurulum rehberi
- [x] Built-in cheatsheet (terminal, git, docker)
- [x] Cross-platform destek (macOS, Linux, Windows)
- [x] Renkli terminal output (`fatih/color`)
- [x] İlk commit & repo oluşturuldu

---

## 🔄 v0.2.0 — Core UX (Active)

### Safety
- [ ] Tehlikeli komut tespiti (`rm -rf`, `reset --hard`, `DROP TABLE` vb.)
- [ ] Tehlikeli komutlar için kırmızı ⚠️ uyarı paneli
- [ ] `--force` flag ile uyarıyı bypass etme seçeneği

### CLI Experience
- [ ] `go build` → tek binary (`pilot`)
- [ ] Global kurulum scripti (`install.sh`)
- [ ] `pilot` yazınca interactive REPL modu açılsın
- [ ] Spinner / loading animasyonu (AI yanıt beklerken)
- [ ] `--dry-run` flag → komutu çalıştırmadan göster

### Cheatsheet
- [ ] npm / yarn komutları eklendi
- [ ] Kubernetes / kubectl komutları eklendi
- [ ] SSH / SCP genişletildi

---

## 🔲 v0.3.0 — History & Memory

- [ ] `pilot history` — geçmiş sorguları listele
- [ ] `pilot history search <query>` — geçmişte ara
- [ ] `pilot history clear` — geçmişi temizle
- [ ] Geçmiş SQLite'da saklanır (`~/.pilot/history.db`)
- [ ] Sık kullanılan komutlar otomatik önerilsin
- [ ] `pilot fav <komut>` — favorilere ekle
- [ ] `pilot fav list` — favorileri listele

---

## 🔲 v0.4.0 — Config & Customization

- [ ] `~/.pilot/config.toml` — kullanıcı ayarları
- [ ] Model seçimi (`openrouter/free`, `gpt-4o`, `claude` vb.)
- [ ] Dil seçimi (`tr`, `en`, `de` vb.) — AI cevap dili
- [ ] Custom cheatsheet desteği (`~/.pilot/cheatsheet.md`)
- [ ] `pilot config set model gpt-4o` komutu
- [ ] `pilot config set lang en` komutu
- [ ] API key güvenli saklama (`keychain` / `secret-service`)

---

## 🔲 v0.5.0 — Lokalizasyon

- [ ] `lang/tr_TR.json` — Türkçe UI metinleri
- [ ] `lang/en_US.json` — İngilizce UI metinleri
- [ ] `lang/de_DE.json` — Almanca UI metinleri
- [ ] Sistem diline göre otomatik algılama
- [ ] `pilot config set lang en` ile manuel değiştirme

---

## 🔲 v1.0.0 — Production Ready

### Distribution
- [ ] `brew install pilot` — Homebrew formula
- [ ] `apt install pilot` — Debian/Ubuntu package
- [ ] `winget install pilot` — Windows Package Manager
- [ ] GitHub Releases — otomatik binary upload (CI/CD)
- [ ] `install.sh` — curl one-liner kurulum

### CI/CD
- [ ] GitHub Actions — lint, test, build pipeline
- [ ] Cross-platform otomatik build (linux/mac/win)
- [ ] Release tagging otomasyonu (`v*` tag → release)
- [ ] Kod coverage raporu

### Documentation
- [ ] `README.md` — tam kurulum ve kullanım rehberi
- [ ] `CONTRIBUTING.md` — katkı rehberi
- [ ] `CHANGELOG.md` — sürüm geçmişi
- [ ] Demo GIF / video (README için)

---

## 💡 v1.x — Future Ideas

- [ ] `pilot chain "X yap, sonra Y yap"` — multi-step komut zinciri
- [ ] Plugin sistemi — kullanıcı kendi cheatsheet'ini yüklesin
- [ ] `pilot explain --verbose` — adım adım detaylı açıklama
- [ ] Shell completion (`zsh`, `bash`, `fish`)
- [ ] `pilot update` — kendi kendini güncelle
- [ ] Offline mod — local Ollama entegrasyonu
- [ ] VS Code extension
- [ ] Web UI (pilot dashboard)
- [ ] Telemetry (opt-in) — en çok sorulan komutlar

---

## 🐛 Known Issues

- [ ] `pilot run` — komut parse'ı bazen ``` bloğu dışındaki metni alıyor
- [ ] Windows'ta renk desteği test edilmedi
- [ ] Çok uzun AI yanıtlarında panel çerçevesi bozuluyor

---

## 📦 Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `github.com/fatih/color` | latest | Renkli terminal output |

---

## 🗓 Milestone Summary

| Version | Focus | Status |
|---------|-------|--------|
| v0.1.0 | MVP | ✅ Done |
| v0.2.0 | Core UX + Safety | 🔄 Active |
| v0.3.0 | History & Memory | 🔲 Planned |
| v0.4.0 | Config & Custom | 🔲 Planned |
| v0.5.0 | Lokalizasyon | 🔲 Planned |
| v1.0.0 | Production Ready | 🔲 Planned |