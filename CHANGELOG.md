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

## [0.9.1] — 2026-04-19

### Added
- `pilot chain <steps>` — chain multiple steps into one command via AI
- `pilot chain --dry <steps>` — preview generated chain without executing
- `⛓️` icon for chain entries in history
- Custom system prompt for chain commands — enforces `&&` / `|` structure

---

## [0.9.0] — 2026-04-19

### Added
- `pilot update` — check for latest release and self-update via `install.sh`
- `internal/update/checker.go` — background version check once per day
- Update hint shown after `ask`, `explain`, `run` when new version available:
  `💡 New version available: vX.X.X → run: pilot update`
- Semver comparison with `isNewer()` — handles major/minor/patch correctly
- Update check timestamp stored in `~/.pilot/.update_check`

---

## [0.8.0] — 2026-04-19

### Added
- Homebrew formula (`brew tap muhofy/tap && brew install pilot`)
- GitHub Actions: auto-update Homebrew formula on release tag
- `homebrew-tap` repo: `Formula/pilot.rb`

---

## [0.7.0] — 2026-04-19

### Added
- `--version` / `-v` / `version` flag — version injected via `-ldflags`
- `internal/ui/spinner.go` — animated braille spinner component
- `ask.go`, `explain.go`, `run.go` — `Loading()` replaced with `Spinner`

### Fixed
- `pilot run` parser: `extractCommand()` rewritten with regex
  - Strategy 1: fenced code block (` ``` `)
  - Strategy 2: inline backtick
  - Strategy 3: first line matching command pattern
  - Previously picked up text outside code blocks

---

## [0.6.0] — 2026-04-19

### Added
- `internal/ui/prompt.go` — arrow-key interactive select component (pure bash, no new deps)
- `pilot config set model` — interactive model picker with arrow keys
- `pilot config set lang` — interactive language picker with arrow keys
- Unknown model warning + confirmation prompt in config
- Shell completion for bash, zsh, fish (`pilot completion <shell>`)
- `pilot history clear` — now uses arrow-key confirmation

### Changed
- `pilot run` — `confirm()` replaced with `ui.Select()` (Yes / No / Exit)
- `install.sh` — full rewrite, clean binary download only, no wizard
- `setup.go` — usage updated with completion command

---

## [0.5.0] — 2026-04-11

### Added
- Localisation system (`internal/locale`)
- `lang/en_US.json` — English UI strings
- `lang/tr_TR.json` — Turkish UI strings
- `locale.T(key, args...)` — typed string lookup with fallback to key
- Auto-detection of system language via `$LANG`, `$LANGUAGE`, `$LC_ALL`
- All locale JSON files embedded into binary via `//go:embed`
- Lang validation in `pilot config set lang` — only supported codes accepted

> Commits: `a30eba1` `f3f6694`

---

## [0.4.0] — 2026-04-10

### Added
- `~/.pilot/config.json` — persistent user settings (lang, model)
- `pilot config set model <model>` — override preferred AI model
- `pilot config set lang <lang>` — set UI language
- `pilot config show` — display current config
- Preferred model prepended to fallback list at startup

> Commits: `a30eba1`

---

## [0.3.0] — 2026-04-10

### Added
- `pilot history` — list last 20 queries
- `pilot history search <keyword>` — case-insensitive keyword search
- `pilot history clear` — delete all history with confirmation prompt
- History stored in `~/.pilot/history.db` using `go.etcd.io/bbolt` (pure Go, no CGO)
- History auto-saved for `ask`, `explain`, and `run` commands

> Commits: `8f3d9d9`

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

> Commits: `553f1eb` `72054f8` `72b9d19` `efc5da6` `0cc1a77` `e42933e` `00d9ecb` `1c4a774`

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

> Commits: `553f1eb`

---

[Unreleased]: https://github.com/muhofy/Pilot-cli/compare/v0.9.1...HEAD
[0.9.1]: https://github.com/muhofy/Pilot-cli/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/muhofy/Pilot-cli/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/muhofy/Pilot-cli/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/muhofy/Pilot-cli/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/muhofy/Pilot-cli/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/muhofy/Pilot-cli/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/muhofy/Pilot-cli/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/muhofy/Pilot-cli/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/muhofy/Pilot-cli/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/muhofy/Pilot-cli/releases/tag/v0.1.0