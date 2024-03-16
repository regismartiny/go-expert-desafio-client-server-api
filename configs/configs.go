package configs

import "github.com/spf13/viper"

var cfg *conf

type conf struct {
	QuotesAPI      string `mapstructure:"QUOTES_API"`
	QuotesAPIToken string `mapstructure:"QUOTES_API_TOKEN"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return cfg, nil
}
