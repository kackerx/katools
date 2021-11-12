package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"net/http"
)

type video struct {
	Name   string
	Author string
	Actor  string
	Img    string
	Desc   string
	Id     string
	Note   string
}

type model struct {
	repos   []*video
	err     error
	spinner spinner.Model
}

func NewModel() model {
	sp := spinner.NewModel()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#800080"))
	return model{
		spinner: sp,
	}
}

var (
	cyan  = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
	green = lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32"))
	gray  = lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
	gold  = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
)

func (m model) Init() tea.Cmd {
	// spinner通过tick触发其改变状态, 需要init中返回tick的cmd
	return tea.Batch(spinner.Tick, fetchVideo)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			fmt.Println("enter")
			return m, nil
		}
	case errMsg:
		m.err = msg
		return m, nil
	case []*video:
		m.repos = msg
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	var s string
	if m.err != nil {
		s = fmt.Sprintf("err: %s", m.err)
	} else if len(m.repos) > 0 {
		for _, v := range m.repos {
			s += gray.Render(v.Name) + "\n"
		}
	} else {
		s = m.spinner.View() + green.Render("fetching...")
	}

	s += "\n"
	s += cyan.Render("q to exit")
	return s
}

type errMsg struct{ error }

func fetchVideo() tea.Msg {
	videos, err := getVideo("https://www.hehedy.com/type/1-1.html")
	if err != nil {
		return errMsg{err}
	}
	return videos
}

func getVideo(url string) ([]*video, error) {
	ret := make([]*video, 0, 1)
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header["user-agent"] = []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Mobile/15E148 Safari/604.1"}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".stui-vodlist__box").Each(func(i int, s *goquery.Selection) {
		var video video
		img, ok := s.Find("a").Attr("data-original")
		if ok {
			video.Img = img
		}

		title := s.Find("h4 a").Text()
		video.Name = title

		actor := s.Find("p").Text()
		video.Actor = actor

		ret = append(ret, &video)
	})

	return ret, err
}
