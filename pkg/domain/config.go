package domain

import "github.com/spf13/viper"

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
