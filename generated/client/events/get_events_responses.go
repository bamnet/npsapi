// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bamnet/npsapi/generated/models"
)

// GetEventsReader is a Reader for the GetEvents structure.
type GetEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsOK creates a GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

/*GetEventsOK handles this case with default header values.

successful operation
*/
type GetEventsOK struct {
	Payload []*models.Event
}

func (o *GetEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsOK  %+v", 200, o.Payload)
}

func (o *GetEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsBadRequest creates a GetEventsBadRequest with default headers values
func NewGetEventsBadRequest() *GetEventsBadRequest {
	return &GetEventsBadRequest{}
}

/*GetEventsBadRequest handles this case with default header values.

Invalid status value
*/
type GetEventsBadRequest struct {
}

func (o *GetEventsBadRequest) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsBadRequest ", 400)
}

func (o *GetEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}