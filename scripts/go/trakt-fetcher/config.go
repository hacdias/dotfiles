package main

import "github.com/spf13/viper"

type config struct {
	Client string
	Secret string
	Token  string
	Output string
}

func parseConfig() (*config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
