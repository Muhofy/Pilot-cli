# Pilot — TODO y Hoja de Ruta

> Asistente de comandos de terminal. Pregunta en lenguaje natural, genera o explica comandos.

---

## Leyenda de Estado

| Símbolo | Significado |
|--------|-------------|
| ✅ | Completado |
| 🔄 | En progreso |
| 🔲 | Planificado |
| 💡 | Idea / Investigación |
| ❌ | Cancelado |

---

## ✅ v0.1.0 — MVP (Completado)

- [x] Arquitectura CLI construida con Go
- [x] Integración con OpenRouter free tier (`openrouter/free`)
- [x] `pilot ask` — lenguaje natural → generación de comandos
- [x] `pilot explain` — explicar comando
- [x] `pilot run` — generar + confirmar + ejecutar
- [x] `pilot setup` — guía de configuración de API key
- [x] Cheatsheet integrado (terminal, git, docker)
- [x] Soporte multiplataforma (macOS, Linux, Windows)
- [x] Salida de terminal con colores (`fatih/color`)
- [x] Commit inicial y repositorio creado

---

## 🔄 v0.2.0 — UX Central (Activo)

### Seguridad
- [ ] Detección de comandos peligrosos (`rm -rf`, `reset --hard`, `DROP TABLE`, etc.)
- [ ] Panel de advertencia rojo ⚠️ para comandos riesgosos
- [ ] Opción para omitir advertencias con `--force`

### Experiencia CLI
- [ ] `go build` → binario único (`pilot`)
- [ ] Script de instalación global (`install.sh`)
- [ ] Modo REPL interactivo al ejecutar `pilot`
- [ ] Animación de carga (esperando respuesta de la IA)
- [ ] Flag `--dry-run` → mostrar comando sin ejecutar

### Cheatsheet
- [ ] Añadir comandos npm / yarn
- [ ] Añadir comandos Kubernetes / kubectl
- [ ] Ampliar soporte SSH / SCP

---

## 🔲 v0.3.0 — Historial y Memoria

- [ ] `pilot history` — listar consultas pasadas
- [ ] `pilot history search <query>` — buscar en historial
- [ ] `pilot history clear` — limpiar historial
- [ ] Guardar historial en SQLite (`~/.pilot/history.db`)
- [ ] Sugerir automáticamente comandos frecuentes
- [ ] `pilot fav <comando>` — añadir a favoritos
- [ ] `pilot fav list` — listar favoritos

---

## 🔲 v0.4.0 — Configuración y Personalización

- [ ] `~/.pilot/config.toml` — configuración del usuario
- [ ] Selección de modelo (`openrouter/free`, `gpt-4o`, `claude`, etc.)
- [ ] Selección de idioma (`tr`, `en`, `de`, etc.)
- [ ] Soporte de cheatsheet personalizado (`~/.pilot/cheatsheet.md`)
- [ ] `pilot config set model gpt-4o`
- [ ] `pilot config set lang en`
- [ ] Almacenamiento seguro de API key (`keychain` / `secret-service`)

---

## 🔲 v0.5.0 — Localización

- [ ] `lang/tr_TR.json` — textos UI en turco
- [ ] `lang/en_US.json` — textos UI en inglés
- [ ] `lang/de_DE.json` — textos UI en alemán
- [ ] Detección automática del idioma del sistema
- [ ] Cambio manual con `pilot config set lang en`

---

## 🔲 v1.0.0 — Listo para Producción

### Distribución
- [ ] `brew install pilot`
- [ ] `apt install pilot`
- [ ] `winget install pilot`
- [ ] GitHub Releases — subida automática de binarios
- [ ] Instalación rápida con `install.sh`

### CI/CD
- [ ] GitHub Actions — lint, test, build
- [ ] Build automático multiplataforma
- [ ] Automatización de releases (`v*`)
- [ ] Reporte de cobertura

### Documentación
- [ ] `README.md`
- [ ] `CONTRIBUTING.md`
- [ ] `CHANGELOG.md`
- [ ] Demo GIF / video

---

## 💡 v1.x — Ideas Futuras

- [ ] `pilot chain "haz X, luego Y"`
- [ ] Sistema de plugins
- [ ] `pilot explain --verbose`
- [ ] Autocompletado de shell
- [ ] `pilot update`
- [ ] Modo offline (Ollama)
- [ ] Extensión VS Code
- [ ] Interfaz web
- [ ] Telemetría (opcional)

---

## 🐛 Problemas Conocidos

- [ ] `pilot run` — parsing incorrecto fuera de ``` bloques
- [ ] Colores no probados en Windows
- [ ] UI se rompe con respuestas largas

---

## 📦 Dependencias

| Paquete | Versión | Propósito |
|---------|--------|-----------|
| `github.com/fatih/color` | latest | Salida con color |