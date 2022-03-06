package app

import (
	"dev-11/internal/config"
	"dev-11/internal/delivery/api/event"
	"dev-11/internal/repository/cache"
	"dev-11/internal/service"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	log.Println("init config")
	cfg := config.NewConfig()

	log.Println("init router")
	router := http.NewServeMux()

	server := &http.Server{
		Addr:           cfg.App.Host + ":" + cfg.App.Port,
		Handler:        router,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	c := cache.NewCache()
	s := service.NewService(c)
	h := event.NewHandler(s)

	log.Println("register handlers")
	h.Register(router)

	go gracefulShutdown([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, server)

	log.Println("start serve")
	return server.ListenAndServe()
}

func gracefulShutdown(signals []os.Signal, closeItems ...io.Closer) {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign
	log.Printf("Caught signal %s. Shutting down...", sig)
	for _, closer := range closeItems {
		err := closer.Close()
		if err != nil {
			fmt.Printf("failed to close %v: %v", closer, err)
		}
	}
}
