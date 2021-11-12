package main

import (
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    focusIndex int
    inputs []textinput.Model
}

func initModel() model {
    m := model{
        inputs: make([]textinput.Model, 3),
    }
    
    t := textinput.Model{}
    for i := range m.inputs {
        switch i {
        case 0:
            t.Placeholder = "name"
        case 1:
            t.Placeholder = "email"
        case 2:
            t.EchoMode = textinput.EchoPassword
            t.Placeholder = "pass"
        }
        m.inputs[i] = t
    }
    return m
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        }
    }
    
    return m, nil
}

func (m model) View() string {
    return ""
}