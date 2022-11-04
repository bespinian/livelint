package livelint

import (
	"testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/matryer/is"
)

func TestChecksModelInit(t *testing.T) {
	t.Parallel()

	t.Run("initializes the Model correctly", func(t *testing.T) {
		t.Parallel()
		is := is.New(t)

		m := InitialModel()

		is.Equal(m.listVisible, false)
		is.Equal(m.textInputVisible, false)
		is.Equal(m.yesNoInputVisible, false)
		is.Equal(m.status.context, "")
	})
}

func TestChecksModelUpdate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it            string
		inputModel    Model
		inputMsg      tea.Msg
		expectedModel Model
		expectedCmd   tea.Cmd
	}{
		{
			it:            "quits correctly",
			inputModel:    InitialModel(),
			inputMsg:      tea.KeyMsg(tea.Key{Type: tea.KeyCtrlC}),
			expectedModel: InitialModel(),
			expectedCmd:   tea.Quit,
		},
		{
			it:         "updates each step's correctly",
			inputModel: InitialModel(),
			inputMsg:   statusMsg{context: "Test context", checks: []check{{title: "Test check 1", checkResults: []CheckResult{{Message: "Msg 1", HasFailed: true}}}}},
			expectedModel: (func() Model {
				m := InitialModel()
				m.status = statusMsg{context: "Test context", checks: []check{{title: "Test check 1", checkResults: []CheckResult{{Message: "Msg 1", HasFailed: true}}}}}
				return m
			})(),
		},
		{
			it:         "sets list choices correctly",
			inputModel: InitialModel(),
			inputMsg:   listChoiceMsg{title: "test list title", items: []string{"item 1", "item 2"}, choice: make(chan int)},
			expectedModel: (func() Model {
				m := InitialModel()
				m.listVisible = true
				m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
				m.list.SetItems([]list.Item{listItem{title: "item 1"}, listItem{title: "item 2"}})
				m.list.Title = "test list title"
				return m
			})(),
			expectedCmd: nil,
		},
		{
			it:         "sets yes / no choices correctly",
			inputModel: InitialModel(),
			inputMsg:   yesNoInputMsg{question: "test question"},
			expectedModel: (func() Model {
				m := InitialModel()
				m.yesNoInputVisible = true
				m.yesNoInput.Title = "test question"
				return m
			})(),
			expectedCmd: nil,
		},
		{
			it:         "sets up text input correctly",
			inputModel: InitialModel(),
			inputMsg:   textInputMsg{question: "test question"},
			expectedModel: (func() Model {
				m := InitialModel()
				m.textInputVisible = true
				m.textInput.Placeholder = "test question"
				return m
			})(),
			expectedCmd: nil,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			m, cmd := tc.inputModel.Update(tc.inputMsg)
			resultModel, ok := m.(Model)
			is.True(ok)

			is.Equal(cmd, tc.expectedCmd)
			is.Equal(resultModel.listVisible, tc.expectedModel.listVisible) // list visible
			if resultModel.listVisible {
				is.Equal(resultModel.list.Title, tc.expectedModel.list.Title)                                                                                                   // list title matches
				is.True(stringSubsequence(extractStrings(resultModel.list.Items(), getListItemTitle, is), extractStrings(tc.expectedModel.list.Items(), getListItemTitle, is))) // all expected items present
				is.True(stringSubsequence(extractStrings(tc.expectedModel.list.Items(), getListItemTitle, is), extractStrings(resultModel.list.Items(), getListItemTitle, is))) // only expected items present
			}
			is.Equal(resultModel.textInputVisible, tc.expectedModel.textInputVisible) // text input visible
			if resultModel.textInputVisible {
				is.Equal(resultModel.textInput.Placeholder, tc.expectedModel.textInput.Placeholder) // placeholder text matches
			}
			is.Equal(resultModel.yesNoInputVisible, tc.expectedModel.yesNoInputVisible) // yes / no input visible
			if resultModel.yesNoInputVisible {
				is.Equal(resultModel.yesNoInput.Title, tc.expectedModel.yesNoInput.Title) // yes / no input title matches
			}
			is.Equal(resultModel.status.context, tc.expectedModel.status.context)                                                                                       // context text matches
			is.Equal(resultModel.error, tc.expectedModel.error)                                                                                                         // error object matches
			is.True(stringSubsequence(extractStrings(resultModel.status.checks, getCheckTitle, is), extractStrings(tc.expectedModel.status.checks, getCheckTitle, is))) // all expected checks present
			is.True(stringSubsequence(extractStrings(tc.expectedModel.status.checks, getCheckTitle, is), extractStrings(resultModel.status.checks, getCheckTitle, is))) // only expected checks present
			for i, check := range resultModel.status.checks {
				is.True(stringSubsequence(extractStrings(check.checkResults, getCheckResultTitle, is), extractStrings(tc.expectedModel.status.checks[i].checkResults, getCheckResultTitle, is))) // all expected checks have all expected results
				is.True(stringSubsequence(extractStrings(tc.expectedModel.status.checks[i].checkResults, getCheckResultTitle, is), extractStrings(check.checkResults, getCheckResultTitle, is))) // all expected checks have only the expected results
			}
		})
	}
}

func stringSubsequence(s1, s2 []string) bool {
	result := true
	if len(s1) > len(s2) {
		result = false
	} else {
		for i, msg := range s1 {
			if s2[i] != msg {
				result = false
				break
			}
		}
	}
	return result
}

func extractStrings[T interface{}](l []T, f func(element T) (string, bool), is *is.I) []string {
	result := []string{}
	for _, t := range l {
		str, ok := f(t)
		if !ok {
			is.Fail()
		}
		result = append(result, str)
	}
	return result
}

func getListItemTitle(item list.Item) (string, bool) {
	li, ok := item.(listItem)
	if !ok {
		return "", ok
	}
	return li.Title(), ok
}

func getCheckTitle(check check) (string, bool) {
	return check.title, true
}

func getCheckResultTitle(result CheckResult) (string, bool) {
	return result.Message, true
}
