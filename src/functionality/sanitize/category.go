package sanitize

import "strings"

func Category(s string) string {
	return identifierRe.ReplaceAllString(strings.TrimSpace(s), "")
}
