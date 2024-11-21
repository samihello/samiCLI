/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Returns a random joke",
	Long:  `Command 'random' gives you a random joke in the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		getRandomJoke()
		fmt.Println()
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
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)

	var joke Joke
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseURL string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseURL,
		nil,
	)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "samiCLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
	}

	return responseBytes
}
