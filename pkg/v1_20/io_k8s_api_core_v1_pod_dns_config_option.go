// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1PodDNSConfigOption PodDNSConfigOption defines DNS resolver options of a pod.
//
// swagger:model io.k8s.api.core.v1.PodDNSConfigOption
type IoK8sAPICoreV1PodDNSConfigOption struct {

	// Required.
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this io k8s api core v1 pod DNS config option
func (m *IoK8sAPICoreV1PodDNSConfigOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 pod DNS config option based on context it is used
func (m *IoK8sAPICoreV1PodDNSConfigOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodDNSConfigOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodDNSConfigOption) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodDNSConfigOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
