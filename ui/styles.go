package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	ActiveHeaderColor      = "#fc5fa3"
	InactivePaneColor      = "#d0a8ff"
	DisabledFileColor      = "#6c7986"
	DirectoryColor         = "#41a1c0"
	NoConstructsColor      = "#fabd2f"
	UnsupportedFileColor   = "#928374"
	CWDColor               = "#d0a8ff"
	FilepathColor          = "#fc5fa3"
	FilepathColorTUI       = "#ffd166"
	TSElementColor         = "#41a1c0"
	DividerColor           = "#6c7986"
	DefaultForegroundColor = "#282828"
	ModeColor              = "#fc5fa3"
	HelpMsgColor           = "#83a598"
	FooterColor            = "#7c6f64"
	HTMLBackgroundColor    = "#1f1f24"
)

var (
	filePathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(FilepathColor))

	filePathStyleTUI = lipgloss.NewStyle().
				Foreground(lipgloss.Color(FilepathColorTUI))

	tsElementStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(TSElementColor))

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(DividerColor))

	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(DefaultForegroundColor))

	fileExplorerStyle = lipgloss.NewStyle().
				Width(45).
				PaddingRight(2).
				PaddingBottom(1)

	activePaneHeaderStyle = baseStyle.
				Align(lipgloss.Left).
				Bold(true).
				Background(lipgloss.Color(ActiveHeaderColor))

	inActivePaneHeaderStyle = activePaneHeaderStyle.
				Background(lipgloss.Color(InactivePaneColor))

	unsupportedFileStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(UnsupportedFileColor))

	cwdStyle = baseStyle.
			PaddingRight(0).
			Foreground(lipgloss.Color(CWDColor))

	modeStyle = baseStyle.
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color(ModeColor))

	helpMsgStyle = baseStyle.
			Bold(true).
			Foreground(lipgloss.Color(HelpMsgColor))
)
