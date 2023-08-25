package cmd

import (
	"cyeng/pkg/question"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
)

var path string
var quizCount int

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
			fmt.Println("no valid path...", err)
			os.Exit(1)
		}

		var d []Data
		if err := json.Unmarshal([]byte(data), &d); err != nil {
			fmt.Println("error occurs...: ", err)
			os.Exit(1)
		}

		// check if quiz count is greater than data length
		if quizCount != 0 && len(d) < quizCount {
			fmt.Println("quiz count is greater than data length...")
			os.Exit(1)
		}

		// shuffle the data
		rand.Shuffle(len(d), func(i, j int) {
			d[i], d[j] = d[j], d[i]
		})

		q := question.NewQuestion()
		// default quiz number is 10 or if less than 10, then the length of data
		quizCount := lo.Ternary(len(d) > 10, 10, len(d))
		for i := 0; i < quizCount; i++ {
			v := d[i]
			if err := q.AskSelectQuestion(v.Question, v.Selections, v.Correct); err != nil {
				fmt.Println("error occurs...: ", err)
				os.Exit(1)
			}
		}

		for _, c := range q.Corrections {
			printQuiz(c)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "./data.json", "path to json file")
	rootCmd.Flags().IntVar(&quizCount, "quizCount", 0, "quiz count")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}

func printQuiz(c question.Correction) {
	fmt.Println("Question: ", c.Question)
	fmt.Println("Selected Answer: ", c.SelectedAns)
	fmt.Println("Correct Answer: ", c.CorrectAns)
	fmt.Println("You are: ", lo.Ternary(c.IsCorrect, "correct", "incorrect"))
	fmt.Println(" ----------------------------------------- ")
}
