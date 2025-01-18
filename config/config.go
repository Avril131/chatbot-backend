package config

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	GoogleID string   `mapstructure:"google_client_id" json:"google_client_id" yaml:"google_client_id"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	OpenAI   OpenAI   `mapstructure:"openai" json:"openai" yaml:"openai"`
}
