#!/bin/bash
set -e

APP="pilot"
CMD="./cmd/pilot/main.go"
OUT="dist"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")

echo "🔨 Building $APP $VERSION..."
mkdir -p $OUT

build() {
  local os=$1 arch=$2 ext=$3
  local out="$OUT/${APP}-${os}-${arch}${ext}"
  echo "  → $out"
  CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build \
    -ldflags="-s -w -X main.Version=$VERSION" \
    -o "$out" $CMD
}

case "${1:-local}" in
  local)
    CGO_ENABLED=0 go build -ldflags="-s -w" -o $APP $CMD
    echo "✅ Built: ./$APP"

    if [ -n "$PREFIX" ]; then
      cp $APP $PREFIX/bin/$APP
      echo "✅ Installed to \$PREFIX/bin/$APP"
    elif [ -d "/usr/local/bin" ]; then
      sudo cp $APP /usr/local/bin/$APP
      echo "✅ Installed to /usr/local/bin/$APP"
    fi
    ;;

  all)
    build linux   amd64 ""
    build linux   arm64 ""
    build darwin  amd64 ""
    build darwin  arm64 ""
    build windows amd64 ".exe"
    echo "✅ All binaries in ./$OUT/"
    ;;

  clean)
    rm -rf $OUT $APP
    echo "✅ Cleaned"
    ;;

  *)
    echo "Usage: ./build.sh [local|all|clean]"
    ;;
esac