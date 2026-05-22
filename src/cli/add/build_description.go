package add

import (
	"am/src/functionality/styles"
	"strings"
)

func buildDescription() string {
	var body strings.Builder

	title := styles.TitleStyle.Render("\nAdd alias")
	subtitle := styles.DescriptionStyle.Render("Add a new alias or update an existing one.")

	body.WriteString(title + "\n")
	body.WriteString(subtitle + "\n\n")

	body.WriteString(styles.SectionStyle.Render("▸ Usage") + "\n")
	for _, cmd := range []string{
		"am a <alias> <command>",
		"am a <alias> <command> <category>",
		"am a <alias> <command> <category> <description>",
	} {
		body.WriteString("    " + styles.CommandStyle.Render(cmd) + "\n")
	}
	body.WriteString("\n")

	body.WriteString(styles.SectionStyle.Render("▸ Arguments") + "\n")
	args := []struct{ name, desc string }{
		{"<alias>      ", "Name for the alias"},
		{"<command>    ", "Shell command to run"},
		{"<category>   ", "Group to organize aliases  " + styles.DescriptionStyle.Render("(optional)")},
		{"<description>", "Short description          " + styles.DescriptionStyle.Render("(optional)")},
	}
	for _, a := range args {
		body.WriteString("    " + styles.TitleStyle.Render(a.name) + "  " + styles.DescriptionStyle.Render(a.desc) + "\n")
	}
	body.WriteString("\n")

	return body.String()
}
