# Contributing to pilot

Thanks for considering a contribution! Here's everything you need to get started.

---

## Development Setup

```bash
git clone https://github.com/muhofy/pilot.git
cd pilot
go mod download
```

Build and install locally:

```bash
./build.sh          # build + install to $PREFIX/bin or /usr/local/bin
./build.sh all      # cross-compile all platforms → dist/
./build.sh clean    # remove build artifacts
```

Run from source without installing:

```bash
go run ./cmd/pilot ask list running containers
```

---

## Project Structure

```
pilot/
├── cmd/pilot/main.go          → entry point, arg routing
├── internal/
│   ├── ai/openrouter.go       → OpenRouter client, model fallback
│   ├── cli/                   → one file per subcommand
│   ├── config/config.go       → ~/.pilot/config.json
│   ├── history/history.go     → bbolt history DB
│   ├── locale/locale.go       → i18n system
│   ├── safety/checker.go      → dangerous command detection
│   └── ui/panel.go            → terminal panel renderer
└── pkg/cheatsheet/
    ├── cheatsheet.go          → CSV loader + prompt builder
    └── cheatsheet.csv         → built-in command reference
```

---

## Adding a New Command

1. Create `internal/cli/<command>.go`
2. Export a function: `func MyCommand(args []string)`
3. Add a `case` in `cmd/pilot/main.go`
4. Add usage line to `cli.Usage()` in `setup.go`
5. Add locale keys to `lang/en_US.json` and `lang/tr_TR.json`

---

## Adding a New Language

1. Create `internal/locale/lang/<locale>.json` (e.g. `de_DE.json`)
2. Copy all keys from `en_US.json` and translate values
3. Add a `case` in `locale.go → Init()`
4. Add to supported langs in `internal/cli/config.go → supportedLangs`
5. Update `pilot config set lang` docs

---

## Cheatsheet Contributions

The built-in cheatsheet lives at `pkg/cheatsheet/cheatsheet.csv`.  
Format: `category,command,description`

```csv
k8s,kubectl get pods -A,list pods across all namespaces
k8s,kubectl logs -f <pod>,follow pod logs in real time
```

Keep descriptions short (under 60 chars) and commands copy-paste ready.

---

## Commit Style

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add shell completion for zsh
fix: parser picks up text outside code block
docs: update README install instructions
chore: bump bbolt to v1.4.3
refactor: extract prompt builder to cheatsheet package
```

---

## Pull Request Checklist

- [ ] `go vet ./...` passes with no errors
- [ ] `go test ./...` passes
- [ ] New features have at least one test
- [ ] Locale keys added for all supported languages
- [ ] CHANGELOG.md updated under `[Unreleased]`

---

## Reporting Bugs

Open an issue with:
- pilot version (`pilot --version`)
- OS and architecture
- Exact command that failed
- Expected vs actual output

---

## License

By contributing, you agree your changes will be licensed under the [GNU General Public License v3.0](LICENSE).