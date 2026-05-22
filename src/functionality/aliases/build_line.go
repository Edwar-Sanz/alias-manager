package aliases

import (
	"am/src/constants"
	"am/src/types"
	"fmt"
)

func BuildLine(entry types.AliasEntry) string {
	return fmt.Sprintf(
		"%s%s%s%s\" %s%s%s%s",
		constants.AliasPrefix,
		entry.Name,
		constants.AliasSeparator,
		entry.Command,
		constants.CatMarker,
		entry.Category,
		constants.DescMarker,
		entry.Desc,
	)
}
