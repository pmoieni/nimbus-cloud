package nbc

import "strings"

func isEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
