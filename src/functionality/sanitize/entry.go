package sanitize

import "am/src/types"

func Entry(entry types.AliasEntry) types.AliasEntry {
	return types.AliasEntry{
		Name:     Name(entry.Name),
		Command:  Command(entry.Command),
		Category: Category(entry.Category),
		Desc:     Desc(entry.Desc),
	}
}
