package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fabriciolfj/rules-elegibility/controller"
)

func appHttp() {
	pc, err := InitController()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/eligibility", controller.RecoveryMiddleware(pc.HandleCreate))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("fail star server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Printf("Server exiting")
}

func listenerProcessRules() {
	app, err := InitListenerProcessWire()

	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	errChan := make(chan error, 2)

	go func() {
		defer func() {
			if err := app.Start(); err != nil {
				errChan <- err
			}
		}()

	}()

	select {
	case err := <-errChan:
		log.Printf("Erro: %v", err)
	case sig := <-sigChan:
		log.Printf("Sinal recebido: %v", sig)
	}

	if err := app.Close(); err != nil {
		log.Printf("Erro ao fechar listener1: %v", err)
	}

}

func main() {
	appHttp()
	listenerProcessRules()

}
