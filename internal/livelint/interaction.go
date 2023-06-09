package livelint

import (
	tea "github.com/charmbracelet/bubbletea"
)

//go:generate moq -out mock_ui_test.go -pkg livelint_test . UserInteraction

type UserInteraction interface {
	DisplayContext(contextMsg string)
	DisplayCheckStart(checkMsg string)
	DisplayCheckResult(checkResult CheckResult)
	DisplayCheckCompletion(completionMsg string, kind SummaryType)
	DisplayError(errorMsg error)
	AskYesNo(question string) bool
	SetVerbose()
	StartSpinner()
	StopSpinner()
}

func NewBubbleTeaUI(program *tea.Program) *BubbleteaUI {
	return &BubbleteaUI{program, statusMsg{}}
}

type BubbleteaUI struct {
	*tea.Program
	statusMsg
}

var _ UserInteraction = &BubbleteaUI{}

func (ui *BubbleteaUI) AskYesNo(question string) bool {
	yesNoResponse := make(chan int)
	ui.Send(yesNoInputMsg{question: question, value: yesNoResponse})
	input := <-yesNoResponse

	return input == 0
}

func (ui *BubbleteaUI) DisplayContext(contextMsg string) {
	ui.statusMsg = initializeStatus(contextMsg)
	ui.Send(ui.statusMsg)
}

func (ui *BubbleteaUI) DisplayCheckStart(checkMsg string) {
	ui.statusMsg.StartCheck(checkMsg)
	ui.Send(ui.statusMsg)
}

func (ui *BubbleteaUI) DisplayCheckResult(checkResult CheckResult) {
	ui.statusMsg.AddCheckResult(checkResult)
	ui.Send(ui.statusMsg)
}

func (ui *BubbleteaUI) DisplayCheckCompletion(completionMsg string, kind SummaryType) {
	ui.statusMsg.CompleteCheck(summaryMsg{text: completionMsg, kind: kind})
	ui.Send(ui.statusMsg)
}

func (ui *BubbleteaUI) DisplayError(errorMsg error) {
	ui.Send(errMsg{err: errorMsg})
}

func (ui *BubbleteaUI) SetVerbose() {
	ui.Send(verboseMsg{verbose: true})
}

func (ui *BubbleteaUI) StartSpinner() {
	ui.Send(showSpinnerMsg{showing: true})
}

func (ui *BubbleteaUI) StopSpinner() {
	ui.Send(showSpinnerMsg{showing: false})
}
