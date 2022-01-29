package product

import (
	"log"
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{Name: "test", Price: 2, SKU: "abc-adc-r"}

	err := p.Validate()

	if err != nil {
		log.Fatal(err)
	}
}
