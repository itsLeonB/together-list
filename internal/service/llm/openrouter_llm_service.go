
// Close releases resources used by the OpenRouter LLM service (no-op)
func (o *openRouterLLMService) Close() error {
	return nil
}

func (o *openRouterLLMService) Close() error {
	// HTTP client doesn't need explicit cleanup
	return nil
}

// Close implements the Service interface (no resources to cleanup for HTTP client)
func (o *openRouterLLMService) Close() error {
	return nil
}