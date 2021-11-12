package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    "os"
)

func main() {
    if err := tea.NewProgram(model{}).Start(); err != nil {
        panic(err)
        os.Exit(1)
    }
}

type model struct {
	quiting bool
	full    bool
}

func (m model) Init() tea.Cmd {
	//初始化数据
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quiting = true
			return m, tea.Quit
        case " ":
            var cmd tea.Cmd
            if m.full {
                cmd = tea.ExitAltScreen
            } else {
                cmd = tea.EnterAltScreen
            }
            m.full = !m.full
            return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
    if m.quiting {
        return "bye\n"
    }
    var mode string
    if m.full {
        mode = "full"
    } else {
        mode = "no full"
    }
    return fmt.Sprintf("you are in %s\n", mode)
}
