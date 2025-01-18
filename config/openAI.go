package config

type OpenAI struct {
	SecretKey string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	MaxTokens  int    `mapstructure:"max_tokens" json:"max_tokens" yaml:"max_token"`
	Model     string `mapstructure:"model" json:"model" yaml:"model"`
}
