// Package classification of Product type
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import "github.com/emrahsariboz/microservices/product"

// swagger:response noContent
type productNoContent struct {
}

//swagger:parameters deleteProduct
type ProductIdParameter struct {
	// The id of the product to delete from
	// in: path
	// required: true
	ID int `json:"id"`
}

//List of products returns in the response
// swagger:response productResponse
type productResponse struct {
	//All products in the system
	//in: body
	Body []product.Product
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
