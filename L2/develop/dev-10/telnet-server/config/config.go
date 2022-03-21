package config

import (
	"flag"
	"sync"
)

type Config struct {
	Host string
	Port string
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.StringVar(&instance.Port, "p", "8080", "port for telnet server")
		flag.StringVar(&instance.Host, "h", "localhost", "hostname")
		flag.Parse()
	})
	return instance
}
