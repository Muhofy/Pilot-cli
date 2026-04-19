package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/muhofy/pilot/internal/ui"
)

const (
	releaseAPI  = "https://api.github.com/repos/muhofy/Pilot-cli/releases/latest"
	installURL  = "https://raw.githubusercontent.com/muhofy/Pilot-cli/main/install.sh"
)

// Update checks for a new version and installs it if confirmed.
func Update(currentVersion string) {
	sp := ui.NewSpinner("Checking for updates...")
	sp.Start()

	latest, err := fetchLatestVersion()
	sp.Stop()

	if err != nil {
		ui.Error("Could not check for updates: " + err.Error())
		return
	}

	current := normalizeVersion(currentVersion)
	remote  := normalizeVersion(latest)

	if current == "dev" {
		color.Yellow("  ⚠️  Development build — cannot compare versions.")
		color.White("  Latest release: %s\n", latest)
		return
	}

	if current == remote {
		ui.Success("Already up to date: " + currentVersion)
		return
	}

	color.Cyan("\n  📦 Update available!\n")
	color.White("  Current : %s\n", currentVersion)
	color.White("  Latest  : %s\n", latest)
	fmt.Println()

	res := ui.Select("Install update?", []ui.Option{
		{Label: "Yes, update now", Value: "yes"},
		{Label: "No, skip",        Value: "no"},
	})

	if res.Value != "yes" {
		color.White("Skipped.")
		return
	}

	runInstallScript()
}

// fetchLatestVersion calls the GitHub releases API and returns the tag name.
func fetchLatestVersion() (string, error) {
	resp, err := http.Get(releaseAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	if data.TagName == "" {
		return "", fmt.Errorf("empty tag_name in response")
	}
	return data.TagName, nil
}

// runInstallScript downloads and runs install.sh.
func runInstallScript() {
	ui.Loading("Downloading latest version...")

	resp, err := http.Get(installURL)
	if err != nil {
		ui.Error("Download failed: " + err.Error())
		return
	}
	defer resp.Body.Close()

	tmp, err := os.CreateTemp("", "pilot-install-*.sh")
	if err != nil {
		ui.Error("Could not create temp file: " + err.Error())
		return
	}
	defer os.Remove(tmp.Name())

	if _, err := io.Copy(tmp, resp.Body); err != nil {
		ui.Error("Write failed: " + err.Error())
		return
	}
	tmp.Close()

	if err := os.Chmod(tmp.Name(), 0755); err != nil {
		ui.Error("chmod failed: " + err.Error())
		return
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		ui.Error("Auto-update is not supported on Windows. Download manually from:")
		color.Cyan("  https://github.com/muhofy/Pilot-cli/releases/latest\n")
		return
	}
	cmd = exec.Command("bash", tmp.Name())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin  = os.Stdin

	if err := cmd.Run(); err != nil {
		ui.Error("Update failed: " + err.Error())
	}
}

// normalizeVersion strips the leading 'v' for comparison.
func normalizeVersion(v string) string {
	return strings.TrimPrefix(strings.TrimSpace(v), "v")
}