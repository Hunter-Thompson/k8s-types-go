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

// IoK8sAPINetworkingV1IngressServiceBackend IngressServiceBackend references a Kubernetes Service as a Backend.
//
// swagger:model io.k8s.api.networking.v1.IngressServiceBackend
type IoK8sAPINetworkingV1IngressServiceBackend struct {

	// Name is the referenced service. The service must exist in the same namespace as the Ingress object.
	// Required: true
	Name *string `json:"name"`

	// Port of the referenced service. A port name or port number is required for a IngressServiceBackend.
	Port *IoK8sAPINetworkingV1ServiceBackendPort `json:"port,omitempty"`
}

// Validate validates this io k8s api networking v1 ingress service backend
func (m *IoK8sAPINetworkingV1IngressServiceBackend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1IngressServiceBackend) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPINetworkingV1IngressServiceBackend) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api networking v1 ingress service backend based on the context it is used
func (m *IoK8sAPINetworkingV1IngressServiceBackend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1IngressServiceBackend) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {
		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IngressServiceBackend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IngressServiceBackend) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1IngressServiceBackend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
