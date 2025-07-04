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

// Close releases resources used by the Google LLM service
func (g *googleLLMService) Close() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}

func newGoogleLLMService(apiKey, model string) *googleLLMService {
	if apiKey == "" {
		log.Fatalf("missing Google AI API Key")
	}
	if model == "" {
		log.Fatalf("Google AI model is not specified")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal("error creating Google GenAI client: ", err)
	}

	modelVar, err := client.Models.Get(ctx, model, nil)
	if err != nil {
		log.Fatalf("error validating Google AI model: %v", err)
	}
	if modelVar == nil {
		log.Fatalf("Google AI model: %s is not found/available", model)
	}

	return &googleLLMService{
		client,
		model,
	}
}

func (gs *googleLLMService) Close() error {
	if gs.client != nil {
		return gs.client.Close()
	}
	return nil
}

func (gs *googleLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	response, err := gs.client.Models.GenerateContent(ctx, gs.model, genai.Text(prompt), nil)
	if err != nil {
		return "", eris.Wrap(err, "error prompting model")
	}

	return response.Text(), nil

// Close releases resources used by the Google LLM service
func (g *googleLLMService) Close() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}
}

// Close releases resources used by the Google LLM service
func (g *googleLLMService) Close() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}

// Close properly cleans up the GenAI client resources
func (gs *googleLLMService) Close() error {
	if gs.client != nil {
		return gs.client.Close()
	}
	return nil
}

func (g *googleLLMService) Close() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}

// Close properly cleans up the GenAI client connection
func (gs *googleLLMService) Close() error {
	if gs.client != nil {
		return gs.client.Close()
	}
	return nil
}

// Close properly cleans up the Google GenAI client resources
func (g *googleLLMService) Close() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}
