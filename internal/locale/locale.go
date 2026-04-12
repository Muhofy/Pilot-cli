package locale

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed lang/en_US.json
var enUS []byte

//go:embed lang/tr_TR.json
var trTR []byte

var strings_ map[string]string

// Init loads the locale based on config lang, falling back to system LANG, then en_US.
func Init(lang string) {
	if lang == "" {
		lang = detectSystemLang()
	}

	var data []byte
	switch normalise(lang) {
	case "tr":
		data = trTR
	default:
		data = enUS
	}

	if err := json.Unmarshal(data, &strings_); err != nil {
		_ = json.Unmarshal(enUS, &strings_)
	}
}

// T returns the localised string for the given key.
// Extra args are passed to fmt.Sprintf if the value contains format verbs.
func T(key string, args ...any) string {
	if strings_ == nil {
		return key
	}

	val, ok := strings_[key]
	if !ok {
		return key
	}
	if len(args) == 0 {
		return val
	}
	return fmt.Sprintf(val, args...)
}

func detectSystemLang() string {
	for _, env := range []string{"LANG", "LANGUAGE", "LC_ALL"} {
		if v := os.Getenv(env); v != "" {
			return v
		}
	}
	return "en"
}

func normalise(lang string) string {
	lang = strings.ToLower(lang)
	lang = strings.Split(lang, "_")[0]
	lang = strings.Split(lang, ".")[0]
	lang = strings.Split(lang, "-")[0]
	return lang
}