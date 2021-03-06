// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPINetworkingV1beta1IngressClassSpec IngressClassSpec provides information about the class of an Ingress.
//
// swagger:model io.k8s.api.networking.v1beta1.IngressClassSpec
type IoK8sAPINetworkingV1beta1IngressClassSpec struct {

	// Controller refers to the name of the controller that should handle this class. This allows for different "flavors" that are controlled by the same controller. For example, you may have different Parameters for the same implementing controller. This should be specified as a domain-prefixed path no more than 250 characters in length, e.g. "acme.io/ingress-controller". This field is immutable.
	Controller string `json:"controller,omitempty" json,yaml:"controller,omitempty"`

	// Parameters is a link to a custom resource containing additional configuration for the controller. This is optional if the controller does not require extra parameters.
	Parameters *IoK8sAPICoreV1TypedLocalObjectReference `json:"parameters,omitempty" json,yaml:"parameters,omitempty"`
}

// Validate validates this io k8s api networking v1beta1 ingress class spec
func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api networking v1beta1 ingress class spec based on the context it is used
func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressClassSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1beta1IngressClassSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
