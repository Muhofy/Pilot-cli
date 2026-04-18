# 🧭 pilot

> Your terminal co-pilot. Ask in natural language, get the right command.

[![CI/CD](https://github.com/muhofy/pilot/actions/workflows/release.yml/badge.svg)](https://github.com/muhofy/pilot/actions)
[![Go Version](https://img.shields.io/badge/go-1.22-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-GPLv3-blue)](LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/muhofy/pilot)](https://github.com/muhofy/pilot/releases/latest)
[![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-lightgrey)](#installation)

```
$ pilot ask find all log files larger than 50MB

┌──────────────────────────────────────────┐
│ pilot ask                                │
├──────────────────────────────────────────┤
  ```
  find . -name "*.log" -size +50M
  ```
  📌 What it does: Recursively searches the current directory
     for .log files larger than 50 megabytes.
└──────────────────────────────────────────┘
```

---

## ✨ Features

- 🤖 **Natural language → terminal command** via free AI (OpenRouter)
- 🔍 **Command explanation** — understand any command instantly
- ▶️  **Generate & execute** — confirm before running
- ⚠️  **Safety checker** — warns on destructive commands (`rm -rf`, `DROP TABLE`, etc.)
- 📜 **History** — search and replay past queries
- ⚙️  **Config** — set preferred AI model and UI language
- 🌍 **i18n** — English and Turkish, auto-detected from system locale
- 📦 **Single binary** — no runtime, no dependencies, cross-platform

---

## 📦 Installation

### One-liner (recommended)

```bash
curl -fsSL https://raw.githubusercontent.com/muhofy/pilot/main/install.sh | bash
```

Works on Linux, macOS, and Termux (Android). Detects your OS and architecture automatically.

### Homebrew (macOS / Linux)

```bash
brew tap muhofy/tap
brew install pilot
```

### Manual

Download the binary for your platform from [Releases](https://github.com/muhofy/pilot/releases/latest):

| Platform | Binary |
|----------|--------|
| Linux x86_64 | `pilot-linux-amd64` |
| Linux ARM64 | `pilot-linux-arm64` |
| macOS x86_64 | `pilot-darwin-amd64` |
| macOS ARM64 (M1/M2/M3) | `pilot-darwin-arm64` |
| Windows x86_64 | `pilot-windows-amd64.exe` |

```bash
chmod +x pilot-linux-amd64
sudo mv pilot-linux-amd64 /usr/local/bin/pilot
```

### Build from source

```bash
git clone https://github.com/muhofy/pilot.git
cd pilot
./build.sh
```

---

## 🚀 Quick Start

**1. Get a free API key**

```bash
pilot setup
```

Go to [openrouter.ai/keys](https://openrouter.ai/keys), create a key, then:

```bash
export OPENROUTER_API_KEY=your_key_here

# Persist across sessions
echo 'export OPENROUTER_API_KEY=your_key_here' >> ~/.zshrc
```

**2. Ask your first question**

```bash
pilot ask compress a folder into tar.gz
```

---

## 📖 Usage

```
pilot ask     <request>         → natural language → generate command
pilot explain <command>         → explain what a command does
pilot run     <request>         → generate + confirm + execute
pilot history                   → show last 20 queries
pilot history search <keyword>  → search history
pilot history clear             → clear all history
pilot config set model <model>  → set preferred AI model
pilot config set lang  <lang>   → set UI language (en, tr)
pilot config show               → display current config
pilot setup                     → API key setup guide
```

### `pilot ask`

Generate a terminal command from a natural language description.

```bash
pilot ask list all running docker containers
pilot ask find files modified in the last 24 hours
pilot ask show disk usage sorted by size
```

### `pilot explain`

Understand what any command does, broken down by each part.

```bash
pilot explain "tar -czf archive.tar.gz ./dist"
pilot explain "git log --oneline --graph --decorate"
pilot explain "find . -name '*.go' | xargs grep -l 'TODO'"
```

### `pilot run`

Generate a command and execute it after confirmation. Dangerous commands require explicit approval.

```bash
pilot run delete all stopped docker containers
# → shows command
# → ⚠️  warning if destructive
# → Run this command? [y/n]:
```

### `pilot history`

```bash
pilot history                    # last 20 queries
pilot history search docker      # search by keyword
pilot history clear              # clear all
```

### `pilot config`

```bash
pilot config set model deepseek/deepseek-chat-v3.1:free
pilot config set lang tr
pilot config show
```

---

## 🤖 AI Models

Pilot uses [OpenRouter](https://openrouter.ai) with a free-tier model fallback chain:

| Priority | Model |
|----------|-------|
| 1 | `deepseek/deepseek-chat-v3.1:free` |
| 2 | `meta-llama/llama-4-maverick:free` |
| 3 | `qwen/qwen3-235b-a22b:free` |
| 4 | `google/gemma-3-27b-it:free` |
| 5 | `openrouter/free` (last resort) |

Override with `pilot config set model <model-id>`.

---

## ⚙️ Configuration

Config is stored at `~/.pilot/config.json`:

```json
{
  "lang": "en",
  "model": "deepseek/deepseek-chat-v3.1:free"
}
```

### Custom cheatsheet

Drop a CSV at `~/.pilot/cheatsheet.csv` to override the built-in command reference:

```csv
category,command,description
k8s,kubectl get pods,list all pods in current namespace
k8s,kubectl logs -f <pod>,follow pod logs
```

---

## 🛡️ Safety

Pilot automatically detects and warns about destructive commands:

| Level | Examples | Behavior |
|-------|----------|----------|
| ⚠️ Warning | `rm`, `sudo`, `chmod`, `git reset --hard` | Yellow warning panel |
| 🔴 Danger | `rm -rf`, `dd if=`, `DROP TABLE`, fork bomb | Red panel + explicit confirm |

---

## 🌍 Localisation

```bash
pilot config set lang en   # English (default)
pilot config set lang tr   # Turkish
```

System language is auto-detected from `$LANG` / `$LANGUAGE` / `$LC_ALL`.

---

## 🏗️ Project Structure

```
pilot/
├── cmd/pilot/main.go          → entry point
├── internal/
│   ├── ai/openrouter.go       → OpenRouter client, model fallback
│   ├── cli/                   → ask, explain, run, setup, history, config
│   ├── config/config.go       → ~/.pilot/config.json
│   ├── history/history.go     → ~/.pilot/history.db (bbolt)
│   ├── locale/locale.go       → i18n, auto-detect, T() helper
│   ├── safety/checker.go      → dangerous command detection
│   └── ui/panel.go            → dynamic terminal panels
└── pkg/cheatsheet/
    ├── cheatsheet.go          → embed + parse CSV
    └── cheatsheet.csv         → built-in command reference
```

---

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

```bash
git clone https://github.com/muhofy/pilot.git
cd pilot
go mod download
go build ./...
go test ./...
```

---

## 📄 License

[MIT](LICENSE) © [muhofy](https://github.com/muhofy)

---

<div align="center">
  <sub>Built with ❤️ in Go · Powered by OpenRouter free tier</sub>
</div>