package aliases

import (
	"am/src/functionality/styles"
	"strings"
)

func FormatCategories(allCategories []string) string {
	var body strings.Builder
	for _, cat := range allCategories {
		body.WriteString(styles.SectionStyle.Render("▸ "+cat) + "\n")
	}
	return body.String()
}
