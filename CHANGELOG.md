# Changelog

All notable changes to **pilot** are documented here.  
Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).  
Versioning follows [Semantic Versioning](https://semver.org/).

---

## [Unreleased] — v1.0.0

### Planned
- `README.md` — full install + usage guide with demo GIF
- `CONTRIBUTING.md` — contribution guide
- `install.sh` — one-liner curl install
- GitHub Actions CI/CD — lint, build, test, release pipeline
- Cross-platform binary upload on `v*` tag push
- Homebrew formula (`brew install muhofy/tap/pilot`)

---

## [0.5.0] — 2026-04-10

### Added
- Localisation system (`internal/locale`)
- `lang/en_US.json` — English UI strings
- `lang/tr_TR.json` — Turkish UI strings
- `locale.T(key, args...)` — typed string lookup with fallback to key
- Auto-detection of system language via `$LANG`, `$LANGUAGE`, `$LC_ALL`
- All locale JSON files embedded into binary via `//go:embed`

---

## [0.4.0] — 2026-04-11

### Added
- `~/.pilot/config.json` — persistent user settings (lang, model)
- `pilot config set model <model>` — override preferred AI model
- `pilot config set lang <lang>` — set UI language
- `pilot config show` — display current config
- Lang validation — only supported language codes accepted
- Preferred model prepended to fallback list at startup

---

## [0.3.0] — 2026-04-10

### Added
- `pilot history` — list last 20 queries
- `pilot history search <keyword>` — case-insensitive keyword search
- `pilot history clear` — delete all history with confirmation prompt
- History stored in `~/.pilot/history.db` using `go.etcd.io/bbolt` (pure Go, no CGO)
- History auto-saved for `ask`, `explain`, and `run` commands

---

## [0.2.0] — 2026-04-10

### Added
- Dangerous command detection (`rm -rf`, `dd if=`, `DROP TABLE`, fork bomb, etc.)
- Warning-level detection (`rm`, `sudo`, `chmod`, `git reset --hard`, etc.)
- Red ⚠️ danger panel and yellow warning panel in `pilot run`
- Dynamic panel width based on terminal size (`golang.org/x/term`)
- Model fallback list — tries 5 models in order, last resort `openrouter/free`
- Response validation — retries on gibberish or empty AI response
- Cheatsheet migrated to embedded CSV (`pkg/cheatsheet/cheatsheet.csv`)
- User cheatsheet override via `~/.pilot/cheatsheet.csv`
- `build.sh` — local build + auto-install + cross-platform (`local` / `all` / `clean`)

### Changed
- All UI text and code comments migrated to English
- Tightened system prompt to prevent non-command AI output
- Migrated to standard Go project layout (`cmd/`, `internal/`, `pkg/`)

### Fixed
- AI response occasionally returning prose instead of a command block

---

## [0.1.0] — 2026-04-10

### Added
- Initial release
- `pilot ask <request>` — natural language → generate terminal command
- `pilot explain <command>` — explain what a command does
- `pilot run <request>` — generate + confirm + execute
- `pilot setup` — API key setup guide
- OpenRouter free tier integration with `openrouter/free` model
- Built-in command cheatsheet (terminal, git, docker, npm)
- Colored terminal output via `github.com/fatih/color`
- Cross-platform support: Linux, macOS, Windows

---

[Unreleased]: https://github.com/muhofy/pilot/compare/v0.5.0...HEAD
[0.5.0]: https://github.com/muhofy/pilot/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/muhofy/pilot/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/muhofy/pilot/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/muhofy/pilot/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/muhofy/pilot/releases/tag/v0.1.0