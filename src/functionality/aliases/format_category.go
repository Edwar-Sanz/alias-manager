package aliases

import (
	"am/src/functionality/styles"
	"am/src/types"
	"strings"

	"charm.land/lipgloss/v2"
)

func FormatCategory(cat string, entries []types.AliasEntry) string {
	var box strings.Builder
	box.WriteString(
		styles.SectionStyle.
			UnderlineStyle(lipgloss.UnderlineCurly).
			Render(" ➜ "+cat) + "\n\n")
	for _, e := range entries {
		box.WriteString("  " + styles.TitleStyle.Render(e.Name) + "  " + styles.DescriptionStyle.Render(e.Desc) + "\n")
		box.WriteString("  " + styles.CommandStyle.Render(e.Command) + "\n\n")
	}
	return box.String()
}
