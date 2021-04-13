package util

import (
	"encoding/json"
	"strings"
	"unicode"
)

// SnakeCase converts the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func SnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}
	s := string(out)
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, "__", "_", -1)
	return s
}

func ToPtr(s string) *string {
	return &s
}

func ToJsonString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func FindString(slice []string, value string) (int, bool) {
	for i, item := range slice {
		if item == value {
			return i, true
		}

	}
	return -1, false
}

func FindStringAt(slice []string, indx int) (string, bool ) {
	if len(slice) <= indx {
		return "", false
	}
	return slice[indx], true
}