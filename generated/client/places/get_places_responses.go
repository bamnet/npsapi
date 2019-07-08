// Code generated by go-swagger; DO NOT EDIT.

package places

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bamnet/npsapi/generated/models"
)

// GetPlacesReader is a Reader for the GetPlaces structure.
type GetPlacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPlacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlacesOK creates a GetPlacesOK with default headers values
func NewGetPlacesOK() *GetPlacesOK {
	return &GetPlacesOK{}
}

/*GetPlacesOK handles this case with default header values.

successful operation
*/
type GetPlacesOK struct {
	Payload []*models.Place
}

func (o *GetPlacesOK) Error() string {
	return fmt.Sprintf("[GET /places][%d] getPlacesOK  %+v", 200, o.Payload)
}

func (o *GetPlacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlacesBadRequest creates a GetPlacesBadRequest with default headers values
func NewGetPlacesBadRequest() *GetPlacesBadRequest {
	return &GetPlacesBadRequest{}
}

/*GetPlacesBadRequest handles this case with default header values.

Invalid status value
*/
type GetPlacesBadRequest struct {
}

func (o *GetPlacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /places][%d] getPlacesBadRequest ", 400)
}

func (o *GetPlacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
