// Code generated by go-swagger; DO NOT EDIT.

package places

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new places API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for places API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPlaces places are shared content assets that are tagged so they can appear in a variety of places on n p s gov data includes a title image short description of the content and link to more information about the asset
*/
func (a *Client) GetPlaces(params *GetPlacesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPlacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlaces",
		Method:             "GET",
		PathPattern:        "/places",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPlacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}