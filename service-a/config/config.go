package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		println(err)
		return
	}
	log.Println("Config file loaded successfully")
	log.Println("Mongo URI:", viper.GetString("MONGO.URI"))
	log.Println("Mongo Database:", viper.GetString("MONGO.DATABASE"))
}
