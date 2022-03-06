package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	App struct {
		Port string `env-default:"8080"`
		Host string `env-default:"localhost"`
	}
}

var once sync.Once
var instance Config

//NewConfig - create config from env
func NewConfig() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println(err.Error())
		}

		instance.App.Port = os.Getenv("PORT")
		instance.App.Host = os.Getenv("HOST")
	})
	return instance
}
