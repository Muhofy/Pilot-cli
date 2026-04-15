#!/bin/bash
set -e

# ─────────────────────────────────────────
#  pilot — One-liner installer
#  Usage: curl -fsSL https://raw.githubusercontent.com/muhofy/pilot/main/install.sh | bash
# ─────────────────────────────────────────

APP="pilot"
REPO="muhofy/pilot"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"

# ── Colors ──────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info()    { echo -e "${CYAN}  ℹ ${NC} $1"; }
success() { echo -e "${GREEN}  ✅ ${NC} $1"; }
warn()    { echo -e "${YELLOW}  ⚠️  ${NC} $1"; }
error()   { echo -e "${RED}  ❌ ${NC} $1"; exit 1; }

# ── Banner ───────────────────────────────
echo -e "${CYAN}"
echo "  ██████╗ ██╗██╗      ██████╗ ████████╗"
echo "  ██╔══██╗██║██║     ██╔═══██╗╚══██╔══╝"
echo "  ██████╔╝██║██║     ██║   ██║   ██║   "
echo "  ██╔═══╝ ██║██║     ██║   ██║   ██║   "
echo "  ██║     ██║███████╗╚██████╔╝   ██║   "
echo "  ╚═╝     ╚═╝╚══════╝ ╚═════╝    ╚═╝   "
echo -e "${NC}"
echo "  Your terminal co-pilot"
echo ""

# ── Detect OS & Arch ─────────────────────
detect_platform() {
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  ARCH=$(uname -m)

  case "$OS" in
    linux)  OS="linux" ;;
    darwin) OS="darwin" ;;
    mingw*|msys*|cygwin*) OS="windows" ;;
    *) error "Unsupported OS: $OS" ;;
  esac

  case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) error "Unsupported architecture: $ARCH" ;;
  esac

  EXT=""
  [ "$OS" = "windows" ] && EXT=".exe"

  BINARY="${APP}-${OS}-${ARCH}${EXT}"
}

# ── Detect install dir ────────────────────
detect_install_dir() {
  # Termux
  if [ -n "$PREFIX" ] && [ -d "$PREFIX/bin" ]; then
    INSTALL_DIR="$PREFIX/bin"
    return
  fi
  # Standard Unix
  if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
    return
  fi
  # Fallback: ~/bin
  INSTALL_DIR="$HOME/bin"
  mkdir -p "$INSTALL_DIR"
  # Add to PATH hint if needed
  if ! echo "$PATH" | grep -q "$HOME/bin"; then
    warn "Add to your shell profile: export PATH=\"\$HOME/bin:\$PATH\""
  fi
}

# ── Fetch latest version tag ──────────────
fetch_version() {
  info "Fetching latest release..."
  if command -v curl &>/dev/null; then
    VERSION=$(curl -fsSL "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  elif command -v wget &>/dev/null; then
    VERSION=$(wget -qO- "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
  else
    error "curl or wget is required"
  fi

  [ -z "$VERSION" ] && error "Could not fetch latest version. Check your internet connection."
  info "Latest version: ${CYAN}${VERSION}${NC}"
}

# ── Download binary ───────────────────────
download_binary() {
  URL="${GITHUB_DL}/${VERSION}/${BINARY}"
  TMP_DIR=$(mktemp -d)
  TMP_FILE="${TMP_DIR}/${APP}"

  info "Downloading ${BINARY}..."

  if command -v curl &>/dev/null; then
    curl -fsSL "$URL" -o "$TMP_FILE" || error "Download failed: $URL"
  else
    wget -qO "$TMP_FILE" "$URL" || error "Download failed: $URL"
  fi

  chmod +x "$TMP_FILE"
  echo "$TMP_FILE"
}

# ── Install ───────────────────────────────
install_binary() {
  local src="$1"
  local dst="${INSTALL_DIR}/${APP}"

  # Need sudo for /usr/local/bin if not writable
  if [ "$INSTALL_DIR" = "/usr/local/bin" ] && [ ! -w "/usr/local/bin" ]; then
    info "Installing to /usr/local/bin (requires sudo)..."
    sudo mv "$src" "$dst"
  else
    mv "$src" "$dst"
  fi

  success "Installed to ${dst}"
}

# ── Verify ────────────────────────────────
verify_install() {
  if command -v "$APP" &>/dev/null; then
    success "pilot is ready! Run: ${CYAN}pilot setup${NC}"
  else
    warn "pilot installed but not in PATH."
    warn "Add ${INSTALL_DIR} to your PATH:"
    echo ""
    echo "    export PATH=\"${INSTALL_DIR}:\$PATH\""
    echo ""
  fi
}

# ── Already installed? ────────────────────
check_existing() {
  if command -v "$APP" &>/dev/null; then
    CURRENT=$("$APP" --version 2>/dev/null || echo "unknown")
    warn "pilot is already installed (${CURRENT})."
    echo -n "  Reinstall/update? [y/n]: "
    read -r ans
    [ "$ans" != "y" ] && [ "$ans" != "yes" ] && { info "Aborted."; exit 0; }
  fi
}

# ── Main ──────────────────────────────────
main() {
  detect_platform
  detect_install_dir
  check_existing
  fetch_version
  TMP_FILE=$(download_binary)
  install_binary "$TMP_FILE"
  verify_install

  echo ""
  echo -e "  ${CYAN}Next steps:${NC}"
  echo "    1. pilot setup       → configure your API key"
  echo "    2. pilot ask <query> → generate a command"
  echo "    3. pilot --help      → show all commands"
  echo ""
}

main