package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	// Initialize LLM
	// If you want to use an alternae LLM OPenAI compatible, just set the follwowing environment variables:
	// export OPENAI_API_BASE=https://api.openai.com/v1
	// export OPENAI_MODEL=gpt-3.5-turbo
	// export OPENAI_API_KEY=your_api_key
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create conversation memory
	chatMemory := memory.NewConversationBuffer()

	fmt.Println("Chat Application Started! Type 'quit' to exit.")

	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "quit" {
			break
		}

		// Get response from LLM
		response, err := llm.GenerateContent(ctx, []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeHuman, input),
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		aiResponse := response.Choices[0].Content
		fmt.Printf("AI: %s\n\n", aiResponse)

		// Store conversation in memory
		chatMemory.ChatHistory.AddUserMessage(ctx, input)
		chatMemory.ChatHistory.AddAIMessage(ctx, aiResponse)
	}
}
