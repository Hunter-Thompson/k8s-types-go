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

// IoK8sAPIPolicyV1beta1IDRange IDRange provides a min/max of an allowed range of IDs.
//
// swagger:model io.k8s.api.policy.v1beta1.IDRange
type IoK8sAPIPolicyV1beta1IDRange struct {

	// max is the end of the range, inclusive.
	// Required: true
	Max *int64 `json:"max" json,yaml:"max"`

	// min is the start of the range, inclusive.
	// Required: true
	Min *int64 `json:"min" json,yaml:"min"`
}

// Validate validates this io k8s api policy v1beta1 ID range
func (m *IoK8sAPIPolicyV1beta1IDRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1IDRange) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIPolicyV1beta1IDRange) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api policy v1beta1 ID range based on context it is used
func (m *IoK8sAPIPolicyV1beta1IDRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1IDRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1IDRange) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1IDRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
