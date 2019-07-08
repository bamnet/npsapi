// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event event
// swagger:model Event
type Event struct {

	// data
	Data []*EventDataItems0 `json:"data"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// start
	Start int32 `json:"start,omitempty"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateData(formats strfmt.Registry) error {

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
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventDataItems0 event data items0
// swagger:model EventDataItems0
type EventDataItems0 struct {

	// Category for event
	Category string `json:"category,omitempty"`

	// Email address for event contact
	ContactEmailAddress string `json:"contactEmailAddress,omitempty"`

	// Name of event contact
	ContactName string `json:"contactName,omitempty"`

	// Phone number for event contact
	ContactTelephoneNumber string `json:"contactTelephoneNumber,omitempty"`

	// Date of next upcoming event
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// End date for event
	// Format: date-time
	DateEnd strfmt.DateTime `json:"dateEnd,omitempty"`

	// Start date for event
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// Array of event dates
	Dates []strfmt.DateTime `json:"dates"`

	// Event description
	Description string `json:"description,omitempty"`

	// Fee information for event
	FeeInfo string `json:"feeInfo,omitempty"`

	// Unique identifier for this event
	ID string `json:"id,omitempty"`

	// images
	Images []*EventDataItems0ImagesItems0 `json:"images"`

	// URL for more information about the event
	InfoURL string `json:"infoURL,omitempty"`

	// The event takes place all day
	IsAllDay bool `json:"isAllDay,omitempty"`

	// The event is free
	IsFree bool `json:"isFree,omitempty"`

	// The event has recurrence
	IsRecurring bool `json:"isRecurring,omitempty"`

	// The event requires registration or reservation
	IsRegResRequired bool `json:"isRegResRequired,omitempty"`

	// The latitude of the event location
	Latitude float32 `json:"latitude,omitempty"`

	// The location the event takes place
	Location string `json:"location,omitempty"`

	// The longitude of the event location
	Longitude float32 `json:"longitude,omitempty"`

	// Name of the organization associated with event
	OrganizationName string `json:"organizationName,omitempty"`

	// Name and designation of the park associated with event
	ParkFullName string `json:"parkFullName,omitempty"`

	// Name of the portal site associated with event
	PortalName string `json:"portalName,omitempty"`

	// Date the event recurrence ends
	// Format: date-time
	RecurrenceDateEnd strfmt.DateTime `json:"recurrenceDateEnd,omitempty"`

	// Recurrence rule for event
	RecurrenceRule string `json:"recurrenceRule,omitempty"`

	// Additional information on required reservation or registration for event
	RegResInfo string `json:"regResInfo,omitempty"`

	// URL for required reservation or registration for event
	RegResURL string `json:"regResURL,omitempty"`

	// Site code of the associated site for event
	SiteCode string `json:"siteCode,omitempty"`

	// Site type of the associated site for event
	SiteType string `json:"siteType,omitempty"`

	// Name of associated subject site for event
	SubjectName string `json:"subjectName,omitempty"`

	// Tags associated with event
	Tags []string `json:"tags"`

	// Additional information about times for event
	TimeDescription string `json:"timeDescription,omitempty"`

	// Time information for event
	Times []*EventDataItems0TimesItems0 `json:"times"`

	// Event title
	Title string `json:"title,omitempty"`

	// Type(s) of event
	Types []string `json:"types"`
}

// Validate validates this event data items0
func (m *EventDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecurrenceDateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventDataItems0) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventDataItems0) validateDateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.DateEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("dateEnd", "body", "date-time", m.DateEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventDataItems0) validateDateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventDataItems0) validateDates(formats strfmt.Registry) error {

	if swag.IsZero(m.Dates) { // not required
		return nil
	}

	for i := 0; i < len(m.Dates); i++ {

		if err := validate.FormatOf("dates"+"."+strconv.Itoa(i), "body", "date-time", m.Dates[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *EventDataItems0) validateImages(formats strfmt.Registry) error {

	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EventDataItems0) validateRecurrenceDateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.RecurrenceDateEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("recurrenceDateEnd", "body", "date-time", m.RecurrenceDateEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventDataItems0) validateTimes(formats strfmt.Registry) error {

	if swag.IsZero(m.Times) { // not required
		return nil
	}

	for i := 0; i < len(m.Times); i++ {
		if swag.IsZero(m.Times[i]) { // not required
			continue
		}

		if m.Times[i] != nil {
			if err := m.Times[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("times" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventDataItems0) UnmarshalBinary(b []byte) error {
	var res EventDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventDataItems0ImagesItems0 Event image
// swagger:model EventDataItems0ImagesItems0
type EventDataItems0ImagesItems0 struct {

	// Image alt text
	AltText string `json:"altText,omitempty"`

	// Image caption
	Caption string `json:"caption,omitempty"`

	// Image credit
	Credit string `json:"credit,omitempty"`

	// Image description
	Description string `json:"description,omitempty"`

	// Image title
	Title string `json:"title,omitempty"`

	// Image URL
	URL string `json:"url,omitempty"`
}

// Validate validates this event data items0 images items0
func (m *EventDataItems0ImagesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventDataItems0ImagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventDataItems0ImagesItems0) UnmarshalBinary(b []byte) error {
	var res EventDataItems0ImagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EventDataItems0TimesItems0 event data items0 times items0
// swagger:model EventDataItems0TimesItems0
type EventDataItems0TimesItems0 struct {

	// Event begins at sunrise
	SunriseStart bool `json:"sunriseStart,omitempty"`

	// Event ends at sunset
	SunsetEnd bool `json:"sunsetEnd,omitempty"`

	// Time event ends
	// Format: date-time
	TimeEnd strfmt.DateTime `json:"timeEnd,omitempty"`

	// Time event begins
	// Format: date-time
	TimeStart strfmt.DateTime `json:"timeStart,omitempty"`
}

// Validate validates this event data items0 times items0
func (m *EventDataItems0TimesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventDataItems0TimesItems0) validateTimeEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("timeEnd", "body", "date-time", m.TimeEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EventDataItems0TimesItems0) validateTimeStart(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeStart) { // not required
		return nil
	}

	if err := validate.FormatOf("timeStart", "body", "date-time", m.TimeStart.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventDataItems0TimesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventDataItems0TimesItems0) UnmarshalBinary(b []byte) error {
	var res EventDataItems0TimesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}