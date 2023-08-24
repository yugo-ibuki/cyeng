package cmd

import (
	"cyeng/pkg/question"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"os"
)

var path string

type Data struct {
	Question   string   `json:"question"`
	Selections []string `json:"selections"`
	Correct    string   `json:"correct"`
}

var rootCmd = &cobra.Command{
	Use:   "fortune-golling",
	Short: "You can get the fortune cookie message from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("error occurs...: ", err)
			os.Exit(1)
		}

		var d []Data
		if err := json.Unmarshal([]byte(data), &d); err != nil {
			fmt.Println("error occurs...: ", err)
			os.Exit(1)
		}

		q := question.NewQuestion()
		for _, v := range d {
			if err := q.AskSelectQuestion(v.Question, v.Selections, v.Correct); err != nil {
				fmt.Println("error occurs...: ", err)
				os.Exit(1)
			}
		}

		for _, c := range q.Corrections {
			fmt.Println("Question: ", c.Question)
			fmt.Println("Selected Answer: ", c.SelectedAns)
			fmt.Println("Correct Answer: ", c.CorrectAns)
			fmt.Println("You are: ", lo.Ternary(c.IsCorrect, "correct", "incorrect"))
			fmt.Println(" ----------------------------------------- ")
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "./data/fortune.json", "path to json file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
