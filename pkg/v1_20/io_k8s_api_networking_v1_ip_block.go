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

// IoK8sAPINetworkingV1IPBlock IPBlock describes a particular CIDR (Ex. "192.168.1.1/24","2001:db9::/64") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
//
// swagger:model io.k8s.api.networking.v1.IPBlock
type IoK8sAPINetworkingV1IPBlock struct {

	// CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64"
	// Required: true
	Cidr *string `json:"cidr" json,yaml:"cidr"`

	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are "192.168.1.1/24" or "2001:db9::/64" Except values will be rejected if they are outside the CIDR range
	Except []string `json:"except" json,yaml:"except"`
}

// Validate validates this io k8s api networking v1 IP block
func (m *IoK8sAPINetworkingV1IPBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1IPBlock) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api networking v1 IP block based on context it is used
func (m *IoK8sAPINetworkingV1IPBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IPBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1IPBlock) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1IPBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
