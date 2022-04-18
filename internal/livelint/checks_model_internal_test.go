package livelint

import (
	"testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/matryer/is"
)

func TestChecksModelInit(t *testing.T) {
	t.Run("initializes the model correctly", func(t *testing.T) {
		t.Parallel()
		is := is.New(t)

		m := initialModel()

		is.Equal(m.listVisible, false)
		is.Equal(m.textInputVisible, false)
		is.Equal(m.yesNoInputVisible, false)
		is.Equal(len(m.messages), 0)
	})
}

func TestChecksModelUpdate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it            string
		inputModel    model
		inputMsg      tea.Msg
		expectedModel model
		expectedCmd   tea.Cmd
	}{
		{
			it:            "quits correctly",
			inputModel:    initialModel(),
			inputMsg:      tea.KeyMsg(tea.Key{Type: tea.KeyCtrlC}),
			expectedModel: initialModel(),
			expectedCmd:   tea.Quit,
		},
		{
			it:         "updates each step's correctly",
			inputModel: initialModel(),
			inputMsg:   stepMsg{Message: "test message"},
			expectedModel: (func() model {
				m := initialModel()
				m.messages = []string{"test message"}
				return m
			})(),
		},
		{
			it:         "updates context message correctly",
			inputModel: initialModel(),
			inputMsg:   contextMsg("test context message"),
			expectedModel: (func() model {
				m := initialModel()
				m.context = "test context message"
				return m
			})(),
			expectedCmd: nil,
		},
		{
			it:         "updates summary message correctly",
			inputModel: initialModel(),
			inputMsg:   summaryMsg("test summary message"),
			expectedModel: (func() model {
				m := initialModel()
				m.messages = []string{"test summary message"}
				return m
			})(),
			expectedCmd: nil,
		},
		{
			it:         "sets list choices correctly",
			inputModel: initialModel(),
			inputMsg:   listChoiceMsg{title: "test list title", items: []string{"item 1", "item 2"}, choice: make(chan int)},
			expectedModel: (func() model {
				m := initialModel()
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
			inputModel: initialModel(),
			inputMsg:   yesNoInputMsg{question: "test question"},
			expectedModel: (func() model {
				m := initialModel()
				m.yesNoInputVisible = true
				m.yesNoInput.Title = "test question"
				return m
			})(),
			expectedCmd: nil,
		},
		{
			it:         "sets up text input correctly",
			inputModel: initialModel(),
			inputMsg:   textInputMsg{question: "test question"},
			expectedModel: (func() model {
				m := initialModel()
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
			resultModel, ok := m.(model)
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
			is.Equal(resultModel.context, tc.expectedModel.context)                     // context text matches
			is.Equal(resultModel.error, tc.expectedModel.error)                         // error object matches
			is.True(stringSubsequence(resultModel.messages, tc.expectedModel.messages)) // all expected messages present
			is.True(stringSubsequence(tc.expectedModel.messages, resultModel.messages)) // only expected messages present
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
