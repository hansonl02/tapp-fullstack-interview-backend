package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	gogpt "github.com/tapp-ai/go-openai"
	"go.uber.org/zap"
)

func main() {
	var log *zap.Logger
	log, _ = zap.NewDevelopment()

	OpenAIApiKey := "TODO: Insert OpenAI Key"
	// create gpt client
	gptClient := gogpt.NewClient(OpenAIApiKey)

	// create http client
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// creates an App object.
	a := &App{
		Log:        log,
		GptClient:  gptClient,
		HttpClient: httpClient,
	}

	// starts a new fiber web server instance.
	r := fiber.New()

	// add CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// use this to check your connection and make sure your localhost is running.
	r.Get("/", a.ServeRoot)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	// sample Endpoint Summarize
	// this is a POST request endpoint at the location http://127.0.0.1:8085/api/v1/summarize
	// refer to summarize.go for a walkthrough of the post request body.
	v1.Post("/summarize", a.Summarize)

	// TODO: Create an endpoint that takes in a business name as an input, and generates
	// a list of possible domains with GPT. Then check which ones are available and returns a list
	// to the front end.
	// Note: You may need to parse out the response of GPT into a usable array, as GPT returns a
	// response in string format
	// Use this endpoint "http://domains.usestyle.ai/api/v1/availability?domain={DOMAIN}" to check if a domain is available
	v1.Post("get-domains", a.GetDomains)

	log.Info("sampleapi started")
	err := r.Listen(":8085")
	if err != nil {
		panic(err)
	}
}
