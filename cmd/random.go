/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Returns a random joke",
	Long:  `Command 'random' gives you a random joke in the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
		getRandomJoke()
	},
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

func getRandomJoke() {
	fmt.Println("Getting random joke for sami...")
}
