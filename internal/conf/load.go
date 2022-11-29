package conf

import (
	"github.com/spf13/viper"
)

func Load(file string, cfg interface{}) error {

	viperInstance := viper.New()
	viperInstance.SetConfigType("toml")
	viperInstance.SetConfigFile(file)

	if err := viperInstance.ReadInConfig(); err != nil {
		return err
	}

	return viperInstance.Unmarshal(cfg)

}
