package livelint

func (n *Livelint) askUserQuestion(question string) string {
	textResponse := make(chan string)
	n.tea.Send(textInputMsg{question: question, value: textResponse})

	input := <-textResponse

	return input
}
