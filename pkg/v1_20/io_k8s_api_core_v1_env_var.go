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

// IoK8sAPICoreV1EnvVar EnvVar represents an environment variable present in a Container.
//
// swagger:model io.k8s.api.core.v1.EnvVar
type IoK8sAPICoreV1EnvVar struct {

	// Name of the environment variable. Must be a C_IDENTIFIER.
	// Required: true
	Name *string `json:"name"`

	// Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	Value string `json:"value,omitempty"`

	// Source for the environment variable's value. Cannot be used if value is not empty.
	ValueFrom *IoK8sAPICoreV1EnvVarSource `json:"valueFrom,omitempty"`
}

// Validate validates this io k8s api core v1 env var
func (m *IoK8sAPICoreV1EnvVar) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EnvVar) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVar) validateValueFrom(formats strfmt.Registry) error {
	if swag.IsZero(m.ValueFrom) { // not required
		return nil
	}

	if m.ValueFrom != nil {
		if err := m.ValueFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 env var based on the context it is used
func (m *IoK8sAPICoreV1EnvVar) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValueFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EnvVar) contextValidateValueFrom(ctx context.Context, formats strfmt.Registry) error {

	if m.ValueFrom != nil {
		if err := m.ValueFrom.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("valueFrom")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("valueFrom")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EnvVar) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EnvVar) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EnvVar
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}