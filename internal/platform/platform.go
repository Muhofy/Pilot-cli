package platform

import (
	"os"
	"runtime"
	"strings"
)

// IsWindows returns true on Windows.
func IsWindows() bool { return runtime.GOOS == "windows" }

// IsMacOS returns true on macOS.
func IsMacOS() bool { return runtime.GOOS == "darwin" }

// IsLinux returns true on Linux (including Termux/WSL).
func IsLinux() bool { return runtime.GOOS == "linux" }

// IsTermux returns true when running inside Termux on Android.
func IsTermux() bool {
	return IsLinux() && os.Getenv("PREFIX") != "" &&
		strings.Contains(os.Getenv("PREFIX"), "com.termux")
}

// IsWSL returns true when running inside Windows Subsystem for Linux.
func IsWSL() bool {
	if !IsLinux() {
		return false
	}
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	lower := strings.ToLower(string(data))
	return strings.Contains(lower, "microsoft") || strings.Contains(lower, "wsl")
}

// Shell returns the current shell name (bash, zsh, fish, powershell, cmd).
func Shell() string {
	if IsWindows() {
		// Check if running in PowerShell
		if os.Getenv("PSModulePath") != "" {
			return "powershell"
		}
		return "cmd"
	}
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "bash"
	}
	parts := strings.Split(shell, "/")
	return parts[len(parts)-1]
}

// ProfilePath returns the shell profile file path for the current shell.
func ProfilePath() string {
	home, _ := os.UserHomeDir()
	switch Shell() {
	case "zsh":
		return home + "/.zshrc"
	case "fish":
		return home + "/.config/fish/config.fish"
	case "powershell":
		// $PROFILE equivalent
		if p := os.Getenv("USERPROFILE"); p != "" {
			return p + `\Documents\PowerShell\Microsoft.PowerShell_profile.ps1`
		}
		return ""
	default:
		return home + "/.bashrc"
	}
}

// BinaryName returns the correct binary name for the current OS.
func BinaryName() string {
	if IsWindows() {
		return "pilot.exe"
	}
	return "pilot"
}

// InstallDir returns the best install directory for the current platform.
func InstallDir() (dir string, needsSudo bool) {
	if IsTermux() {
		return os.Getenv("PREFIX") + "/bin", false
	}
	if IsWindows() {
		if p := os.Getenv("LOCALAPPDATA"); p != "" {
			return p + `\pilot`, false
		}
		home, _ := os.UserHomeDir()
		return home + `\.pilot\bin`, false
	}
	if _, err := os.Stat("/usr/local/bin"); err == nil {
		if f, err := os.OpenFile("/usr/local/bin/.write_test", os.O_CREATE|os.O_WRONLY, 0600); err == nil {
			f.Close()
			os.Remove("/usr/local/bin/.write_test")
			return "/usr/local/bin", false
		}
		return "/usr/local/bin", true
	}
	home, _ := os.UserHomeDir()
	return home + "/bin", false
}