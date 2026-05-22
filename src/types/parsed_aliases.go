package types

type ParsedAliases struct {
	Categories []string
	ByCategory map[string][]AliasEntry
}
