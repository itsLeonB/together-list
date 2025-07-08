package config

type whatsappConfigLoader struct {
	Env         string `split_words:"true" required:"true"`
	DatabaseUrl string `split_words:"true" required:"true"`
}

func (cl *whatsappConfigLoader) Load() *Config {
	return &Config{
		Env:         cl.Env,
		DatabaseUrl: cl.DatabaseUrl,
	}
}
