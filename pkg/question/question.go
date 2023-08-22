package question

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type Question struct {
	Answer []string
}

func NewQuestion() *Question {
	return &Question{
		Answer: []string{},
	}
}

func (q *Question) addResult(ans string) {
	q.Answer = append(q.Answer, ans)
}

func (q *Question) AskSelectQuestion(label string, items []string) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	q.addResult(result)
}

func (q *Question) AskInputQuestion(label string) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	q.addResult(result)
}
