package sanitize

import "strings"

func Command(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, `"`)
	return strings.TrimSpace(s)
}
