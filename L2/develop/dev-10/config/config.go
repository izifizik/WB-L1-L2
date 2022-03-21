package config

import (
	"flag"
	"sync"
)

type Config struct {
	Timeout string
	Host    string
	Port    string
}

var once sync.Once
var instance Config

func NewConfig() Config {
	once.Do(func() {
		flag.StringVar(&instance.Timeout, "timeout", "10s", "time limit")
		flag.StringVar(&instance.Port, "port", "8080", "port for telnet server")
		flag.StringVar(&instance.Host, "host", "localhost", "hostname")
		flag.Parse()
	})
	return instance
}
