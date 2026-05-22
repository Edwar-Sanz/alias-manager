package list

import (
	"am/src/functionality/styles"
	"strings"
)

func buildDescription() string {
	var body strings.Builder

	title := styles.TitleStyle.Render("\nList aliases")
	subtitle := styles.DescriptionStyle.Render("List your aliases, optionally filtered by category.")

	body.WriteString(title + "\n")
	body.WriteString(subtitle + "\n\n")

	body.WriteString(styles.SectionStyle.Render("▸ Usage") + "\n")
	for _, cmd := range []string{
		"am l",
		"am l -c <category>",
		"am l -C",
	} {
		body.WriteString("    " + styles.CommandStyle.Render(cmd) + "\n")
	}
	body.WriteString("\n")

	body.WriteString(styles.SectionStyle.Render("▸ Flags") + "\n")
	flags := []struct{ name, desc string }{
		{"-c, --category  ", "Filter aliases by category"},
		{"-C, --categories", "List category names only"},
	}
	for _, f := range flags {
		body.WriteString("    " + styles.TitleStyle.Render(f.name) + "  " + styles.DescriptionStyle.Render(f.desc) + "\n")
	}
	body.WriteString("\n")

	return body.String()
}
