// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1ServiceStatus ServiceStatus represents the current status of a service.
//
// swagger:model io.k8s.api.core.v1.ServiceStatus
type IoK8sAPICoreV1ServiceStatus struct {

	// Current service state
	Conditions []*IoK8sApimachineryPkgApisMetaV1Condition `json:"conditions" json,yaml:"conditions"`

	// LoadBalancer contains the current status of the load-balancer, if one is present.
	LoadBalancer *IoK8sAPICoreV1LoadBalancerStatus `json:"loadBalancer,omitempty" json,yaml:"loadBalancer,omitempty"`
}

// Validate validates this io k8s api core v1 service status
func (m *IoK8sAPICoreV1ServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoadBalancer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ServiceStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1ServiceStatus) validateLoadBalancer(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api core v1 service status based on the context it is used
func (m *IoK8sAPICoreV1ServiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLoadBalancer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ServiceStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1ServiceStatus) contextValidateLoadBalancer(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICoreV1ServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ServiceStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
