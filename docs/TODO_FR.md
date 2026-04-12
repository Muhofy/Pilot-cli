# Pilot — TODO & Feuille de Route

> Assistant de commandes terminal. Posez des questions en langage naturel, générez ou expliquez des commandes.

---

## Légende des Statuts

| Symbole | Signification |
|--------|---------------|
| ✅ | Terminé |
| 🔄 | En cours |
| 🔲 | Planifié |
| 💡 | Idée / Recherche |
| ❌ | Annulé |

---

## ✅ v0.1.0 — MVP (Terminé)

- [x] Architecture CLI construite avec Go
- [x] Intégration OpenRouter free tier (`openrouter/free`)
- [x] `pilot ask` — langage naturel → génération de commandes
- [x] `pilot explain` — expliquer une commande
- [x] `pilot run` — générer + confirmer + exécuter
- [x] `pilot setup` — guide de configuration API key
- [x] Cheatsheet intégré (terminal, git, docker)
- [x] Support multiplateforme (macOS, Linux, Windows)
- [x] Sortie terminal colorée (`fatih/color`)
- [x] Commit initial & dépôt créé

---

## 🔄 v0.2.0 — UX Principale (Actif)

### Sécurité
- [ ] Détection des commandes dangereuses (`rm -rf`, `reset --hard`, `DROP TABLE`, etc.)
- [ ] Panneau d’avertissement rouge ⚠️
- [ ] Option de contournement avec `--force`

### Expérience CLI
- [ ] Binaire unique (`pilot`)
- [ ] Script d’installation global (`install.sh`)
- [ ] Mode REPL interactif
- [ ] Animation de chargement (spinner)
- [ ] `--dry-run` — afficher sans exécuter

### Cheatsheet
- [ ] Ajouter npm / yarn
- [ ] Ajouter Kubernetes / kubectl
- [ ] Étendre SSH / SCP

---

## 🔲 v0.3.0 — Historique & Mémoire

- [ ] Historique des commandes
- [ ] Recherche dans l’historique
- [ ] Nettoyage de l’historique
- [ ] Stockage SQLite
- [ ] Suggestions automatiques
- [ ] Favoris

---

## 🔲 v0.4.0 — Configuration

- [ ] `config.toml`
- [ ] Sélection du modèle
- [ ] Sélection de la langue
- [ ] Cheatsheet personnalisé
- [ ] Commandes de config
- [ ] Stockage sécurisé API key

---

## 🔲 v0.5.0 — Localisation

- [ ] UI multilingue
- [ ] Détection automatique
- [ ] Override manuel

---

## 🔲 v1.0.0 — Prêt pour Production

### Distribution
- [ ] Homebrew
- [ ] apt
- [ ] winget
- [ ] Releases GitHub
- [ ] Installation via script

### CI/CD
- [ ] GitHub Actions
- [ ] Builds multiplateformes
- [ ] Automatisation des releases
- [ ] Coverage

### Documentation
- [ ] README
- [ ] CONTRIBUTING
- [ ] CHANGELOG
- [ ] Démo

---

## 💡 v1.x — Idées Futures

- [ ] Chaînage de commandes
- [ ] Système de plugins
- [ ] Mode verbose
- [ ] Autocomplétion shell
- [ ] Mise à jour automatique
- [ ] Mode offline
- [ ] Extension VS Code
- [ ] Interface Web
- [ ] Télémétrie (optionnelle)

---

## 🐛 Problèmes Connus

- [ ] Bug de parsing
- [ ] Couleurs Windows non testées
- [ ] UI casse avec réponses longues

---

## 📦 Dépendances

| Package | Version | Usage |
|---------|--------|-------|
| `github.com/fatih/color` | latest | Couleurs terminal |