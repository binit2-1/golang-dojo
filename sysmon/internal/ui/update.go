package ui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	//key press
	case tea.KeyMsg:
		//which keys pressed?
		switch msg.String() {

		// exit program
		case "q", "ctrl+c":
			return m, tea.Quit

		}
	
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return m, nil
}
