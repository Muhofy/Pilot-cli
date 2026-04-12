package safety

import "strings"

// Level represents the danger level of a command.
type Level int

const (
	Safe    Level = iota
	Warning // Requires caution
	Danger  // Requires explicit confirmation
)

// Result holds the safety check outcome.
type Result struct {
	Level  Level
	Reason string
}

// dangerPatterns are commands that can cause irreversible damage.
var dangerPatterns = []string{
	"rm -rf",
	"rm -fr",
	"dd if=",
	"mkfs",
	":(){:|:&};:", // fork bomb
	"> /dev/sda",
	"chmod -R 777",
	"DROP TABLE",
	"DROP DATABASE",
	"TRUNCATE",
}

// warningPatterns are commands that require extra care.
var warningPatterns = []string{
	"rm ",
	"git reset --hard",
	"git clean -fd",
	"git push --force",
	"kill -9",
	"pkill",
	"sudo ",
	"chmod",
	"chown",
	"format",
}

// Check evaluates a command string and returns a safety Result.
func Check(cmd string) Result {
	lower := strings.ToLower(cmd)

	for _, p := range dangerPatterns {
		if strings.Contains(lower, strings.ToLower(p)) {
			return Result{
				Level:  Danger,
				Reason: "This command can cause irreversible damage: " + p,
			}
		}
	}

	for _, p := range warningPatterns {
		if strings.Contains(lower, strings.ToLower(p)) {
			return Result{
				Level:  Warning,
				Reason: "This command requires caution: " + p,
			}
		}
	}

	return Result{Level: Safe}
}