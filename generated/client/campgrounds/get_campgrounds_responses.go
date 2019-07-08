// Code generated by go-swagger; DO NOT EDIT.

package campgrounds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bamnet/npsapi/generated/models"
)

// GetCampgroundsReader is a Reader for the GetCampgrounds structure.
type GetCampgroundsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCampgroundsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCampgroundsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCampgroundsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCampgroundsOK creates a GetCampgroundsOK with default headers values
func NewGetCampgroundsOK() *GetCampgroundsOK {
	return &GetCampgroundsOK{}
}

/*GetCampgroundsOK handles this case with default header values.

successful operation
*/
type GetCampgroundsOK struct {
	Payload []*models.Campground
}

func (o *GetCampgroundsOK) Error() string {
	return fmt.Sprintf("[GET /campgrounds][%d] getCampgroundsOK  %+v", 200, o.Payload)
}

func (o *GetCampgroundsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCampgroundsBadRequest creates a GetCampgroundsBadRequest with default headers values
func NewGetCampgroundsBadRequest() *GetCampgroundsBadRequest {
	return &GetCampgroundsBadRequest{}
}

/*GetCampgroundsBadRequest handles this case with default header values.

Invalid status value
*/
type GetCampgroundsBadRequest struct {
}

func (o *GetCampgroundsBadRequest) Error() string {
	return fmt.Sprintf("[GET /campgrounds][%d] getCampgroundsBadRequest ", 400)
}

func (o *GetCampgroundsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
