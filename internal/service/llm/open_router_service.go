package llm

import (
	"context"
	"log"

	"github.com/eduardolat/openroutergo"
	"github.com/rotisserie/eris"
)

type openRouterLLMService struct {
	client *openroutergo.Client
	model  string
}

func newOpenRouterService(apiKey, model string) LLMService {
	if apiKey == "" {
		log.Fatalf("missing OpenRouter API Key")
	}
	if model == "" {
		log.Fatalf("OpenRouter model is not specified")
	}

	client, err := openroutergo.NewClient().
		WithAPIKey(apiKey).
		WithRefererTitle("TogetherList").
		Create()

	if err != nil {
		log.Fatalf("error creating open router client: %v", err)
	}

	return &openRouterLLMService{
		client,
		model,
	}
}

func (ls *openRouterLLMService) GetResponse(ctx context.Context, prompt string) (string, error) {
	_, response, err := ls.client.
		NewChatCompletion().
		WithModel(ls.model).
		// WithSystemMessage("You are a helpful assistant expert in geography.").
		WithUserMessage(prompt).
		Execute()
	if err != nil {
		return "", eris.Wrap(err, "error retrieving response")
	}

	if len(response.Choices) == 0 {
		return "", eris.New("no response choices returned from OpenRouter API")
	}

	return response.Choices[0].Message.Content, nil
}
