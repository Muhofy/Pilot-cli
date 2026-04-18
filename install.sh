#!/bin/bash
set -e

# ─────────────────────────────────────────────────────────────
#  pilot — Installer
#  curl -fsSL https://raw.githubusercontent.com/muhofy/Pilot-cli/main/install.sh | bash
# ─────────────────────────────────────────────────────────────

APP="pilot"
REPO="muhofy/Pilot-cli"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"

R='\033[0;31m'
G='\033[0;32m'
C='\033[0;36m'
DIM='\033[2m'
BOLD='\033[1m'
NC='\033[0m'

info()    { echo -e "${C}  ❯ ${NC}$1"; }
success() { echo -e "${G}  ✓ ${NC}$1"; }
error()   { echo -e "${R}  ✗ ${NC}$1"; exit 1; }
nl()      { echo ""; }

# ── Banner ────────────────────────────────────────────────────
echo -e "${C}"
echo "  ██████╗ ██╗██╗      ██████╗ ████████╗"
echo "  ██╔══██╗██║██║     ██╔═══██╗╚══██╔══╝"
echo "  ██████╔╝██║██║     ██║   ██║   ██║   "
echo "  ██╔═══╝ ██║██║     ██║   ██║   ██║   "
echo "  ██║     ██║███████╗╚██████╔╝   ██║   "
echo "  ╚═╝     ╚═╝╚══════╝ ╚═════╝    ╚═╝   "
echo -e "${NC}"
echo -e "  ${BOLD}Your terminal co-pilot${NC}"
echo -e "  ${DIM}────────────────────────────────────────${NC}"
nl

# ── Detect platform ───────────────────────────────────────────
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux)  OS="linux" ;;
  darwin) OS="darwin" ;;
  *)      error "Unsupported OS: $OS" ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) error "Unsupported architecture: $ARCH" ;;
esac

BINARY="${APP}-${OS}-${ARCH}"

# ── Detect install dir ────────────────────────────────────────
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

# ── Fetch latest version ──────────────────────────────────────
info "Fetching latest release..."
if command -v curl &>/dev/null; then
  VERSION=$(curl -fsSL "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
else
  VERSION=$(wget -qO- "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
fi
[ -z "$VERSION" ] && error "Could not fetch latest version."
success "Version: ${BOLD}${VERSION}${NC}"

# ── Download ──────────────────────────────────────────────────
URL="${GITHUB_DL}/${VERSION}/${BINARY}"
TMP=$(mktemp)

info "Downloading ${BINARY}..."
if command -v curl &>/dev/null; then
  curl -fsSL "$URL" -o "$TMP"
else
  wget -qO "$TMP" "$URL"
fi

# ── Install ───────────────────────────────────────────────────
chmod +x "$TMP"
if [ "${NEEDS_SUDO}" = "1" ]; then
  sudo mv "$TMP" "${INSTALL_DIR}/${APP}"
else
  mv "$TMP" "${INSTALL_DIR}/${APP}"
fi

success "Installed to ${INSTALL_DIR}/${APP}"

# ── Done ──────────────────────────────────────────────────────
nl
echo -e "  ${DIM}────────────────────────────────────────${NC}"
echo -e "  ${G}${BOLD}pilot installed! 🎉${NC}"
echo -e "  ${DIM}────────────────────────────────────────${NC}"
nl
echo -e "  ${C}❯${NC} Run setup to configure:"
echo -e "    ${DIM}pilot setup${NC}"
nl