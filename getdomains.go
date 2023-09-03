package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	gogpt "github.com/tapp-ai/go-openai"
	"go.uber.org/zap"
)

type GetDomains struct {
	Name string `json:"name"`
}

type AvailabilityData struct {
}

//	 GetDomains gets a list of available domains from a given business name
//	 {
//			"name": "name of business"
//	 }
func (a *App) GetDomains(c *fiber.Ctx) error {
	name := GetDomains{}
	err := c.BodyParser(&name)
	if err != nil {
		a.Log.Error("error parsing name into the struct")
		return c.JSON(ErrorResponse("error parsing name into the struct"))
	}

	prompt := fmt.Sprintf(`Generate a comma-separated list of twenty potential website domains for my business named %s, in CSV format with the data and nothing else.`, name.Name)
	resp, err := a.GptClient.CreateChatCompletion(
		c.Context(),
		gogpt.ChatCompletionRequest{
			Model: gogpt.GPT3Dot5Turbo,
			Messages: []gogpt.ChatCompletionMessage{
				{
					Role:    gogpt.ChatMessageRoleUser,
					Content: prompt,
				},
			},
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

	domains := resp.Choices[0].Message.Content
	domains = strings.ReplaceAll(domains, "\n", "")

	requestURL := fmt.Sprintf(`https://domains.usestyle.ai/api/v1/availability?domains=%s`, domains)
	res, err := a.HttpClient.Get(requestURL)
	if err != nil {
		a.Log.Error("Failed during availability API call", zap.Error(err))
		return c.JSON(ErrorResponse("Failed during availability API call"))
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	defer res.Body.Close()
	if err != nil {
		a.Log.Error("Failed decoding availability API response", zap.Error(err))
		return c.JSON(ErrorResponse("Failed decoding availability API response"))
	}

	var availableDomains []string
	for k, v := range data {
		if k == "data" {
			for domain, available := range v.(map[string]interface{}) {
				if available.(bool) {
					availableDomains = append(availableDomains, domain)
				}
			}
			break
		}
	}

	return c.JSON(SuccessResponse(availableDomains))
}
