package aliases

import (
	"am/src/constants"
	"am/src/types"
	"strings"
)

func ParseLine(line string) (entry types.AliasEntry, ok bool) {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, constants.AliasPrefix) {
		return
	}
	line = strings.TrimPrefix(line, constants.AliasPrefix)

	catIdx := strings.Index(line, constants.CatMarker)
	if catIdx == -1 {
		return
	}
	nameCmd := strings.TrimSpace(line[:catIdx])
	meta := line[catIdx+len(constants.CatMarker):]

	descIdx := strings.Index(meta, constants.DescMarker)
	if descIdx == -1 {
		return
	}
	category := meta[:descIdx]
	if category == "" {
		category = constants.UnCategorizedCategory
	}
	desc := meta[descIdx+len(constants.DescMarker):]

	eqIdx := strings.Index(nameCmd, constants.AliasSeparator)
	if eqIdx == -1 {
		return
	}
	name := nameCmd[:eqIdx]
	command := strings.TrimSuffix(nameCmd[eqIdx+len(constants.AliasSeparator):], `"`)

	return types.AliasEntry{Name: name, Command: command, Desc: desc, Category: category}, true
}
