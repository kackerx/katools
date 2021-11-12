package main

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
	"time"
)

type model int

func (m model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Duration(time.Second), func(t time.Time) tea.Msg {
		return time.Time(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case time.Time:
		m -= 1
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
    return fmt.Sprintf("quit %d", m)
}

func main() {
 	p := tea.NewProgram(model(5), tea.WithAltScreen())
	if err := p.Start(); err != nil {
        panic(err)
	}
}