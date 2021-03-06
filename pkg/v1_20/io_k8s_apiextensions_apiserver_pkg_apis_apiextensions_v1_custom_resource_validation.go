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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation CustomResourceValidation is a list of validation methods for CustomResources.
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceValidation
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation struct {

	// openAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
	OpenAPIV3Schema *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"openAPIV3Schema,omitempty" json,yaml:"openAPIV3Schema,omitempty"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource validation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpenAPIV3Schema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) validateOpenAPIV3Schema(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenAPIV3Schema) { // not required
		return nil
	}

	if m.OpenAPIV3Schema != nil {
		if err := m.OpenAPIV3Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openAPIV3Schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openAPIV3Schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource validation based on the context it is used
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOpenAPIV3Schema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) contextValidateOpenAPIV3Schema(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenAPIV3Schema != nil {
		if err := m.OpenAPIV3Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openAPIV3Schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openAPIV3Schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
