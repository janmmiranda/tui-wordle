package tui

import (
	"bytes"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/janmmiranda/tui-wordle/internal/clients"
)

type WordleModel struct {
	WordleList   list.Model
	Guesses      []WordleItem
	Width        int
	Height       int
	Tries        int
	WordleClient clients.Client
	Input        textinput.Model
	Err          error
}

type WordleItem struct {
	Word          string
	CheckedWordle clients.Wordle
}

func (w WordleItem) Title() string       { return w.Word }
func (w WordleItem) Description() string { return "" }
func (w WordleItem) FilterValue() string { return w.Word }
func (w WordleItem) View() string {
	var buffer bytes.Buffer

	if w.CheckedWordle.WasCorrect {
		buffer.WriteString(GreenStyle.Render(w.Word))
	} else {
		for _, charInfo := range w.CheckedWordle.CharacterInfo {
			var styledChar string
			if charInfo.Scoring.InWord && charInfo.Scoring.CorrectIdx {
				styledChar = GreenStyle.Render(charInfo.Char)
			} else if charInfo.Scoring.InWord {
				styledChar = YellowStyle.Render(charInfo.Char)
			} else {
				styledChar = NormalStyle.Render(charInfo.Char)
			}
			buffer.WriteString(styledChar)
		}
	}

	return buffer.String()
}

type ConvertibleToListItem interface {
	list.Item
}

func convertToItems[T ConvertibleToListItem](elements []T) []list.Item {
	items := make([]list.Item, len(elements))
	for i, element := range elements {
		items[i] = element
	}
	return items
}

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}
