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

// IoK8sAPICoreV1HTTPHeader HTTPHeader describes a custom header to be used in HTTP probes
//
// swagger:model io.k8s.api.core.v1.HTTPHeader
type IoK8sAPICoreV1HTTPHeader struct {

	// The header field name
	// Required: true
	Name *string `json:"name"`

	// The header field value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this io k8s api core v1 HTTP header
func (m *IoK8sAPICoreV1HTTPHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1HTTPHeader) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1HTTPHeader) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 HTTP header based on context it is used
func (m *IoK8sAPICoreV1HTTPHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1HTTPHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1HTTPHeader) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1HTTPHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
