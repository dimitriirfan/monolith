package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartRESTServer() {
	cfg := LoadServerConfig()
	fmt.Println(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!\n"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server running at port %d\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	<-sigChan
	log.Printf("server stopped gracefully\n")

}
