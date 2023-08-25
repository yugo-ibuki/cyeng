package question

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type Question struct {
	Corrections []Correction
}

type Correction struct {
	Question    string
	SelectedAns string
	CorrectAns  string
	IsCorrect   bool
}

func NewQuestion() *Question {
	return &Question{
		Corrections: []Correction{},
	}
}

func (q *Question) AskSelectQuestion(label string, items []string, correct string) error {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	q.Corrections = append(q.Corrections, Correction{
		Question:    label,
		SelectedAns: result,
		CorrectAns:  correct,
		IsCorrect:   result == correct,
	})

	return nil
}
