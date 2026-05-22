package styles

import "charm.land/lipgloss/v2"

var BoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Padding(1, 2)
