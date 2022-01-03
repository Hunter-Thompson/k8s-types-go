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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec CustomResourceDefinitionSpec describes how a user wants their resource to appear
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionSpec
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec struct {

	// conversion defines conversion settings for the CRD.
	Conversion *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion `json:"conversion,omitempty"`

	// group is the API group of the defined custom resource. The custom resources are served under `/apis/<group>/...`. Must match the name of the CustomResourceDefinition (in the form `<names.plural>.<group>`).
	// Required: true
	Group *string `json:"group"`

	// names specify the resource and kind names for the custom resource.
	// Required: true
	Names *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames `json:"names"`

	// preserveUnknownFields indicates that object fields which are not specified in the OpenAPI schema should be preserved when persisting to storage. apiVersion, kind, metadata and known fields inside metadata are always preserved. This field is deprecated in favor of setting `x-preserve-unknown-fields` to true in `spec.versions[*].schema.openAPIV3Schema`. See https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#pruning-versus-preserving-unknown-fields for details.
	PreserveUnknownFields bool `json:"preserveUnknownFields,omitempty"`

	// scope indicates whether the defined custom resource is cluster- or namespace-scoped. Allowed values are `Cluster` and `Namespaced`.
	// Required: true
	Scope *string `json:"scope"`

	// versions is the list of all API versions of the defined custom resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	// Required: true
	Versions []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionVersion `json:"versions"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource definition spec
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) validateConversion(formats strfmt.Registry) error {
	if swag.IsZero(m.Conversion) { // not required
		return nil
	}

	if m.Conversion != nil {
		if err := m.Conversion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversion")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) validateNames(formats strfmt.Registry) error {

	if err := validate.Required("names", "body", m.Names); err != nil {
		return err
	}

	if m.Names != nil {
		if err := m.Names.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("names")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("names")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) validateVersions(formats strfmt.Registry) error {

	if err := validate.Required("versions", "body", m.Versions); err != nil {
		return err
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource definition spec based on the context it is used
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConversion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) contextValidateConversion(ctx context.Context, formats strfmt.Registry) error {

	if m.Conversion != nil {
		if err := m.Conversion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversion")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) contextValidateNames(ctx context.Context, formats strfmt.Registry) error {

	if m.Names != nil {
		if err := m.Names.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("names")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("names")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Versions); i++ {

		if m.Versions[i] != nil {
			if err := m.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}