// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sApimachineryPkgApisMetaV1Preconditions Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions
type IoK8sApimachineryPkgApisMetaV1Preconditions struct {

	// Specifies the target ResourceVersion
	ResourceVersion string `json:"resourceVersion,omitempty" json,yaml:"resourceVersion,omitempty"`

	// Specifies the target UID.
	UID string `json:"uid,omitempty" json,yaml:"uid,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 preconditions
func (m *IoK8sApimachineryPkgApisMetaV1Preconditions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 preconditions based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1Preconditions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1Preconditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1Preconditions) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1Preconditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
