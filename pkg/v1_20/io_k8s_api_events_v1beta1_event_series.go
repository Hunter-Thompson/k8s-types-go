// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPIEventsV1beta1EventSeries EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
//
// swagger:model io.k8s.api.events.v1beta1.EventSeries
type IoK8sAPIEventsV1beta1EventSeries struct {

	// count is the number of occurrences in this series up to the last heartbeat time.
	// Required: true
	Count *int32 `json:"count" json,yaml:"count"`

	// lastObservedTime is the time when last Event from the series was seen before last heartbeat.
	// Required: true
	// Format: date-time
	LastObservedTime *IoK8sApimachineryPkgApisMetaV1MicroTime `json:"lastObservedTime" json,yaml:"lastObservedTime"`
}

// Validate validates this io k8s api events v1beta1 event series
func (m *IoK8sAPIEventsV1beta1EventSeries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastObservedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIEventsV1beta1EventSeries) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIEventsV1beta1EventSeries) validateLastObservedTime(formats strfmt.Registry) error {

	if err := validate.Required("lastObservedTime", "body", m.LastObservedTime); err != nil {
		return err
	}

	if err := validate.Required("lastObservedTime", "body", m.LastObservedTime); err != nil {
		return err
	}

	if m.LastObservedTime != nil {
		if err := m.LastObservedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastObservedTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastObservedTime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api events v1beta1 event series based on the context it is used
func (m *IoK8sAPIEventsV1beta1EventSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastObservedTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIEventsV1beta1EventSeries) contextValidateLastObservedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastObservedTime != nil {
		if err := m.LastObservedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastObservedTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastObservedTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIEventsV1beta1EventSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIEventsV1beta1EventSeries) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIEventsV1beta1EventSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
