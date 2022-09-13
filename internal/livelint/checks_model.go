package livelint

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	width               = 120
	paddingTopBottom    = 2
	listHeight          = 8
	listPaddingTop      = 1
	headerMarginTop     = 1
	listItemPaddingLeft = 2
)

type model struct {
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
	spinnerVisible    bool
	spinner           spinner.Model
	status            statusMsg
	verbose           bool
}

func initialModel() model {
	yesNoItems := []list.Item{
		listItem{title: "Yes"},
		listItem{title: "No"},
	}

	return model{
		list:              list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		textInput:         textinput.New(),
		yesNoInput:        list.New(yesNoItems, list.NewDefaultDelegate(), len(yesNoItems)*width, len(yesNoItems)*listHeight),
		listVisible:       false,
		textInputVisible:  false,
		yesNoInputVisible: false,
		spinnerVisible:    false,
		spinner:           spinner.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	listStyle := lipgloss.NewStyle().
		Width(width).
		Margin(1, paddingTopBottom)

	m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#3273DC"))
	m.spinner.Spinner = spinner.Pulse

	doc := strings.Builder{}

	if len(m.status.context) > 0 {
		doc.WriteString(m.assembleHeaderBar())
	}

	doc.WriteString(m.assembleLists() + "\n")

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
	if m.spinnerVisible {
		spinnerStr := listStyle.Render(m.spinner.View())
		doc.WriteString(spinnerStr + "\n")
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

	case verboseMsg:
		m.verbose = msg.verbose

	case statusMsg:
		m.status = msg

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

	case showSpinnerMsg:
		if msg.showing {
			m.spinnerVisible = true
			return m, m.spinner.Tick
		}
		m.spinnerVisible = false
		return m, nil

	}

	var cmd tea.Cmd
	m.list, _ = m.list.Update(msg)
	m.yesNoInput, _ = m.yesNoInput.Update(msg)
	m.textInput, _ = m.textInput.Update(msg)
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

type summaryType int

const (
	success summaryType = iota
)

type summaryMsg struct {
	text string
	kind summaryType
}

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

type verboseMsg struct {
	verbose bool
}

type statusMsg struct {
	context string
	checks  []check
}

type check struct {
	title        string
	checkResults []CheckResult
	outcome      summaryMsg
}

func initalizeVerbose(verbose bool) verboseMsg {
	return verboseMsg{verbose: verbose}
}

func initalizeStatus(context string) statusMsg {
	return statusMsg{context: context, checks: []check{}}
}

func (s *statusMsg) StartCheck(title string) {
	s.checks = append(s.checks, check{title: title, checkResults: []CheckResult{}})
}

func (s *statusMsg) AddCheckResult(checkResult CheckResult) {
	s.checks[len(s.checks)-1].checkResults = append(s.checks[len(s.checks)-1].checkResults, checkResult)
}

func (s *statusMsg) CompleteCheck(outcome summaryMsg) {
	s.checks[len(s.checks)-1].outcome = outcome
}

type showSpinnerMsg struct {
	showing bool
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

func mapStrings[T interface{}](items []T, f func(item T) string) []string {
	result := []string{}
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

func (m model) assembleLists() string {
	var (
		gray = lipgloss.AdaptiveColor{Light: "#B2BEB5", Dark: "#818589"}

		subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
		list   = lipgloss.NewStyle().
			BorderForeground(subtle).
			Height(listHeight).
			Width(width).
			PaddingTop(listPaddingTop)

		listHeader = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderBottom(true).
				BorderForeground(subtle).
				PaddingTop(listPaddingTop).
				Render

		good      = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
		checkMark = lipgloss.NewStyle().SetString("✔").
				Foreground(good).
				PaddingRight(listItemPaddingLeft).
				String()

		bad   = lipgloss.AdaptiveColor{Light: "#EA9999", Dark: "#E06666"}
		cross = lipgloss.NewStyle().SetString("✘").
			Foreground(bad).
			PaddingRight(listItemPaddingLeft).
			String()

		action     = lipgloss.AdaptiveColor{Light: "#FFBF00", Dark: "#FFEA00"}
		arrowRight = lipgloss.NewStyle().SetString("⮕").
				Foreground(action).
				PaddingLeft(listItemPaddingLeft).
				PaddingRight(listItemPaddingLeft).
				String()

		listItem        = lipgloss.NewStyle().PaddingLeft(listItemPaddingLeft)
		listItemDetails = func(c CheckResult) string {
			details := ""
			if m.verbose && len(c.Details) > 0 {
				details = "\n\n     " + lipgloss.NewStyle().SetString(strings.Join(c.Details, "\n     ")).
					Foreground(gray).
					String() + "\n"
			}
			return details
		}
		listItemInstructions = func(c CheckResult) string {
			instructions := ""
			if c.Instructions != "" {
				instructions = "\n\n" + listItem.Copy().Render(
					arrowRight+lipgloss.NewStyle().SetString(c.Instructions).Bold(true).String(),
				) + "\n"
			}
			return instructions
		}
		listItemSuccess = func(c CheckResult) string {
			return listItem.Copy().Render(checkMark+c.Message) +
				listItemDetails(c) +
				listItemInstructions(c)
		}
		listItemError = func(c CheckResult) string {
			return listItem.Copy().Render(cross+c.Message) +
				listItemDetails(c) +
				listItemInstructions(c)
		}
	)

	summary := []string{}
	for _, check := range m.status.checks {
		summary = append(summary, listHeader(check.title))
		summary = append(summary, mapStrings(check.checkResults, func(c CheckResult) string {
			if c.HasFailed {
				return listItemError(c)
			}
			return listItemSuccess(c)
		})...)
		summary = append(summary, "")
	}

	return list.Render(
		lipgloss.JoinVertical(lipgloss.Left, summary...),
	)
}

func (m model) assembleHeaderBar() string {
	var (
		statusNugget = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Padding(0, 1)

		statusBarStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

		statusStyle = lipgloss.NewStyle().
				Inherit(statusBarStyle).
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#3273DC")).
				Padding(0, 1).
				MarginRight(1)

		modeStyle = statusNugget.Copy().
				Background(lipgloss.Color("#3273DC")).
				Align(lipgloss.Right)

		statusText = lipgloss.NewStyle().Inherit(statusBarStyle)
	)

	w := lipgloss.Width
	statusKey := statusStyle.Render("livelint")
	verbose := ""
	if m.verbose {
		verbose = "verbose"
	}
	encoding := modeStyle.Render(verbose)
	statusVal := statusText.Copy().
		Width(width - w(statusKey) - w(encoding)).
		Render(m.status.context)

	return statusBarStyle.Width(width).MarginTop(headerMarginTop).Render(lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		encoding,
	))
}
