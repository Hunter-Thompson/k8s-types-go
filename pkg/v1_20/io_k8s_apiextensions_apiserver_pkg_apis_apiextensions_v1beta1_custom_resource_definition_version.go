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
	"github.com/go-openapi/validate"
)

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion CustomResourceDefinitionVersion describes a version for CRD.
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinitionVersion
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion struct {

	// additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. Top-level and per-version columns are mutually exclusive. Per-version columns must not all be set to identical values (top-level columns should be used instead). If no top-level or per-version columns are specified, a single column displaying the age of the custom resource is used.
	AdditionalPrinterColumns []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition `json:"additionalPrinterColumns" json,yaml:"additionalPrinterColumns"`

	// deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
	Deprecated bool `json:"deprecated,omitempty" json,yaml:"deprecated,omitempty"`

	// deprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
	DeprecationWarning string `json:"deprecationWarning,omitempty" json,yaml:"deprecationWarning,omitempty"`

	// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	// Required: true
	Name *string `json:"name" json,yaml:"name"`

	// schema describes the schema used for validation and pruning of this version of the custom resource. Top-level and per-version schemas are mutually exclusive. Per-version schemas must not all be set to identical values (top-level validation schema should be used instead).
	Schema *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceValidation `json:"schema,omitempty" json,yaml:"schema,omitempty"`

	// served is a flag enabling/disabling this version from being served via REST APIs
	// Required: true
	Served *bool `json:"served" json,yaml:"served"`

	// storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
	// Required: true
	Storage *bool `json:"storage" json,yaml:"storage"`

	// subresources specify what subresources this version of the defined custom resource have. Top-level and per-version subresources are mutually exclusive. Per-version subresources must not all be set to identical values (top-level subresources should be used instead).
	Subresources *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceSubresources `json:"subresources,omitempty" json,yaml:"subresources,omitempty"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1beta1 custom resource definition version
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalPrinterColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubresources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateAdditionalPrinterColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalPrinterColumns) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalPrinterColumns); i++ {
		if swag.IsZero(m.AdditionalPrinterColumns[i]) { // not required
			continue
		}

		if m.AdditionalPrinterColumns[i] != nil {
			if err := m.AdditionalPrinterColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalPrinterColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalPrinterColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateServed(formats strfmt.Registry) error {

	if err := validate.Required("served", "body", m.Served); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) validateSubresources(formats strfmt.Registry) error {
	if swag.IsZero(m.Subresources) { // not required
		return nil
	}

	if m.Subresources != nil {
		if err := m.Subresources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subresources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subresources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s apiextensions apiserver pkg apis apiextensions v1beta1 custom resource definition version based on the context it is used
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalPrinterColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubresources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) contextValidateAdditionalPrinterColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalPrinterColumns); i++ {

		if m.AdditionalPrinterColumns[i] != nil {
			if err := m.AdditionalPrinterColumns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalPrinterColumns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalPrinterColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {
		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) contextValidateSubresources(ctx context.Context, formats strfmt.Registry) error {

	if m.Subresources != nil {
		if err := m.Subresources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subresources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subresources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinitionVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
