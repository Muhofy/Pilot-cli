#!/bin/bash
set -e

APP="pilot"
REPO="muhofy/Pilot-cli"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"

G='\033[0;32m'
C='\033[0;36m'
R='\033[0;31m'
Y='\033[1;33m'
BOLD='\033[1m'
DIM='\033[2m'
NC='\033[0m'

ok()    { echo -e "${G}  ✓${NC}  $1"; }
info()  { echo -e "${C}  ·${NC}  $1"; }
fail()  { echo -e "${R}  ✗${NC}  $1"; exit 1; }
warn()  { echo -e "${Y}  !${NC}  $1"; }

# ── Detect OS & Arch ─────────────────────────────────────────
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux)  ;;
  darwin) ;;
  *)      fail "Unsupported OS: $OS" ;;
esac

case "$ARCH" in
  x86_64|amd64)  ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *)             fail "Unsupported architecture: $ARCH" ;;
esac

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
info "Fetching latest version..."

if command -v curl &>/dev/null; then
  VERSION=$(curl -fsSL "$GITHUB_API" \
    | grep '"tag_name"' \
    | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
elif command -v wget &>/dev/null; then
  VERSION=$(wget -qO- "$GITHUB_API" \
    | grep '"tag_name"' \
    | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
else
  fail "curl or wget is required"
fi

[ -z "$VERSION" ] && fail "Could not fetch latest version. Check your connection."

ok "Latest version: ${BOLD}$VERSION${NC}"

# ── Download binary ───────────────────────────────────────────
BINARY="${APP}-${OS}-${ARCH}"
URL="${GITHUB_DL}/${VERSION}/${BINARY}"
TMP=$(mktemp)

info "Downloading ${BINARY}..."

if command -v curl &>/dev/null; then
  curl -fsSL "$URL" -o "$TMP" || fail "Download failed: $URL"
else
  wget -qO "$TMP" "$URL" || fail "Download failed: $URL"
fi

# ── Install ───────────────────────────────────────────────────
chmod +x "$TMP"

if [ "${NEEDS_SUDO}" = "1" ]; then
  warn "Installing to /usr/local/bin (requires sudo)..."
  sudo mv "$TMP" "${INSTALL_DIR}/${APP}"
else
  mv "$TMP" "${INSTALL_DIR}/${APP}"
fi

ok "Installed to ${INSTALL_DIR}/${APP}"

# ── Verify ────────────────────────────────────────────────────
if ! command -v "$APP" &>/dev/null; then
  warn "${INSTALL_DIR} may not be in your PATH."
  warn "Add this to your shell profile:"
  echo ""
  echo "    export PATH=\"${INSTALL_DIR}:\$PATH\""
fi

# ── Done ──────────────────────────────────────────────────────
echo ""
echo -e "  ${G}${BOLD}pilot ${VERSION} installed successfully!${NC}"
echo ""
echo -e "  ${DIM}Next: configure your API key${NC}"
echo -e "  ${C}  pilot setup${NC}"
echo ""