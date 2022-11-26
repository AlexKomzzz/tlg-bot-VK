package config

import "github.com/spf13/viper"

type Messages struct {
	Responses
	Errors
	dataVKapi
}

type dataVKapi struct {
	AppID   int    `mapstructure:"app_id"`
	Scope   string `mapstructure:"scope"`
	Version string `mapstructure:"version"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	RedirectURL       string `mapstructure:"redirectURL"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	UnknownCommand    string `mapstructure:"unknown_command"`
	LinkSaved         string `mapstructure:"link_saved"`
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Config struct {
	TelegramToken  string
	ClientSecretVK string

	BotURL     string `mapstructure:"bot_url"`
	BoltDBFile string `mapstructure:"db_file"`

	Messages Messages
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.vk_api", &cfg.Messages.dataVKapi); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("token_tlg"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token_tlg")

	if err := viper.BindEnv("client_secret"); err != nil {
		return err
	}
	cfg.ClientSecretVK = viper.GetString("client_secret")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
