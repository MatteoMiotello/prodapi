package bootstrap

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic("init config failed " + err.Error())
	}
}
