package handlers

import (
	"context"
	"encoding/json"
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

func (p *products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(product.Product)

	p.l.Println("PRODCT PASSED", prod)
	product.AddProduct(&prod)
}

func (p *products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err2 := strconv.Atoi(vars["id"])

	if err2 != nil {
		log.Println(err2)
	}

	p.l.Println("Handle PUT request ", id)

	prod := r.Context().Value(KeyProduct{}).(product.Product)

	product.UpdateProduct(id, prod)

}

type KeyProduct struct{}

// Creates Middleware - A step before or after an http request.
func (p *products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		np := product.NewProduct()

		err := json.NewDecoder(r.Body).Decode(np)

		if err != nil {
			http.Error(rw, "Cannot unmarshall the json", http.StatusBadRequest)
		}

		err = np.Validate()

		if err != nil {
			p.l.Println("[ERROR] validating product")
			http.Error(rw, "Error validating product", http.StatusBadGateway)
			return
		}
		// Every http request has context object embedded in it
		// which is used to store information during the lifetime of request. (Used for passing request scope values)
		// A good use case is pass info between middleware and handlers.
		// Another use case is to check the request creater is authenticated, etc.
		ctx := context.WithValue(r.Context(), KeyProduct{}, *np)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
