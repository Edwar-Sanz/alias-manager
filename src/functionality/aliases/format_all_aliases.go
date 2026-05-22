package aliases

import (
	"am/src/constants"
	"am/src/types"
	"strings"
)

func FormatAllAliases(parsed types.ParsedAliases) string {
	var body strings.Builder
	body.WriteString("\n" + constants.Separator + "\n")
	for _, cat := range parsed.Categories {
		body.WriteString(FormatCategory(cat, parsed.ByCategory[cat]))
		body.WriteString(constants.Separator + "\n")
	}

	return body.String()
}
