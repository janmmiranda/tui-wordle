package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/janmmiranda/tui-wordle/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialWordleModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
