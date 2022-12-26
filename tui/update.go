package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m simpleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k", "^[[A":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", "^[[B":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "l", "^[[C", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}

	}
	return m, nil
}
