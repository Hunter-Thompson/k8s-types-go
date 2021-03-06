// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIPolicyV1beta1AllowedHostPath AllowedHostPath defines the host volume conditions that will be enabled by a policy for pods to use. It requires the path prefix to be defined.
//
// swagger:model io.k8s.api.policy.v1beta1.AllowedHostPath
type IoK8sAPIPolicyV1beta1AllowedHostPath struct {

	// pathPrefix is the path prefix that the host volume must match. It does not support `*`. Trailing slashes are trimmed when validating the path prefix with a host path.
	//
	// Examples: `/foo` would allow `/foo`, `/foo/` and `/foo/bar` `/foo` would not allow `/food` or `/etc/foo`
	PathPrefix string `json:"pathPrefix,omitempty" json,yaml:"pathPrefix,omitempty"`

	// when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly.
	ReadOnly bool `json:"readOnly,omitempty" json,yaml:"readOnly,omitempty"`
}

// Validate validates this io k8s api policy v1beta1 allowed host path
func (m *IoK8sAPIPolicyV1beta1AllowedHostPath) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api policy v1beta1 allowed host path based on context it is used
func (m *IoK8sAPIPolicyV1beta1AllowedHostPath) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1AllowedHostPath) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1AllowedHostPath) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1AllowedHostPath
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
