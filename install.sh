#!/bin/bash
set -e

# ─────────────────────────────────────────────────────────────
#  pilot — Installer
#  curl -fsSL https://raw.githubusercontent.com/muhofy/Pilot-cli/main/install.sh -o install.sh && bash install.sh
# ─────────────────────────────────────────────────────────────

APP="pilot"
REPO="muhofy/Pilot-cli"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"
ENV_FILE="$HOME/.pilot_env"
CONFIG_FILE="$HOME/.pilot/config.json"

# ── ANSI ─────────────────────────────────────────────────────
R='\033[0;31m'
G='\033[0;32m'
Y='\033[1;33m'
C='\033[0;36m'
DIM='\033[2m'
BOLD='\033[1m'
NC='\033[0m'

info()    { echo -e "${C}  ❯ ${NC}$1"; }
success() { echo -e "${G}  ✓ ${NC}$1"; }
warn()    { echo -e "${Y}  ⚠ ${NC}$1"; }
error()   { echo -e "${R}  ✗ ${NC}$1"; exit 1; }
dim()     { echo -e "${DIM}    $1${NC}"; }
nl()      { echo ""; }

# ── Banner ────────────────────────────────────────────────────
banner() {
  clear
  echo -e "${C}"
  echo "  ██████╗ ██╗██╗      ██████╗ ████████╗"
  echo "  ██╔══██╗██║██║     ██╔═══██╗╚══██╔══╝"
  echo "  ██████╔╝██║██║     ██║   ██║   ██║   "
  echo "  ██╔═══╝ ██║██║     ██║   ██║   ██║   "
  echo "  ██║     ██║███████╗╚██████╔╝   ██║   "
  echo "  ╚═╝     ╚═╝╚══════╝ ╚═════╝    ╚═╝   "
  echo -e "${NC}"
  echo -e "  ${BOLD}Your terminal co-pilot${NC}  ${DIM}— Setup Wizard${NC}"
  echo -e "  ${DIM}────────────────────────────────────────${NC}"
  nl
}

# ── Arrow-key select ─────────────────────────────────────────
arrow_select() {
  local question="$1"
  shift
  local options=("$@")
  local n=${#options[@]}
  local cursor=0

  local old_stty
  old_stty=$(stty -g 2>/dev/null) || true

  _restore() {
    stty "$old_stty" 2>/dev/null || true
    tput cnorm 2>/dev/null || true
  }
  trap _restore EXIT INT TERM

  tput civis 2>/dev/null || true
  stty raw -echo 2>/dev/null || true

  _draw() {
    if [ "$1" != "first" ]; then
      for (( i=0; i<n+1; i++ )); do
        tput cuu1 2>/dev/null
        tput el   2>/dev/null
      done
    fi
    echo -e "  ${Y}?${NC} ${BOLD}${question}${NC}"
    for (( i=0; i<n; i++ )); do
      if [ $i -eq $cursor ]; then
        echo -e "  ${C}❯ ${BOLD}${options[$i]}${NC}"
      else
        echo -e "  ${DIM}  ${options[$i]}${NC}"
      fi
    done
  }

  _draw "first"

  while true; do
    local k1 k2 k3
    IFS= read -r -s -n1 k1
    if [[ "$k1" == $'\x1b' ]]; then
      IFS= read -r -s -n1 -t 0.1 k2 || true
      IFS= read -r -s -n1 -t 0.1 k3 || true
    fi

    if [[ "$k1" == $'\x1b' && "$k2" == '[' && "$k3" == 'A' ]] || [[ "$k1" == 'k' ]]; then
      (( cursor > 0 )) && (( cursor-- )) || cursor=$(( n - 1 ))
      _draw
    elif [[ "$k1" == $'\x1b' && "$k2" == '[' && "$k3" == 'B' ]] || [[ "$k1" == 'j' ]]; then
      (( cursor < n-1 )) && (( cursor++ )) || cursor=0
      _draw
    elif [[ "$k1" == '' || "$k1" == $'\n' || "$k1" == $'\r' ]]; then
      break
    elif [[ "$k1" == $'\x03' ]]; then
      _restore
      nl
      error "Aborted."
    fi
  done

  _restore
  trap - EXIT INT TERM
  nl

  ARROW_RESULT="${options[$cursor]}"
  ARROW_INDEX=$cursor
}

# ── Progress bar ──────────────────────────────────────────────
download_with_progress() {
  local url="$1"
  local out="$2"
  local bar_width=40
  local spinner_chars="⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
  local spin_idx=0

  _draw_bar() {
    local pct=$1 speed=$2
    local filled=$(( pct * bar_width / 100 ))
    local empty=$(( bar_width - filled ))
    local bar="" i
    for (( i=0; i<filled; i++ )); do bar+="█"; done
    for (( i=0; i<empty;  i++ )); do bar+="░"; done
    local spin="${spinner_chars:$spin_idx:1}"
    spin_idx=$(( (spin_idx + 1) % ${#spinner_chars} ))
    printf "\r  ${C}%s${NC} [${G}%s${NC}] ${BOLD}%3d%%${NC}  ${DIM}%s${NC}  " \
      "$spin" "$bar" "$pct" "$speed"
  }

  nl
  info "Downloading ${APP} ${VERSION}..."
  nl

  if command -v curl &>/dev/null; then
    curl -fsSL --progress-bar "$url" -o "$out" 2>&1 | while IFS= read -r line; do
      local pct
      pct=$(echo "$line" | grep -oE '[0-9]+' | head -1 || echo 0)
      [ -n "$pct" ] && _draw_bar "$pct" ""
    done
  elif command -v wget &>/dev/null; then
    wget -q --show-progress "$url" -O "$out" 2>&1 | while IFS= read -r line; do
      local pct speed
      pct=$(echo "$line" | grep -oP '\d+(?=%)' | tail -1 || echo 0)
      speed=$(echo "$line" | grep -oP '[\d.]+[KMG]B/s' | tail -1 || echo "")
      [ -n "$pct" ] && _draw_bar "$pct" "$speed"
    done
  else
    error "curl or wget is required"
  fi

  printf "\r  ${G}✓${NC} [$(printf '█%.0s' $(seq 1 $bar_width))] ${BOLD}100%%${NC}            \n"
  nl
}

# ── Detect platform ───────────────────────────────────────────
detect_platform() {
  local os arch
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  arch=$(uname -m)

  case "$os" in
    linux)  OS="linux" ;;
    darwin) OS="darwin" ;;
    *)      error "Unsupported OS: $os" ;;
  esac

  case "$arch" in
    x86_64|amd64)  ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) error "Unsupported architecture: $arch" ;;
  esac

  BINARY="${APP}-${OS}-${ARCH}"
}

# ── Detect install dir ────────────────────────────────────────
detect_install_dir() {
  if [ -n "$PREFIX" ] && [ -d "$PREFIX/bin" ]; then
    INSTALL_DIR="$PREFIX/bin"
  elif [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
  elif [ -d "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
    NEEDS_SUDO=1
  else
    INSTALL_DIR="$HOME/bin"
    mkdir -p "$INSTALL_DIR"
  fi
}

# ── Detect shell ──────────────────────────────────────────────
detect_shell() {
  DETECTED_SHELL=$(basename "$SHELL" 2>/dev/null || echo "bash")
}

# ── Fetch latest version ──────────────────────────────────────
fetch_version() {
  if command -v curl &>/dev/null; then
    VERSION=$(curl -fsSL "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  else
    VERSION=$(wget -qO- "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  fi
  [ -z "$VERSION" ] && error "Could not fetch latest version."
}

# ── Write config ──────────────────────────────────────────────
write_config() {
  mkdir -p "$HOME/.pilot"
  cat > "$CONFIG_FILE" <<EOF
{
  "lang": "${1}",
  "model": "${2}"
}
EOF
}

# ── Write API key ─────────────────────────────────────────────
write_api_key() {
  echo "export OPENROUTER_API_KEY=${1}" > "$ENV_FILE"
  chmod 600 "$ENV_FILE"
}

# ── Add to shell profile ──────────────────────────────────────
add_to_profile() {
  local shell="$1" profile=""
  case "$shell" in
    zsh)  profile="$HOME/.zshrc" ;;
    fish) profile="$HOME/.config/fish/config.fish" ;;
    *)    profile="$HOME/.bashrc" ;;
  esac

  local src='[ -f "$HOME/.pilot_env" ] && source "$HOME/.pilot_env"'
  local cmp=""
  case "$shell" in
    zsh)  cmp='eval "$(pilot completion zsh)"' ;;
    fish) cmp='pilot completion fish | source' ;;
    *)    cmp='eval "$(pilot completion bash)"' ;;
  esac

  mkdir -p "$(dirname "$profile")"
  grep -qF '.pilot_env'       "$profile" 2>/dev/null || echo "$src" >> "$profile"
  grep -qF 'pilot completion' "$profile" 2>/dev/null || echo "$cmp" >> "$profile"

  success "Profile updated: ${DIM}${profile}${NC}"
}

# ═════════════════════════════════════════════════════════════
#  MAIN
# ═════════════════════════════════════════════════════════════
main() {
  banner
  detect_platform
  detect_install_dir
  detect_shell

  # ── Step 1: Language ───────────────────────────────────────
  echo -e "  ${BOLD}Step 1 of 4${NC}  ${DIM}Language${NC}"
  nl
  arrow_select "Select your language" "English" "Turkish / Türkçe"
  case "$ARROW_INDEX" in
    1) LANG_CODE="tr" ;;
    *) LANG_CODE="en" ;;
  esac
  success "Language: ${BOLD}${LANG_CODE}${NC}"
  nl

  # ── Step 2: AI Model ───────────────────────────────────────
  echo -e "  ${BOLD}Step 2 of 4${NC}  ${DIM}AI Model${NC}"
  nl
  arrow_select "Select AI model" \
    "DeepSeek v3.1  (recommended)" \
    "Llama 4 Maverick" \
    "Qwen3 235B" \
    "Gemma 3 27B" \
    "Auto (fallback)"
  case "$ARROW_INDEX" in
    0) MODEL="deepseek/deepseek-chat-v3.1:free" ;;
    1) MODEL="meta-llama/llama-4-maverick:free" ;;
    2) MODEL="qwen/qwen3-235b-a22b:free" ;;
    3) MODEL="google/gemma-3-27b-it:free" ;;
    *) MODEL="" ;;
  esac
  success "Model: ${BOLD}${MODEL:-auto}${NC}"
  nl

  # ── Step 3: Shell ──────────────────────────────────────────
  echo -e "  ${BOLD}Step 3 of 4${NC}  ${DIM}Shell  (detected: ${DETECTED_SHELL})${NC}"
  nl
  arrow_select "Select your shell" "bash" "zsh" "fish"
  CHOSEN_SHELL="$ARROW_RESULT"
  success "Shell: ${BOLD}${CHOSEN_SHELL}${NC}"
  nl

  # ── Step 4: API Key ────────────────────────────────────────
  echo -e "  ${BOLD}Step 4 of 4${NC}  ${DIM}OpenRouter API Key${NC}"
  nl
  dim "Get a free key at: https://openrouter.ai/keys"
  nl

  local api_key=""
  while true; do
    printf "  ${Y}❯${NC} Paste your API key: "
    read -rs api_key
    nl
    if [ -z "$api_key" ]; then
      warn "API key cannot be empty."
    elif [[ ! "$api_key" == sk-or-* ]]; then
      warn "Key doesn't look right (should start with sk-or-)."
      printf "  ${Y}?${NC} Continue anyway? [y/n]: "
      read -r confirm
      [[ "$confirm" == "y" || "$confirm" == "yes" ]] && break
    else
      break
    fi
  done
  success "API key saved"
  nl

  # ── Download + Install ─────────────────────────────────────
  fetch_version
  local url="${GITHUB_DL}/${VERSION}/${BINARY}"
  local tmp_bin
  tmp_bin=$(mktemp)

  download_with_progress "$url" "$tmp_bin"
  chmod +x "$tmp_bin"

  if [ "${NEEDS_SUDO}" = "1" ]; then
    info "Installing to ${INSTALL_DIR} (requires sudo)..."
    sudo mv "$tmp_bin" "${INSTALL_DIR}/${APP}"
  else
    mv "$tmp_bin" "${INSTALL_DIR}/${APP}"
  fi

  write_config "$LANG_CODE" "$MODEL"
  write_api_key "$api_key"
  add_to_profile "$CHOSEN_SHELL"

  # ── Done ───────────────────────────────────────────────────
  nl
  echo -e "  ${DIM}────────────────────────────────────────${NC}"
  echo -e "  ${G}${BOLD}pilot is ready! 🎉${NC}"
  echo -e "  ${DIM}────────────────────────────────────────${NC}"
  nl
  echo -e "  ${C}❯${NC} Reload your shell:"
  if [ "$CHOSEN_SHELL" = "fish" ]; then
    dim "source ~/.config/fish/config.fish"
  else
    dim "source ~/.${CHOSEN_SHELL}rc"
  fi
  nl
  echo -e "  ${C}❯${NC} Try it:"
  dim "pilot ask list all running docker containers"
  dim "pilot explain 'git rebase -i HEAD~3'"
  dim "pilot run compress the dist folder"
  nl
}

main