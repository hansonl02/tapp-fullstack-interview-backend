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

const AvailabilityURL = "https://domains.usestyle.ai/api/v1/availability"

//	 GetDomains gets a list of available domains from a given business name
//	 {
//			"name": "name of business"
//	 }
func (a *App) GetDomains(c *fiber.Ctx) error {
	name := GetDomains{}
	err := c.BodyParser(&name)
	if err != nil {
		a.Log.Error("Error parsing name into the struct")
		return c.JSON(ErrorResponse("Error parsing name into the struct"))
	}

	prompt := fmt.Sprintf(
		`Generate a comma-separated list of twenty potential website domains for my business named %s ending in .com, in CSV format with the data and nothing else.`,
		name.Name,
	)
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
		a.Log.Error("Error during GPT call", zap.Error(err))
		return c.JSON(ErrorResponse("Error during GPT call"))
	}
	domains := strings.ReplaceAll(resp.Choices[0].Message.Content, "\n", "")

	availabilityRequestURL := fmt.Sprintf(AvailabilityURL+`?domains=%s`, domains)
	res, err := a.HttpClient.Get(availabilityRequestURL)
	if err != nil {
		a.Log.Error("Error during availability API call", zap.Error(err))
		return c.JSON(ErrorResponse("Error during availability API call"))
	}

	var availabilityResponse map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&availabilityResponse)
	defer res.Body.Close()
	if err != nil {
		a.Log.Error("Failed decoding availability API response", zap.Error(err))
		return c.JSON(ErrorResponse("Failed decoding availability API response"))
	}

	var availableDomains []string
	if data, ok := availabilityResponse["data"]; ok {
		for domain, available := range data.(map[string]interface{}) {
			if available.(bool) {
				availableDomains = append(availableDomains, domain)
			}
		}
		if len(availableDomains) == 0 {
			a.Log.Error("Found no available domains", zap.Error(err))
			return c.JSON(ErrorResponse("Found no available domains"))
		}
		return c.JSON(SuccessResponse(availableDomains))
	} else {
		a.Log.Error("Failed decoding availability API response", zap.Error(err))
		return c.JSON(ErrorResponse("Failed decoding availability API response"))
	}
}
