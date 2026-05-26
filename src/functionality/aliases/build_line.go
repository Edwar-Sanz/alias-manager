package aliases

import (
	"am/src/constants"
	"am/src/types"
	"fmt"
	"strings"
)

func BuildLine(entry types.AliasEntry) string {
	q := `"`
	if strings.Contains(entry.Command, `"`) {
		q = `'`
	}
	return fmt.Sprintf(
		"%s%s=%s%s%s %s%s%s%s",
		constants.AliasPrefix,
		entry.Name,
		q,
		entry.Command,
		q,
		constants.CatMarker,
		entry.Category,
		constants.DescMarker,
		entry.Desc,
	)
}
