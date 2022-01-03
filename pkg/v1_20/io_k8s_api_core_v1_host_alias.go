// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1HostAlias HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
//
// swagger:model io.k8s.api.core.v1.HostAlias
type IoK8sAPICoreV1HostAlias struct {

	// Hostnames for the above IP address.
	Hostnames []string `json:"hostnames"`

	// IP address of the host file entry.
	IP string `json:"ip,omitempty"`
}

// Validate validates this io k8s api core v1 host alias
func (m *IoK8sAPICoreV1HostAlias) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 host alias based on context it is used
func (m *IoK8sAPICoreV1HostAlias) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1HostAlias) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1HostAlias) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1HostAlias
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
