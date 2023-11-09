package main

import "github.com/spf13/viper"

type config struct {
	Key    string
	User   string
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
