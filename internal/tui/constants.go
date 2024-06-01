package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	GreenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	YellowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00"))
	NormalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
)

type WordleDelegate struct{}

func (d WordleDelegate) Height() int                               { return 1 }
func (d WordleDelegate) Spacing() int                              { return 0 }
func (d WordleDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d WordleDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	wi, ok := listItem.(WordleItem)
	if !ok {
		return
	}
	str := wi.View()

	fn := NormalStyle.Render

	fmt.Fprint(w, fn(str))
}
