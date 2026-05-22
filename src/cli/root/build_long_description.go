package root

import (
	"am/src/functionality/styles"
	"strings"
)

func BuildLongDescription() string {

	sections := []struct {
		title    string
		commands []string
	}{
		{
			"List aliases",
			[]string{
				"am l",
				"am l -c <category>",
				"am l -C",
			},
		},
		{
			"Add or update an alias",
			[]string{
				"am a <alias> <command>",
				"am a <alias> <command> <category>",
				"am a <alias> <command> <category> <description>",
			},
		},
		{
			"Remove an alias",
			[]string{"am r <alias>"},
		},
	}

	var body strings.Builder
	title := styles.TitleStyle.Render("\n" + "Alias Manager")
	subtitle := styles.DescriptionStyle.Render("A CLI to manage your frequently used command aliases.")

	body.WriteString(title + "\n")
	body.WriteString(subtitle + "\n\n")

	for _, s := range sections {
		body.WriteString(styles.SectionStyle.Render("▸ "+s.title) + "\n")
		for _, cmd := range s.commands {
			body.WriteString("    " + styles.CommandStyle.Render(cmd) + "\n")
		}
		body.WriteString("\n")
	}

	return body.String()
}
