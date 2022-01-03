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

// IoK8sAPICoreV1ConfigMapKeySelector Selects a key from a ConfigMap.
//
// swagger:model io.k8s.api.core.v1.ConfigMapKeySelector
type IoK8sAPICoreV1ConfigMapKeySelector struct {

	// The key to select.
	// Required: true
	Key *string `json:"key" json,yaml:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty" json,yaml:"name,omitempty"`

	// Specify whether the ConfigMap or its key must be defined
	Optional bool `json:"optional,omitempty" json,yaml:"optional,omitempty"`
}

// Validate validates this io k8s api core v1 config map key selector
func (m *IoK8sAPICoreV1ConfigMapKeySelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ConfigMapKeySelector) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 config map key selector based on context it is used
func (m *IoK8sAPICoreV1ConfigMapKeySelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ConfigMapKeySelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ConfigMapKeySelector) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ConfigMapKeySelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
