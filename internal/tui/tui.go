package tui

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/janmmiranda/tui-wordle/internal/clients"
)

func InitialWordleModel() *WordleModel {
	wordleClient := clients.NewClient(5 * time.Second)
	delegate := WordleDelegate{}
	wordleList := list.New([]list.Item{}, delegate, 0, 0)
	wordleList.Title = "Wordle"
	wordleList.SetShowStatusBar(false)
	wordleList.SetFilteringEnabled(false)
	wordleList.SetShowPagination(false)

	ti := textinput.New()
	ti.Placeholder = "Enter Guess, 6 remaining"
	ti.Focus()
	ti.CharLimit = 5
	ti.Width = 20

	w := &WordleModel{
		WordleList:   wordleList,
		Guesses:      make([]WordleItem, 0),
		Width:        40,
		Height:       20,
		Tries:        6,
		WordleClient: wordleClient,
		Input:        ti,
	}
	return w
}

func (w *WordleModel) Init() tea.Cmd {
	return textinput.Blink
}

func (w *WordleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var inputCmd, listCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "ctrl+c":
		// 	return w, tea.Quit
		// case "q":
		// 	break
		case "enter":
			if w.Tries > 0 {
				w.Err = nil
				guess := w.Input.Value()
				res, err := w.checkGuess(guess)
				if err != nil {
					w.Input.Reset()
					w.Err = err
					return w, nil
				}
				w.Guesses = append(w.Guesses, WordleItem{Word: guess, CheckedWordle: res})
				w.WordleList.InsertItem(len(w.Guesses)-1, w.Guesses[len(w.Guesses)-1])
				if res.WasCorrect {
					w.Err = errors.New("You won!")
					w.Input.Blur()
				} else {
					w.Input.Reset()
					w.Tries -= 1
					w.Input.Placeholder = fmt.Sprintf("Enter guess, %d remaining", w.Tries)
				}
			} else {
				w.Input.Reset()
				w.Err = errors.New("Game Over!")
				w.Input.Blur()
				return w, nil
			}
		}
	case tea.WindowSizeMsg:
		w.WordleList.SetSize(msg.Width, msg.Height)
	}
	w.Input, inputCmd = w.Input.Update(msg)
	w.WordleList, listCmd = w.WordleList.Update(msg)
	cmds = append(cmds, inputCmd, listCmd)
	return w, tea.Batch(cmds...)
}

func (w *WordleModel) View() string {
	var errMsg strings.Builder
	if w.Err != nil {
		fmt.Fprintf(&errMsg, "\n %s\n", w.Err.Error())
	} else {
		errMsg.WriteString("")
	}

	return lipgloss.Place(
		40,
		20,
		lipgloss.Left,
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Render(w.WordleList.View()),
			w.Input.View(),
			errMsg.String(),
		),
	)
}

func (w *WordleModel) checkGuess(guess string) (clients.Wordle, error) {
	res, err := w.WordleClient.CheckWord(guess)
	if err != nil {
		return clients.Wordle{}, errMsg{err}
	}

	return res, nil
}
