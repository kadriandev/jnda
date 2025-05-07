package visualizer

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kadriandev/jnda/model"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title    string
	desc     string
	status   string
	due_date time.Time
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type _model struct {
	table table.Model
}

func (m _model) Init() tea.Cmd {
	return nil
}

func (m _model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m _model) View() string {
	return docStyle.Render(m.table.View())
}

func ViewTasks(tasks []model.Task) {
	cols := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Title", Width: 10},
		{Title: "Desc", Width: 40},
		{Title: "Status", Width: 10},
	}

	rows := []table.Row{}
	for _, task := range tasks {
		row := []string{strconv.FormatInt(task.Id, 10), task.Title, task.Description, task.Status}
		rows = append(rows, row)
	}
	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := _model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
