package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	var client *openai.Client
	if err := godotenv.Load(); err == nil {
		client = openai.NewClient(os.Getenv("GPTKEY"))
	} else {
		client = openai.NewClient("api-key")
	}
	for {
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: GetInput(),
					},
				},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Choices[0].Message.Content)
	}

}

func GetInput() (text string) {
	fmt.Print("-> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return
}
