package livelint

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	isVerbose         bool
	context           string
	messages          []string
	error             error
	choice            chan int
	list              list.Model
	listVisible       bool
	textInputVisible  bool
	textInput         textinput.Model
	textResponse      chan string
	yesNoInputVisible bool
	yesNoInput        list.Model
	YesNoResponse     chan int
}

func initialModel() model {
	return model{
		messages:          []string{},
		list:              list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		textInput:         textinput.New(),
		yesNoInput:        list.New([]list.Item{listItem{title: "yes"}, listItem{title: "no"}}, list.NewDefaultDelegate(), 0, 0),
		listVisible:       false,
		textInputVisible:  false,
		yesNoInputVisible: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	const width = 72
	const paddingTopBottom = 2

	titleStyle := lipgloss.NewStyle().
		Width(width).
		MarginTop(1).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#485fc7")).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		Padding(1, 0).
		Bold(true)

	contextStyle := lipgloss.NewStyle().
		Width(width).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("#485fc7")).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1, 1).
		Foreground(lipgloss.Color("202"))

	listStyle := lipgloss.NewStyle().
		Width(width).
		Margin(1, paddingTopBottom)

	doc := strings.Builder{}
	header := titleStyle.Render("livelint")

	if len(m.context) > 0 {
		context := contextStyle.Render(m.context)
		header = lipgloss.JoinVertical(lipgloss.Center, header, context)
	}
	doc.WriteString(header + "\n")

	stepStyle := lipgloss.NewStyle().Width(width).Padding(0, paddingTopBottom)

	for _, msg := range m.messages {
		step := stepStyle.Render(msg)
		doc.WriteString(step + "\n")
	}

	if m.listVisible {
		listStr := listStyle.Render(m.list.View())
		doc.WriteString(listStr + "\n")
	}
	if m.textInputVisible {
		textInputStr := listStyle.Render(m.textInput.View())
		doc.WriteString(textInputStr + "\n")
	}
	if m.yesNoInputVisible {
		listStr := listStyle.Render(m.yesNoInput.View())
		doc.WriteString(listStr + "\n")
	}
	return doc.String()
}

// nolint:ireturn
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.listVisible {
				index := m.list.Index()
				m.choice <- index
				m.listVisible = false
				return m, nil
			}
			if m.textInputVisible {
				value := m.textInput.Value()
				m.textResponse <- value
				m.textInputVisible = false
				return m, nil
			}
			if m.yesNoInputVisible {
				index := m.yesNoInput.Index()
				m.YesNoResponse <- index
				m.yesNoInputVisible = false
				return m, nil
			}
		}

	case stepMsg:
		m.messages = append(m.messages, msg.Message)
		return m, nil

	case contextMsg:
		m.context = string(msg)
		return m, nil

	case summaryMsg:
		m.messages = append(m.messages, string(msg))
		return m, nil

	case listChoiceMsg:
		m.list = list.New(getListItems(msg.items), list.NewDefaultDelegate(), 0, 0)
		m.list.Title = msg.title
		m.choice = msg.choice
		m.listVisible = true
		return m, nil

	case textInputMsg:
		m.textInput = textinput.New()
		m.textInput.Placeholder = msg.question
		m.textInput.Focus()
		m.textInput.CharLimit = 156
		m.textInput.Width = 20
		m.textInputVisible = true
		m.textResponse = msg.value
		return m, nil

	case yesNoInputMsg:
		m.yesNoInput.Title = msg.question
		m.YesNoResponse = msg.value
		m.yesNoInputVisible = true
		return m, nil

	case errMsg:
		m.error = msg.err
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.list, _ = m.list.Update(msg)
	m.yesNoInput, _ = m.yesNoInput.Update(msg)
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

type contextMsg string

type stepMsg CheckResult

type summaryMsg string

type listChoiceMsg struct {
	title  string
	items  []string
	choice chan int
}

type textInputMsg struct {
	question string
	value    chan string
}

type yesNoInputMsg struct {
	question string
	value    chan int
}

type errMsg struct {
	err error
}

type listItem struct {
	title string
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return "" }
func (i listItem) FilterValue() string { return i.title }

func getListItems(items []string) []list.Item {
	result := []list.Item{}
	for _, itemString := range items {
		result = append(result, listItem{title: itemString})
	}
	return result
}
