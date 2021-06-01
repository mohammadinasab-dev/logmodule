package configuration

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

//uppercase or lowercase?
type config struct {
	Environment    string `json:"environment"`
	ElasticAddress string `json:"elasticsearch"`
	Logstash       string `json:"logstash"`
}

//error return type or not?
func GetEnvironment(conf *config) string {
	return conf.Environment
}

func LoadSetup(path string) (*config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("setup")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return &config{}, err
	}
	conf := &config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		//log or fmt?
		fmt.Printf("unable to decode into config struct, %v", err)
		return &config{}, errors.New("run.environment variable not set or Not defined")
	}
	return conf, nil
}
