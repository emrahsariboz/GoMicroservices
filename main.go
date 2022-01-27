package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/emrahsariboz/microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	p := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/", p)

	s := &http.Server{
		Addr:    ":9090",
		Handler: sm,

		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			fmt.Println("**********************************")
			l.Fatal(err)
		}
	}()

	// Waits until the client is done and server wont accept anything further.

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
