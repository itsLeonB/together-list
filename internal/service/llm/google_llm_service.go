package llm

import (
	"context"
	"log"

	"github.com/rotisserie/eris"
	"google.golang.org/genai"
)

type googleLLMService struct {
	client *genai.Client
	model  string
}

func newGoogleLLMService(apiKey, model string) *googleLLMService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal("error creating Google GenAI client: ", err)
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
