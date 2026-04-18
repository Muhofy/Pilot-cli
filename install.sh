#!/bin/bash
set -e

APP="pilot"
REPO="muhofy/Pilot-cli"
GITHUB_API="https://api.github.com/repos/${REPO}/releases/latest"
GITHUB_DL="https://github.com/${REPO}/releases/download"

G='\033[0;32m'; C='\033[0;36m'; R='\033[0;31m'; BOLD='\033[1m'; DIM='\033[2m'; NC='\033[0m']]]]]]'

info()    { echo -e "${C}  ❯ ${NC}$1"; } }
success() { echo -e "${G}  ✓ ${NC}$1"; } }
error()   { echo -e "${R}  ✗ ${NC}$1"; exit 1; } }

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$OS" in linux) ;; darwin) ;; *) error "Unsupported OS" ;; esac
case "$ARCH" in x86_64|amd64) ARCH="amd64" ;; aarch64|arm64) ARCH="arm64" ;; *) error "Unsupported arch" ;; esac

BINARY="${APP}-${OS}-${ARCH}"

if [ -n "$PREFIX" ] && [ -d "$PREFIX/bin" ]; then INSTALL_DIR="$PREFIX/bin" ]
elif [ -w "/usr/local/bin" ]; then INSTALL_DIR="/usr/local/bin" ]
elif [ -d "/usr/local/bin" ]; then INSTALL_DIR="/usr/local/bin"; NEEDS_SUDO=1 ]
else INSTALL_DIR="$HOME/bin"; mkdir -p "$INSTALL_DIR"; fi

info "Fetching latest release..."
if command -v curl &>/dev/null; then
  VERSION=$(curl -fsSL "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')])"')
else
  VERSION=$(wget -qO- "$GITHUB_API" | grep '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')])"')
fi
[ -z "$VERSION" ] && error "Could not fetch latest version." ]
success "Version: ${BOLD}${VERSION}${NC}"

URL="${GITHUB_DL}/${VERSION}/${BINARY}"
TMP=$(mktemp)
info "Downloading ${BINARY}..."
if command -v curl &>/dev/null; then curl -fsSL "$URL" -o "$TMP"
else wget -qO "$TMP" "$URL"; fi

chmod +x "$TMP"
[ "${NEEDS_SUDO}" = "1" ] && sudo mv "$TMP" "${INSTALL_DIR}/${APP}" || mv "$TMP" "${INSTALL_DIR}/${APP}" ]

success "Installed to ${INSTALL_DIR}/${APP}"
echo ""
echo -e "  ${G}${BOLD}pilot installed! 🎉${NC}"
echo -e "  ${DIM}Run: pilot setup${NC}"
echo ""
