# Pilot — TODO & Roadmap

> Terminal command assistant. Ask in natural language, generate or explain commands.

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

- [x] CLI architecture built with Go
- [x] OpenRouter free tier integration (`openrouter/free`)
- [x] `pilot ask` — natural language → command generation
- [x] `pilot explain` — explain command
- [x] `pilot run` — generate + confirm + execute
- [x] `pilot setup` — API key setup guide
- [x] Built-in cheatsheet (terminal, git, docker)
- [x] Cross-platform support (macOS, Linux, Windows)
- [x] Colored terminal output (`fatih/color`)
- [x] Initial commit & repository created

---

## 🔄 v0.2.0 — Core UX (Active)

### Safety
- [ ] Dangerous command detection (`rm -rf`, `reset --hard`, `DROP TABLE`, etc.)
- [ ] Red ⚠️ warning panel for risky commands
- [ ] Option to bypass warning with `--force` flag

### CLI Experience
- [ ] `go build` → single binary (`pilot`)
- [ ] Global installation script (`install.sh`)
- [ ] Launch interactive REPL mode when running `pilot`
- [ ] Spinner / loading animation (while waiting for AI response)
- [ ] `--dry-run` flag → show command without executing

### Cheatsheet
- [ ] Add npm / yarn commands
- [ ] Add Kubernetes / kubectl commands
- [ ] Extend SSH / SCP coverage

---

## 🔲 v0.3.0 — History & Memory

- [ ] `pilot history` — list past queries
- [ ] `pilot history search <query>` — search history
- [ ] `pilot history clear` — clear history
- [ ] Store history in SQLite (`~/.pilot/history.db`)
- [ ] Auto-suggest frequently used commands
- [ ] `pilot fav <command>` — add to favorites
- [ ] `pilot fav list` — list favorites

---

## 🔲 v0.4.0 — Config & Customization

- [ ] `~/.pilot/config.toml` — user configuration
- [ ] Model selection (`openrouter/free`, `gpt-4o`, `claude`, etc.)
- [ ] Language selection (`tr`, `en`, `de`, etc.) — AI response language
- [ ] Custom cheatsheet support (`~/.pilot/cheatsheet.md`)
- [ ] `pilot config set model gpt-4o` command
- [ ] `pilot config set lang en` command
- [ ] Secure API key storage (`keychain` / `secret-service`)

---

## 🔲 v0.5.0 — Localization

- [ ] `lang/tr_TR.json` — Turkish UI texts
- [ ] `lang/en_US.json` — English UI texts
- [ ] `lang/de_DE.json` — German UI texts
- [ ] Automatic system language detection
- [ ] Manual override with `pilot config set lang en`

---

## 🔲 v1.0.0 — Production Ready

### Distribution
- [ ] `brew install pilot` — Homebrew formula
- [ ] `apt install pilot` — Debian/Ubuntu package
- [ ] `winget install pilot` — Windows Package Manager
- [ ] GitHub Releases — automatic binary uploads (CI/CD)
- [ ] `install.sh` — curl one-liner installation

### CI/CD
- [ ] GitHub Actions — lint, test, build pipeline
- [ ] Cross-platform automated builds (linux/mac/win)
- [ ] Release tagging automation (`v*` tag → release)
- [ ] Code coverage report

### Documentation
- [ ] `README.md` — full setup and usage guide
- [ ] `CONTRIBUTING.md` — contribution guide
- [ ] `CHANGELOG.md` — version history
- [ ] Demo GIF / video (for README)

---

## 💡 v1.x — Future Ideas

- [ ] `pilot chain "do X, then Y"` — multi-step command chaining
- [ ] Plugin system — allow users to load custom cheatsheets
- [ ] `pilot explain --verbose` — step-by-step detailed explanation
- [ ] Shell completion (`zsh`, `bash`, `fish`)
- [ ] `pilot update` — self-update feature
- [ ] Offline mode — local Ollama integration
- [ ] VS Code extension
- [ ] Web UI (pilot dashboard)
- [ ] Telemetry (opt-in) — most requested commands

---

## 🐛 Known Issues

- [ ] `pilot run` — sometimes parses text outside of ``` blocks
- [ ] Color support not tested on Windows
- [ ] UI panel breaks on very long AI responses

---

## 📦 Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| `github.com/fatih/color` | latest | Colored terminal output |

---

## 🗓 Milestone Summary

| Version | Focus | Status |
|---------|-------|--------|
| v0.1.0 | MVP | ✅ Done |
| v0.2.0 | Core UX + Safety | 🔄 Active |
| v0.3.0 | History & Memory | 🔲 Planned |
| v0.4.0 | Config & Custom | 🔲 Planned |
| v0.5.0 | Localization | 🔲 Planned |
| v1.0.0 | Production Ready | 🔲 Planned |