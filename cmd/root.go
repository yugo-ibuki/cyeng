package cmd

import (
	"fmt"
	"fortune-golling/pkg/question"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fortune-golling",
	Short: "You can get the fortune cookie message from the command line.",
	Run: func(cmd *cobra.Command, args []string) {
		q := question.NewQuestion()
		q.AskInputQuestion("What is your name?")
		q.AskSelectQuestion("What is your favorite food?", []string{"Ramen", "Carry", "Sushi"})

		fmt.Println("answer: ", q.Answer)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print("error occurs...: ", err)
		os.Exit(1)
	}
}
