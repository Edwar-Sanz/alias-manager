package sanitize

import "strings"

func Desc(s string) string {
	return strings.TrimSpace(s)
}
