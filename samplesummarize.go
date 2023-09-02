package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	gogpt "github.com/tapp-ai/go-openai"
	"go.uber.org/zap"
)

type Summarize struct {
	Content string `json:"content"`
}

//	 Summarize is a sample POST request endpoint with the following JSON body
//	 {
//			"content": "content to be summarized"
//	 }
func (a *App) Summarize(c *fiber.Ctx) error {
	// parse the object into the summarize struct
	content := Summarize{}
	err := c.BodyParser(&content)
	if err != nil {
		a.Log.Error("error parsing content into the struct")
		return c.JSON(ErrorResponse("error parsing content into the struct"))
	}

	// write a sample summarize prompt
	prompt := fmt.Sprintf(`Summarize the following text.
	Input: %s
	Output:`, content.Content)

	// call the chat completion
	resp, _ := a.GptClient.CreateChatCompletion(
		c.Context(), gogpt.ChatCompletionRequest{
			Model: gogpt.GPT3Dot5Turbo,
			Messages: []gogpt.ChatCompletionMessage{
				{
					Role:    gogpt.ChatMessageRoleSystem,
					Content: "You're a summarizing chat bot that's only role is to summarize text",
				},
				{
					Role:    gogpt.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 2000,
		},
	)
	if err != nil {
		a.Log.Error("Error in SEO summarizer", zap.Error(err))
		return c.JSON(ErrorResponse("Error in SEO summarizer"))
	}

	// this line is for debugging the response
	choices, err := json.Marshal(resp)
	if err != nil {
		a.Log.Error("Failed to marshal response", zap.Error(err))
		return c.JSON(ErrorResponse("Failed to marshal response"))
	}
	a.Log.Info(string(choices))

	// return the final response
	return c.JSON(SuccessResponse(resp.Choices[0].Message.Content))
}
