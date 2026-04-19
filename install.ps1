# ─────────────────────────────────────────────────────────────
#  pilot — Windows Installer
#  Run: irm https://raw.githubusercontent.com/muhofy/Pilot-cli/main/install.ps1 | iex
# ─────────────────────────────────────────────────────────────

$ErrorActionPreference = "Stop"

$APP     = "pilot"
$REPO    = "muhofy/Pilot-cli"
$API_URL = "https://api.github.com/repos/$REPO/releases/latest"
$DL_BASE = "https://github.com/$REPO/releases/download"

# ── Colors ───────────────────────────────────────────────────
function Write-Info  ($msg) { Write-Host "  · $msg" -ForegroundColor Cyan }
function Write-Ok    ($msg) { Write-Host "  ✓ $msg" -ForegroundColor Green }
function Write-Warn  ($msg) { Write-Host "  ! $msg" -ForegroundColor Yellow }
function Write-Fail  ($msg) { Write-Host "  ✗ $msg" -ForegroundColor Red; exit 1 }

# ── Banner ────────────────────────────────────────────────────
Write-Host ""
Write-Host "  ██████╗ ██╗██╗      ██████╗ ████████╗" -ForegroundColor Cyan
Write-Host "  ██╔══██╗██║██║     ██╔═══██╗╚══██╔══╝" -ForegroundColor Cyan
Write-Host "  ██████╔╝██║██║     ██║   ██║   ██║   " -ForegroundColor Cyan
Write-Host "  ██╔═══╝ ██║██║     ██║   ██║   ██║   " -ForegroundColor Cyan
Write-Host "  ██║     ██║███████╗╚██████╔╝   ██║   " -ForegroundColor Cyan
Write-Host "  ╚═╝     ╚═╝╚══════╝ ╚═════╝    ╚═╝   " -ForegroundColor Cyan
Write-Host ""
Write-Host "  Your terminal co-pilot" -ForegroundColor White
Write-Host "  ────────────────────────────────────────" -ForegroundColor DarkGray
Write-Host ""

# ── Detect arch ───────────────────────────────────────────────
$arch = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else {
    Write-Fail "Only 64-bit Windows is supported."
}
$BINARY = "pilot-windows-$arch.exe"

# ── Fetch latest version ──────────────────────────────────────
Write-Info "Fetching latest release..."
try {
    $release = Invoke-RestMethod -Uri $API_URL -Headers @{ "User-Agent" = "pilot-installer" }
    $VERSION = $release.tag_name
} catch {
    Write-Fail "Could not fetch latest version: $_"
}
Write-Ok "Latest version: $VERSION"

# ── Install dir ───────────────────────────────────────────────
$INSTALL_DIR = "$env:LOCALAPPDATA\pilot"
New-Item -ItemType Directory -Force -Path $INSTALL_DIR | Out-Null

# ── Download ──────────────────────────────────────────────────
$URL     = "$DL_BASE/$VERSION/$BINARY"
$OUT     = "$INSTALL_DIR\pilot.exe"
$TMP     = "$env:TEMP\pilot-download.exe"

Write-Info "Downloading $BINARY..."
try {
    Invoke-WebRequest -Uri $URL -OutFile $TMP -UseBasicParsing
} catch {
    Write-Fail "Download failed: $_"
}

Move-Item -Force $TMP $OUT
Write-Ok "Installed to $OUT"

# ── Add to PATH ───────────────────────────────────────────────
$currentPath = [System.Environment]::GetEnvironmentVariable("PATH", "User")
if ($currentPath -notlike "*$INSTALL_DIR*") {
    [System.Environment]::SetEnvironmentVariable(
        "PATH",
        "$currentPath;$INSTALL_DIR",
        "User"
    )
    Write-Ok "Added $INSTALL_DIR to PATH"
    Write-Warn "Restart your terminal for PATH changes to take effect."
} else {
    Write-Ok "$INSTALL_DIR already in PATH"
}

# ── Done ──────────────────────────────────────────────────────
Write-Host ""
Write-Host "  ────────────────────────────────────────" -ForegroundColor DarkGray
Write-Host "  pilot $VERSION installed successfully!" -ForegroundColor Green
Write-Host "  ────────────────────────────────────────" -ForegroundColor DarkGray
Write-Host ""
Write-Host "  Next: configure your API key" -ForegroundColor DarkGray
Write-Host "  > pilot setup" -ForegroundColor Cyan
Write-Host ""