package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	checkInterval = 24 * time.Hour
	stampFile     = ".update_check"
	releaseAPI    = "https://api.github.com/repos/muhofy/Pilot-cli/releases/latest"
)

// CheckInBackground silently checks for a new version once per day.
// If a newer version is found, prints a one-line hint to stdout.
func CheckInBackground(currentVersion string) {
	if currentVersion == "dev" {
		return
	}

	// Rate-limit: only check once per day
	if !shouldCheck() {
		return
	}

	go func() {
		latest, err := fetchLatest()
		if err != nil {
			return
		}
		writeStamp()

		current := normalize(currentVersion)
		remote  := normalize(latest)

		if current == remote {
			return
		}

		// Compare semver numerically
		if isNewer(remote, current) {
			fmt.Println()
			color.Yellow("  💡 New version available: %s → run: pilot update\n", latest)
		}
	}()
}

// shouldCheck returns true if 24h have passed since last check.
func shouldCheck() bool {
	path, err := stampPath()
	if err != nil {
		return true
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return true
	}
	ts, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
	if err != nil {
		return true
	}
	return time.Since(time.Unix(ts, 0)) > checkInterval
}

// writeStamp records the current timestamp.
func writeStamp() {
	path, err := stampPath()
	if err != nil {
		return
	}
	os.WriteFile(path, []byte(strconv.FormatInt(time.Now().Unix(), 10)), 0600)
}

func stampPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".pilot", stampFile), nil
}

func fetchLatest() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(releaseAPI)
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
	return data.TagName, nil
}

func normalize(v string) string {
	return strings.TrimPrefix(strings.TrimSpace(v), "v")
}

// isNewer returns true if a > b (simple semver compare).
func isNewer(a, b string) bool {
	ap := parseSemver(a)
	bp := parseSemver(b)
	for i := 0; i < 3; i++ {
		if ap[i] > bp[i] { return true  }
		if ap[i] < bp[i] { return false }
	}
	return false
}

func parseSemver(v string) [3]int {
	parts := strings.SplitN(v, ".", 3)
	var out [3]int
	for i, p := range parts {
		if i >= 3 { break }
		n, _ := strconv.Atoi(strings.Split(p, "-")[0])
		out[i] = n
	}
	return out
}