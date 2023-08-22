package question

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type Question struct {
	Answer []string
}

func NewQuestion() *Question {
	return &Question{
		Answer: []string{},
	}
}

func (q *Question) AskSelectQuestion(label string, items []string) error {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	q.Answer = append(q.Answer, result)
	return nil
}

func (q *Question) AskInputQuestion(label string) error {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	q.Answer = append(q.Answer, result)
	return nil
}
