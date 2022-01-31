package product

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator"
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

var ErrProductNotFound = fmt.Errorf("Product not found")

type Products []*Product

func validateSKU(fl validator.FieldLevel) bool {

	// Data format for SKU: abc-abcd-dfsdf

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true

}

func (p *Product) Validate() error {

	validate := validator.New()

	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)

}

func UpdateProduct(id int, p Product) {
	for i := 0; i < len(ProductList); i++ {
		if ProductList[i].ID == id {
			ProductList[i] = &p
		}
	}
}

// ToDo: complete delete request.
func DeleteProduct(id int) error {
	p := FindProduct(id)

	if p < 0 {
		return ErrProductNotFound
	}

	ProductList = append(ProductList[:p], ProductList[p+1:]...)
	return nil
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

func FindProduct(index int) int {
	for i, p := range ProductList {
		if p.ID == index {
			return i
		}
	}

	return -1
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
