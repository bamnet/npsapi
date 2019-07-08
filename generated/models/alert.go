// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Alert alert
// swagger:model Alert
type Alert struct {

	// data
	Data []*AlertDataItems0 `json:"data"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// start
	Start int32 `json:"start,omitempty"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this alert
func (m *Alert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Alert) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Alert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Alert) UnmarshalBinary(b []byte) error {
	var res Alert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertDataItems0 alert data items0
// swagger:model AlertDataItems0
type AlertDataItems0 struct {

	// Alert type: danger, caution, information, or park closure
	// Enum: [Danger Caution Information Park Closure]
	Category string `json:"category,omitempty"`

	// Alert description
	Description string `json:"description,omitempty"`

	// Unique identifier for an alert record
	ID string `json:"id,omitempty"`

	// A variable width character code that uniquely identifies a specific park
	ParkCode string `json:"parkCode,omitempty"`

	// Alert title
	Title string `json:"title,omitempty"`

	// Link to more information about the alert, if available
	URL string `json:"url,omitempty"`
}

// Validate validates this alert data items0
func (m *AlertDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var alertDataItems0TypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Danger","Caution","Information","Park Closure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertDataItems0TypeCategoryPropEnum = append(alertDataItems0TypeCategoryPropEnum, v)
	}
}

const (

	// AlertDataItems0CategoryDanger captures enum value "Danger"
	AlertDataItems0CategoryDanger string = "Danger"

	// AlertDataItems0CategoryCaution captures enum value "Caution"
	AlertDataItems0CategoryCaution string = "Caution"

	// AlertDataItems0CategoryInformation captures enum value "Information"
	AlertDataItems0CategoryInformation string = "Information"

	// AlertDataItems0CategoryParkClosure captures enum value "Park Closure"
	AlertDataItems0CategoryParkClosure string = "Park Closure"
)

// prop value enum
func (m *AlertDataItems0) validateCategoryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, alertDataItems0TypeCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AlertDataItems0) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertDataItems0) UnmarshalBinary(b []byte) error {
	var res AlertDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}