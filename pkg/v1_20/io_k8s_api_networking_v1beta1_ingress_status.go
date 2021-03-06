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

// IoK8sAPINetworkingV1beta1IngressStatus IngressStatus describe the current state of the Ingress.
//
// swagger:model io.k8s.api.networking.v1beta1.IngressStatus
type IoK8sAPINetworkingV1beta1IngressStatus struct {

	// LoadBalancer contains the current status of the load-balancer.
	LoadBalancer *IoK8sAPICoreV1LoadBalancerStatus `json:"loadBalancer,omitempty" json,yaml:"loadBalancer,omitempty"`
}

// Validate validates this io k8s api networking v1beta1 ingress status
func (m *IoK8sAPINetworkingV1beta1IngressStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBalancer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressStatus) validateLoadBalancer(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadBalancer) { // not required
		return nil
	}

	if m.LoadBalancer != nil {
		if err := m.LoadBalancer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadBalancer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loadBalancer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api networking v1beta1 ingress status based on the context it is used
func (m *IoK8sAPINetworkingV1beta1IngressStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoadBalancer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressStatus) contextValidateLoadBalancer(ctx context.Context, formats strfmt.Registry) error {

	if m.LoadBalancer != nil {
		if err := m.LoadBalancer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loadBalancer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("loadBalancer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1beta1IngressStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
