// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteProductReader is a Reader for the DeleteProduct structure.
type DeleteProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProductOK creates a DeleteProductOK with default headers values
func NewDeleteProductOK() *DeleteProductOK {
	return &DeleteProductOK{}
}

/* DeleteProductOK describes a response with status code 200, with default header values.

DeleteProductOK delete product o k
*/
type DeleteProductOK struct {
}

func (o *DeleteProductOK) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductOK ", 200)
}

func (o *DeleteProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
