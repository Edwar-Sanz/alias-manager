package aliases

import (
	"am/src/types"
	"strings"
)

func ParseAliases(content string) types.ParsedAliases {
	result := types.ParsedAliases{
		Categories: []string{},
		ByCategory: map[string][]types.AliasEntry{},
	}

	for line := range strings.SplitSeq(strings.TrimSpace(content), "\n") {
		entry, ok := ParseLine(line)
		if !ok {
			continue
		}
		if _, exists := result.ByCategory[entry.Category]; !exists {
			result.Categories = append(result.Categories, entry.Category)
		}
		result.ByCategory[entry.Category] = append(result.ByCategory[entry.Category], entry)
	}

	return result
}
