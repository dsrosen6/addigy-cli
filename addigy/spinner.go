package addigy

import (
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

const addigyBlue = "#00C7CC"

var spinnerTheme = lipgloss.NewStyle().Foreground(lipgloss.Color(addigyBlue))

func RunWithSpinner(title string, action func()) {
	_ = spinner.New().Type(spinner.Line).Style(spinnerTheme).Title(title).Action(action).Run()
}
