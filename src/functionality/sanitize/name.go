package sanitize

import (
	"regexp"
	"strings"
)

var identifierRe = regexp.MustCompile(`[^a-zA-Z0-9\-_]`)

func Name(s string) string {
	return identifierRe.ReplaceAllString(strings.TrimSpace(s), "")
}
