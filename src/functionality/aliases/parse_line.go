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

	eqIdx := strings.Index(nameCmd, "=")
	if eqIdx == -1 {
		return
	}
	name := nameCmd[:eqIdx]
	rest := nameCmd[eqIdx+1:]

	var q byte
	if len(rest) > 0 && (rest[0] == '"' || rest[0] == '\'') {
		q = rest[0]
		rest = rest[1:]
	}
	command := rest
	if q != 0 {
		command = strings.TrimSuffix(rest, string(q))
	}

	return types.AliasEntry{Name: name, Command: command, Desc: desc, Category: category}, true
}
