#!/bin/bash
set -e

# ─────────────────────────────────────────────────────────────
#  pilot — Installer
#  curl -fsSL https://raw.githubusercontent.com/muhofy/pilot/main/install.sh | bash
# ─────────────────────────────────────────────────────────────

APP="pilot"
REPO="muhofy/pilot"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"
ENV_FILE="$HOME/.pilot_env"
CONFIG_FILE="$HOME/.pilot/config.json"

# ── ANSI ─────────────────────────────────────────────────────
R='\033[0;31m'   # red
G='\033[0;32m'   # green
Y='\033[1;33m'   # yellow
C='\033[0;36m'   # cyan
B='\033[1;34m'   # blue
W='\033[0;37m'   # white
DIM='\033[2m'
BOLD='\033[1m'
NC='\033[0m'

# ── Helpers ───────────────────────────────────────────────────
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

# ── Arrow-key select (pure bash) ─────────────────────────────
# Usage: arrow_select "Question" "opt1" "opt2" ...
# Sets ARROW_RESULT to chosen value, ARROW_INDEX to index
arrow_select() {
  local question="$1"
  shift
  local options=("$@")
  local n=${#options[@]}
  local cursor=0

  # Save terminal state
  local old_stty
  old_stty=$(stty -g 2>/dev/null) || true

  # Hide cursor, raw mode
  tput civis 2>/dev/null || true
  stty raw -echo 2>/dev/null || true

  _restore() {
    stty "$old_stty" 2>/dev/null || true
    tput cnorm 2>/dev/null || true
  }
  trap _restore EXIT INT TERM

  _draw() {
    # Clear drawn lines
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
    local key
    key=$(dd bs=3 count=1 2>/dev/null)

    case "$key" in
      # Arrow Up / k
      $'\x1b\x5b\x41'|k)
        (( cursor > 0 )) && (( cursor-- )) || cursor=$(( n - 1 ))
        _draw ;;
      # Arrow Down / j
      $'\x1b\x5b\x42'|j)
        (( cursor < n-1 )) && (( cursor++ )) || cursor=0
        _draw ;;
      # Enter
      $'\x0d'|$'\x0a'|'')
        break ;;
      # Ctrl+C / ESC
      $'\x03'|$'\x1b')
        _restore
        nl
        error "Aborted." ;;
    esac
  done

  _restore
  trap - EXIT INT TERM

  nl
  ARROW_RESULT="${options[$cursor]}"
  ARROW_INDEX=$cursor
}

# ── Progress bar ──────────────────────────────────────────────
# Usage: download_with_progress <url> <output_file>
download_with_progress() {
  local url="$1"
  local out="$2"
  local tmp_file
  tmp_file=$(mktemp)

  nl
  echo -e "  ${C}Downloading pilot...${NC}"
  nl

  local bar_width=40
  local filled=0
  local spinner_chars="⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
  local spin_idx=0

  _draw_bar() {
    local pct=$1
    local speed=$2
    filled=$(( pct * bar_width / 100 ))
    local empty=$(( bar_width - filled ))
    local bar=""
    for (( i=0; i<filled; i++ )); do bar+="█"; done
    for (( i=0; i<empty;  i++ )); do bar+="░"; done
    local spin="${spinner_chars:$spin_idx:1}"
    spin_idx=$(( (spin_idx + 1) % ${#spinner_chars} ))
    printf "\r  ${C}%s${NC} [${G}%s${NC}] ${BOLD}%3d%%${NC}  ${DIM}%s${NC}  " \
      "$spin" "$bar" "$pct" "$speed"
  }

  # Try curl with progress parsing
  if command -v curl &>/dev/null; then
    curl -fsSL \
      --progress-bar \
      -w "%{speed_download}" \
      "$url" -o "$out" 2>"$tmp_file" &
    local curl_pid=$!

    while kill -0 $curl_pid 2>/dev/null; do
      # Parse curl progress from stderr (# chars = percent)
      local progress
      progress=$(wc -c < "$tmp_file" 2>/dev/null || echo 0)
      local pct=$(( (progress % 101) ))
      _draw_bar "$pct" "..."
      sleep 0.1
    done

    wait $curl_pid
    printf "\r  ${G}✓${NC} [$(printf '█%.0s' $(seq 1 $bar_width))] ${BOLD}100%%${NC}            \n"

  elif command -v wget &>/dev/null; then
    wget -q --show-progress "$url" -O "$out" 2>&1 | while IFS= read -r line; do
      local pct
      pct=$(echo "$line" | grep -oP '\d+(?=%)' | tail -1 || echo 0)
      local speed
      speed=$(echo "$line" | grep -oP '[\d.]+[KMG]B/s' | tail -1 || echo "")
      [ -n "$pct" ] && _draw_bar "$pct" "$speed"
    done
    printf "\r  ${G}✓${NC} [$(printf '█%.0s' $(seq 1 $bar_width))] ${BOLD}100%%${NC}            \n"
  else
    error "curl or wget is required"
  fi

  rm -f "$tmp_file"
  nl
}

# ── Detect platform ───────────────────────────────────────────
detect_platform() {
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  ARCH=$(uname -m)

  case "$OS" in
    linux)  OS="linux" ;;
    darwin) OS="darwin" ;;
    *)      error "Unsupported OS: $OS" ;;
  esac

  case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) error "Unsupported architecture: $ARCH" ;;
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
  info "Fetching latest release..."
  if command -v curl &>/dev/null; then
    VERSION=$(curl -fsSL "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  else
    VERSION=$(wget -qO- "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  fi
  [ -z "$VERSION" ] && error "Could not fetch latest version."
  success "Latest version: ${BOLD}${VERSION}${NC}"
}

# ── Write config ──────────────────────────────────────────────
write_config() {
  local lang="$1"
  local model="$2"
  mkdir -p "$HOME/.pilot"
  cat > "$CONFIG_FILE" <<EOF
{
  "lang": "${lang}",
  "model": "${model}"
}
EOF
}

# ── Write API key ─────────────────────────────────────────────
write_api_key() {
  local key="$1"
  echo "export OPENROUTER_API_KEY=${key}" > "$ENV_FILE"
  chmod 600 "$ENV_FILE"
}

# ── Add to shell profile ──────────────────────────────────────
add_to_profile() {
  local shell="$1"
  local profile=""

  case "$shell" in
    zsh)  profile="$HOME/.zshrc" ;;
    fish) profile="$HOME/.config/fish/config.fish" ;;
    *)    profile="$HOME/.bashrc" ;;
  esac

  local source_line='[ -f "$HOME/.pilot_env" ] && source "$HOME/.pilot_env"'
  local completion_line=""

  case "$shell" in
    zsh)  completion_line='eval "$(pilot completion zsh)"' ;;
    fish) completion_line='pilot completion fish | source' ;;
    *)    completion_line='eval "$(pilot completion bash)"' ;;
  esac

  # Avoid duplicates
  if [ -f "$profile" ]; then
    grep -qF '.pilot_env' "$profile" || echo "$source_line"     >> "$profile"
    grep -qF 'pilot completion' "$profile" || echo "$completion_line" >> "$profile"
  else
    echo "$source_line"     >> "$profile"
    echo "$completion_line" >> "$profile"
  fi

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
  arrow_select "Select your language" \
    "English" \
    "Turkish / Türkçe"

  case "$ARROW_INDEX" in
    0) LANG_CODE="en" ;;
    1) LANG_CODE="tr" ;;
    *) LANG_CODE="en" ;;
  esac
  success "Language: ${BOLD}${LANG_CODE}${NC}"
  nl

  # ── Step 2: AI Model ───────────────────────────────────────
  echo -e "  ${BOLD}Step 2 of 4${NC}  ${DIM}AI Model${NC}"
  nl
  arrow_select "Select AI model" \
    "DeepSeek v3.1     (recommended)" \
    "Llama 4 Maverick" \
    "Qwen3 235B" \
    "Gemma 3 27B" \
    "Auto (fallback)"

  case "$ARROW_INDEX" in
    0) MODEL="deepseek/deepseek-chat-v3.1:free" ;;
    1) MODEL="meta-llama/llama-4-maverick:free" ;;
    2) MODEL="qwen/qwen3-235b-a22b:free" ;;
    3) MODEL="google/gemma-3-27b-it:free" ;;
    4) MODEL="" ;;
  esac
  success "Model: ${BOLD}${MODEL:-auto}${NC}"
  nl

  # ── Step 3: Shell ─────────────────────────────────────────
  echo -e "  ${BOLD}Step 3 of 4${NC}  ${DIM}Shell  ${DIM}(detected: ${DETECTED_SHELL})${NC}"
  nl
  arrow_select "Select your shell" \
    "bash" \
    "zsh" \
    "fish"

  CHOSEN_SHELL="$ARROW_RESULT"
  success "Shell: ${BOLD}${CHOSEN_SHELL}${NC}"
  nl

  # ── Step 4: API Key ───────────────────────────────────────
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
      warn "API key cannot be empty. Try again."
    elif [[ ! "$api_key" == sk-or-* ]]; then
      warn "Key doesn't look right (should start with sk-or-). Continue anyway?"
      printf "  ${Y}?${NC} [y/n]: "
      read -r confirm
      [[ "$confirm" == "y" || "$confirm" == "yes" ]] && break
    else
      break
    fi
  done

  success "API key saved"
  nl

  # ── Download ──────────────────────────────────────────────
  fetch_version
  nl

  local url="${GITHUB_DL}/${VERSION}/${BINARY}"
  local tmp_bin
  tmp_bin=$(mktemp)

  download_with_progress "$url" "$tmp_bin"
  chmod +x "$tmp_bin"

  # Install binary
  if [ "${NEEDS_SUDO}" = "1" ]; then
    info "Installing to ${INSTALL_DIR} (requires sudo)..."
    sudo mv "$tmp_bin" "${INSTALL_DIR}/${APP}"
  else
    mv "$tmp_bin" "${INSTALL_DIR}/${APP}"
  fi

  # Write config + API key + profile
  write_config "$LANG_CODE" "$MODEL"
  write_api_key "$api_key"
  add_to_profile "$CHOSEN_SHELL"

  # ── Done ──────────────────────────────────────────────────
  nl
  echo -e "  ${DIM}────────────────────────────────────────${NC}"
  echo -e "  ${G}${BOLD}pilot is ready!${NC}"
  echo -e "  ${DIM}────────────────────────────────────────${NC}"
  nl
  echo -e "  ${C}❯${NC} Reload your shell:"
  echo -e "    ${DIM}source ~/${CHOSEN_SHELL == "fish" && echo ".config/fish/config.fish" || echo ".${CHOSEN_SHELL}rc"}${NC}"
  nl
  echo -e "  ${C}❯${NC} Try it:"
  dim "pilot ask list all running docker containers"
  dim "pilot explain 'git rebase -i HEAD~3'"
  dim "pilot run compress the dist folder"
  nl
}

main