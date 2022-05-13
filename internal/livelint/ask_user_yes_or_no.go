package livelint

func (n *Livelint) askUserYesOrNo(msg string) bool {
	yesNoResponse := make(chan int)
	n.tea.Send(yesNoInputMsg{question: msg, value: yesNoResponse})
	input := <-yesNoResponse

	return input == 0
}
