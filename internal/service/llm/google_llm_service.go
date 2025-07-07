package llm

import (
	"context"

	"github.com/itsLeonB/together-list/internal/logging"
	"github.com/rotisserie/eris"
	"google.golang.org/genai"
)

type googleLLMService struct {
	client *genai.Client
	model  string
}

func newGoogleLLMService(apiKey, model string) *googleLLMService {
	if apiKey == "" {
		logging.Fatal("missing Google AI API Key")
	}
	if model == "" {
		logging.Fatal("Google AI model is not specified")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		logging.Fatalf("error creating Google GenAI client: %v", err)
	}

	modelVar, err := client.Models.Get(ctx, model, nil)
	if err != nil {
		logging.Fatalf("error validating Google AI model: %v", err)
	}
	if modelVar == nil {
		logging.Fatalf("Google AI model: %s is not found/available", model)
	}

	return &googleLLMService{
		client,
		model,
	}
}

func (gs *googleLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	response, err := gs.client.Models.GenerateContent(ctx, gs.model, genai.Text(prompt), nil)
	if err != nil {
		return "", eris.Wrap(err, "error prompting model")
	}

	return response.Text(), nil
}
