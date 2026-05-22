package remove

import (
	"am/src/functionality/styles"
	"strings"
)

func buildDescription() string {
	var body strings.Builder

	title := styles.TitleStyle.Render("\nRemove alias")
	subtitle := styles.DescriptionStyle.Render("Remove an existing alias by name.")

	body.WriteString(title + "\n")
	body.WriteString(subtitle + "\n\n")

	body.WriteString(styles.SectionStyle.Render("▸ Usage") + "\n")
	body.WriteString("    " + styles.CommandStyle.Render("am r <alias>") + "\n\n")

	body.WriteString(styles.SectionStyle.Render("▸ Arguments") + "\n")
	body.WriteString("    " + styles.TitleStyle.Render("<alias>") + "  " + styles.DescriptionStyle.Render("Name of the alias to remove") + "\n\n")

	return body.String()
}
