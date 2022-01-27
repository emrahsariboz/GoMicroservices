package product

import (
	"fmt"
	"time"
)

type Product struct {
	ID        int     `json:"Id"`
	Name      string  `json:"Name"`
	Price     float32 `json:"Price"`
	CreatedOn string  `json:"-"`
	UpdateOn  string  `json:"-"`
	DeleteOn  string  `json:"-"`
}

type Products []*Product

func UpdateProduct(id int, p Product) {
	for i := 0; i < len(ProductList); i++ {
		if ProductList[i].ID == id {
			ProductList[i] = &p
		}
	}
}

func GetProducts() Products {
	fmt.Println(ProductList)
	return ProductList
}

func AddProduct(p *Product) {
	p.ID = GetNextId()
	fmt.Println(*p)
	ProductList = append(ProductList, p)
}

func GetNextId() int {
	curr := ProductList[len(ProductList)-1]
	return curr.ID + 1
}

func NewProduct() *Product {
	return &Product{}
}

var ProductList = []*Product{
	&Product{
		ID:        1,
		Name:      "Computer",
		Price:     244.00,
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Calculator",
		Price:     32.00,
		CreatedOn: time.Now().UTC().String(),
		UpdateOn:  time.Now().UTC().String(),
	},
}
