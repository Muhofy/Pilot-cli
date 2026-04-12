# Pilot — TODO & Roadmap

> Terminal-Befehlsassistent. Stelle Fragen in natürlicher Sprache, generiere oder erkläre Befehle.

---

## Statuslegende

| Symbol | Bedeutung |
|--------|-----------|
| ✅ | Erledigt |
| 🔄 | In Bearbeitung |
| 🔲 | Geplant |
| 💡 | Idee / Forschung |
| ❌ | Abgebrochen |

---

## ✅ v0.1.0 — MVP (Erledigt)

- [x] CLI-Architektur mit Go erstellt
- [x] OpenRouter Free-Tier Integration (`openrouter/free`)
- [x] `pilot ask` — natürliche Sprache → Befehls­generierung
- [x] `pilot explain` — Befehl erklären
- [x] `pilot run` — generieren + bestätigen + ausführen
- [x] `pilot setup` — API-Key Setup-Anleitung
- [x] Integriertes Cheatsheet (Terminal, Git, Docker)
- [x] Plattformübergreifende Unterstützung
- [x] Farbige Terminalausgabe (`fatih/color`)
- [x] Initial Commit & Repository erstellt

---

## 🔄 v0.2.0 — Core UX (Aktiv)

### Sicherheit
- [ ] Erkennung gefährlicher Befehle (`rm -rf`, `reset --hard`, `DROP TABLE`, etc.)
- [ ] Rotes ⚠️ Warnpanel
- [ ] `--force` zum Überspringen

### CLI-Erlebnis
- [ ] Einzel-Binary (`pilot`)
- [ ] Globales Installationsskript
- [ ] Interaktiver REPL-Modus
- [ ] Ladeanimation (Spinner)
- [ ] `--dry-run`

### Cheatsheet
- [ ] npm / yarn hinzufügen
- [ ] Kubernetes / kubectl hinzufügen
- [ ] SSH / SCP erweitern

---

## 🔲 v0.3.0 — Verlauf & Speicher

- [ ] Verlauf anzeigen
- [ ] Verlauf durchsuchen
- [ ] Verlauf löschen
- [ ] Speicherung in SQLite
- [ ] Häufige Befehle vorschlagen
- [ ] Favoriten-System

---

## 🔲 v0.4.0 — Konfiguration

- [ ] `config.toml`
- [ ] Modellwahl
- [ ] Sprachwahl
- [ ] Benutzerdefinierte Cheatsheets
- [ ] Config-Kommandos
- [ ] Sichere API-Key Speicherung

---

## 🔲 v0.5.0 — Lokalisierung

- [ ] Mehrsprachige UI
- [ ] Automatische Erkennung
- [ ] Manuelle Einstellung

---

## 🔲 v1.0.0 — Produktionsreif

### Distribution
- [ ] Homebrew
- [ ] apt
- [ ] winget
- [ ] Releases
- [ ] Install Script

### CI/CD
- [ ] GitHub Actions
- [ ] Cross-Builds
- [ ] Release Automation
- [ ] Coverage

### Dokumentation
- [ ] README
- [ ] CONTRIBUTING
- [ ] CHANGELOG
- [ ] Demo

---

## 💡 v1.x — Zukunft

- [ ] Multi-Step Commands
- [ ] Plugin-System
- [ ] Verbose Explain
- [ ] Shell Completion
- [ ] Auto Update
- [ ] Offline Mode
- [ ] VS Code Extension
- [ ] Web UI
- [ ] Telemetrie

---

## 🐛 Bekannte Probleme

- [ ] Parsing Bug
- [ ] Windows Farben
- [ ] UI Probleme bei langen Antworten

---

## 📦 Abhängigkeiten

| Paket | Version | Zweck |
|-------|--------|-------|
| `github.com/fatih/color` | latest | Farbige Ausgabe |