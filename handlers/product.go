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

type Product struct {
	ID        int     `json:"Id"`
	Name      string  `json:"Name" validate:"required"`
	Price     float32 `json:"Price" validate:"gt=0"`
	SKU       string  `json:"sku" validate:"required,sku"`
	CreatedOn string  `json:"-"`
	UpdateOn  string  `json:"-"`
	DeleteOn  string  `json:"-"`
}

type products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *products {
	return &products{l}
}

//swagger:route GET /products products listProducts
//Returns a list of products
//Responses:
//	200: productResponse
func (p *products) GetProduct(rw http.ResponseWriter, r *http.Request) {

	lp := product.GetProducts()

	// This converts the lp object into json format and calls he Write method of rw.
	err2 := json.NewEncoder(rw).Encode(lp)

	if err2 != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

//swagger:route DELETE /products/{id} products deleteProduct
//Deletes a product with given id from database
//responses:
//200: noContent

//DeleteProduct deletes a product from the database
func (p *products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Delete request is called ")
	vars := mux.Vars(r)
	id, err2 := strconv.Atoi(vars["id"])

	if err2 != nil {
		log.Println(err2)
	}

	p.l.Println("Handle Delete request ", id)

	productErr := product.DeleteProduct(id)

	if productErr != nil {
		p.l.Println("[ERROR] deleting record", productErr)
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
