package main

import (
	"fmt"
	"github.com/bndr/gotabulate"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type model struct {
	todos    [][]string
	cursor   int
	selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "e":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "n":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "enter", " ":
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

func (m model) View() string {
	var s string
	for i, item := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// Render row
		//s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item)
		item[0] += cursor + checked + item[0]
		
		// footer
	}
	fmt.Println(m.todos)
	tabulate := gotabulate.Create(m.todos)
	s = tabulate.Render("simple")
    s += "Press q to quit\n"
	return s
}

func initModel() model {
	s1 := []string{"tt", "hehe"}
	s2 := []string{"as", "yyyy"}
	return model{
		todos:    [][]string{s1, s2},
		selected: make(map[int]struct{}),
	}
}

func main() {
	
	
	
	//tabulate.SetHeaders([]string{"Cost", "Status"})

	//fmt.Println(tabulate.Render("simple"))
    p := tea.NewProgram(initModel())
    if err := p.Start(); err != nil {
       fmt.Printf("%v", err)
       os.Exit(1)
    }
}
