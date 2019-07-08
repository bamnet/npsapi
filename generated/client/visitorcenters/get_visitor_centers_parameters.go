// Code generated by go-swagger; DO NOT EDIT.

package visitorcenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVisitorCentersParams creates a new GetVisitorCentersParams object
// with the default values initialized.
func NewGetVisitorCentersParams() *GetVisitorCentersParams {
	var ()
	return &GetVisitorCentersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVisitorCentersParamsWithTimeout creates a new GetVisitorCentersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVisitorCentersParamsWithTimeout(timeout time.Duration) *GetVisitorCentersParams {
	var ()
	return &GetVisitorCentersParams{

		timeout: timeout,
	}
}

// NewGetVisitorCentersParamsWithContext creates a new GetVisitorCentersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVisitorCentersParamsWithContext(ctx context.Context) *GetVisitorCentersParams {
	var ()
	return &GetVisitorCentersParams{

		Context: ctx,
	}
}

// NewGetVisitorCentersParamsWithHTTPClient creates a new GetVisitorCentersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVisitorCentersParamsWithHTTPClient(client *http.Client) *GetVisitorCentersParams {
	var ()
	return &GetVisitorCentersParams{
		HTTPClient: client,
	}
}

/*GetVisitorCentersParams contains all the parameters to send to the API endpoint
for the get visitor centers operation typically these are written to a http.Request
*/
type GetVisitorCentersParams struct {

	/*Fields
	  A comma delimited list of resource properties to include in the JSON response in addition to the default properties. The available properties for each resource are listed in the documentation for each resource. Invalid property values will be ignored.

	*/
	Fields []string
	/*Limit
	  Number of results to return per request. Default is 50.

	*/
	Limit *int64
	/*ParkCode
	  A comma delimited list of park codes (each 4-10 characters in length).

	*/
	ParkCode []string
	/*Q
	  Term to search on

	*/
	Q *string
	/*Sort
	  A comma delimited list of resource properties to sort the results by. Each resource identifies which properties are 'sortable'. Ascending order is assumed for each property. If descending order is desired, the unary negative should prefix the property name. The sortable properties are listed in the documentation for each resource. Invalid property values will be ignored.

	*/
	Sort []string
	/*Start
	  Get the next [limit] results starting with this number.

	*/
	Start *int64
	/*StateCode
	  A comma delimited list of 2 character state codes.

	*/
	StateCode []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get visitor centers params
func (o *GetVisitorCentersParams) WithTimeout(timeout time.Duration) *GetVisitorCentersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get visitor centers params
func (o *GetVisitorCentersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get visitor centers params
func (o *GetVisitorCentersParams) WithContext(ctx context.Context) *GetVisitorCentersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get visitor centers params
func (o *GetVisitorCentersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get visitor centers params
func (o *GetVisitorCentersParams) WithHTTPClient(client *http.Client) *GetVisitorCentersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get visitor centers params
func (o *GetVisitorCentersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get visitor centers params
func (o *GetVisitorCentersParams) WithFields(fields []string) *GetVisitorCentersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get visitor centers params
func (o *GetVisitorCentersParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get visitor centers params
func (o *GetVisitorCentersParams) WithLimit(limit *int64) *GetVisitorCentersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get visitor centers params
func (o *GetVisitorCentersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithParkCode adds the parkCode to the get visitor centers params
func (o *GetVisitorCentersParams) WithParkCode(parkCode []string) *GetVisitorCentersParams {
	o.SetParkCode(parkCode)
	return o
}

// SetParkCode adds the parkCode to the get visitor centers params
func (o *GetVisitorCentersParams) SetParkCode(parkCode []string) {
	o.ParkCode = parkCode
}

// WithQ adds the q to the get visitor centers params
func (o *GetVisitorCentersParams) WithQ(q *string) *GetVisitorCentersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get visitor centers params
func (o *GetVisitorCentersParams) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the get visitor centers params
func (o *GetVisitorCentersParams) WithSort(sort []string) *GetVisitorCentersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get visitor centers params
func (o *GetVisitorCentersParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithStart adds the start to the get visitor centers params
func (o *GetVisitorCentersParams) WithStart(start *int64) *GetVisitorCentersParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get visitor centers params
func (o *GetVisitorCentersParams) SetStart(start *int64) {
	o.Start = start
}

// WithStateCode adds the stateCode to the get visitor centers params
func (o *GetVisitorCentersParams) WithStateCode(stateCode []string) *GetVisitorCentersParams {
	o.SetStateCode(stateCode)
	return o
}

// SetStateCode adds the stateCode to the get visitor centers params
func (o *GetVisitorCentersParams) SetStateCode(stateCode []string) {
	o.StateCode = stateCode
}

// WriteToRequest writes these params to a swagger request
func (o *GetVisitorCentersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "multi")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesParkCode := o.ParkCode

	joinedParkCode := swag.JoinByFormat(valuesParkCode, "multi")
	// query array param parkCode
	if err := r.SetQueryParam("parkCode", joinedParkCode...); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "multi")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	valuesStateCode := o.StateCode

	joinedStateCode := swag.JoinByFormat(valuesStateCode, "multi")
	// query array param stateCode
	if err := r.SetQueryParam("stateCode", joinedStateCode...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}