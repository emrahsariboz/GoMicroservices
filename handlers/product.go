package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/emrahsariboz/microservices/product"
	"github.com/gorilla/mux"
)

type products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *products {
	return &products{l}
}

// func (p *products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	// This will get the data as JSON format.
// 	if r.Method == http.MethodGet {
// 		p.getProduct(rw, r)
// 		return
// 	} else if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		// Expect the ID in URI.
// 		// r := regexp.MustCompile(`/([0-9]+)`)
// 		// g := r.FindAllStringSubmatch(r.URL.Path, -1)
// 		p.l.Println("The request is ", r.Method)
// 		t := r.URL.Path

// 		if len(t) == 1 {
// 			p.l.Println("Invalid ID")
// 			return
// 		}
// 		id, _ := strconv.Atoi(string(t[1:]))

// 		p.l.Println("The id is ", id)
// 		p.updateProduct(id, rw, r)
// 	}

// 	// Handle UPDATE
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *products) GetProduct(rw http.ResponseWriter, r *http.Request) {

	lp := product.GetProducts()

	// This converts the lp object into json format and calls he Write method of rw.
	err2 := json.NewEncoder(rw).Encode(lp)

	if err2 != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *products) addProduct(rw http.ResponseWriter, r *http.Request) {

	np := product.NewProduct()

	err := json.NewDecoder(r.Body).Decode(np)

	if err != nil {
		fmt.Println("Something wrong")
	}

	p.l.Printf("Prod: %#v", np)

	product.AddProduct(np)
}

func (p *products) PutProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err2 := strconv.Atoi(vars["id"])

	if err2 != nil {
		p.l.Println("Something wrong")
	}
	np := product.NewProduct()

	err := json.NewDecoder(r.Body).Decode(np)

	if err != nil {
		fmt.Println("Something wrong")
	}

	product.UpdateProduct(id, *np)
}
