package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetResponse(client gpt3.Client, ctx context.Context, ques string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			ques,
		},
		MaxTokens:   gpt3.IntPtr(200),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Print(err)
		os.Exit(13)
	}

}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) {
	return 0, nil
}

func main() {
	log.SetOutput(new(NullWriter))
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")

	if apiKey == "" {
		log.Fatal("ERROR WHILE LOADING API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Say Something ('quit' to end): ")
				if !scanner.Scan() {
					break
				}
				ques := scanner.Text()
				switch ques {
				case "quit":
					quit = true
				default:
					GetResponse(client, ctx, ques)
				}
			}
		},
	}
	rootCmd.Execute()
}
