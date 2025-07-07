package config

type whatsappConfigLoader struct {
	Env              string `split_words:"true" required:"true"`
	DatabaseUrl      string `split_words:"true" required:"true"`
	MessageKeyword   string `split_words:"true" required:"true"`
	Timezone         string
	AttachWorker     bool   `split_words:"true"`
	NotionApiKey     string `split_words:"true" required:"true"`
	NotionDatabaseId string `split_words:"true" required:"true"`
}

func (cl *whatsappConfigLoader) Load() *Config {
	return &Config{
		Env:              cl.Env,
		DatabaseUrl:      cl.DatabaseUrl,
		MessageKeyword:   cl.MessageKeyword,
		Timezone:         cl.Timezone,
		AttachWorker:     cl.AttachWorker,
		NotionApiKey:     cl.NotionApiKey,
		NotionDatabaseId: cl.NotionDatabaseId,
	}
}
