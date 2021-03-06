// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1SecretReference SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
//
// swagger:model io.k8s.api.core.v1.SecretReference
type IoK8sAPICoreV1SecretReference struct {

	// Name is unique within a namespace to reference a secret resource.
	Name string `json:"name,omitempty" json,yaml:"name,omitempty"`

	// Namespace defines the space within which the secret name must be unique.
	Namespace string `json:"namespace,omitempty" json,yaml:"namespace,omitempty"`
}

// Validate validates this io k8s api core v1 secret reference
func (m *IoK8sAPICoreV1SecretReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 secret reference based on context it is used
func (m *IoK8sAPICoreV1SecretReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1SecretReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1SecretReference) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1SecretReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
