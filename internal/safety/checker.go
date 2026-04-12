package safety

import "strings"

type Level int

const (
	Safe    Level = iota
	Warning       // Dikkat gerektirir
	Danger        // Onay zorunlu
)

type Result struct {
	Level   Level
	Reason  string
}

var dangerPatterns = []string{
	"rm -rf",
	"rm -fr",
	"dd if=",
	"mkfs",
	":(){:|:&};:",  // fork bomb
	"> /dev/sda",
	"chmod -R 777",
	"DROP TABLE",
	"DROP DATABASE",
	"TRUNCATE",
}

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

func Check(cmd string) Result {
	lower := strings.ToLower(cmd)

	for _, p := range dangerPatterns {
		if strings.Contains(lower, strings.ToLower(p)) {
			return Result{
				Level:  Danger,
				Reason: "Bu komut geri alınamaz hasara yol açabilir: " + p,
			}
		}
	}

	for _, p := range warningPatterns {
		if strings.Contains(lower, strings.ToLower(p)) {
			return Result{
				Level:  Warning,
				Reason: "Bu komut dikkat gerektirir: " + p,
			}
		}
	}

	return Result{Level: Safe}
}