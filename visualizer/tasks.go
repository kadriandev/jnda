package visualizer

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kadriandev/lazytask/model"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title string
  desc string
  status string
  due_date time.Time
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type _model struct {
	list list.Model
}

func (m _model) Init() tea.Cmd {
	return nil
}

func (m _model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m _model) View() string {
	return docStyle.Render(m.list.View())
}

func ViewTasks(tasks []model.Task) {
  var items []list.Item
  for _, task := range(tasks) {
    item := item{
      title: task.Title,
      desc: task.Description,
    }
    items = append(items, item)
  }
	m := _model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Tasks"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
