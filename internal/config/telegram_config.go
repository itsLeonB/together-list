package config

type telegramConfigLoader struct {
	Env              string `split_words:"true" required:"true"`
	DatabaseUrl      string `split_words:"true" required:"true"`
	TelegramBotToken string `split_words:"true" required:"true"`
}

func (cl *telegramConfigLoader) Load() *Config {
	return &Config{
		Env:              cl.Env,
		DatabaseUrl:      cl.DatabaseUrl,
		TelegramBotToken: cl.TelegramBotToken,
	}
}
