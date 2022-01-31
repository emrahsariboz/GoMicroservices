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
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	p := handlers.NewProducts(l)

	sm := mux.NewRouter()

	// sm.Handle("/product", p).Method("GET")

	// Get Handler
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", p.GetProduct)

	// Put Router
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", p.UpdateProduct)

	// Delete Router
	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/{id:[0-9]+}", p.DeleteProduct)

	// Post Router
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", p.AddProduct)
	postRouter.Use(p.MiddlewareProductValidation)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	//docs handler
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	l.Println("Starting server at")
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
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
