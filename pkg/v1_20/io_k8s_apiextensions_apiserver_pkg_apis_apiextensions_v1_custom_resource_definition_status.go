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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinitionStatus
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus struct {

	// acceptedNames are the names that are actually being used to serve discovery. They may be different than the names in spec.
	AcceptedNames *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionNames `json:"acceptedNames,omitempty" json,yaml:"acceptedNames,omitempty"`

	// conditions indicate state for particular aspects of a CustomResourceDefinition
	Conditions []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionCondition `json:"conditions" json,yaml:"conditions"`

	// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from `spec.versions` while they exist in this list.
	StoredVersions []string `json:"storedVersions" json,yaml:"storedVersions"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource definition status
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) validateAcceptedNames(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedNames) { // not required
		return nil
	}

	if m.AcceptedNames != nil {
		if err := m.AcceptedNames.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedNames")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedNames")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) validateConditions(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource definition status based on the context it is used
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcceptedNames(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) contextValidateAcceptedNames(ctx context.Context, formats strfmt.Registry) error {

	if m.AcceptedNames != nil {
		if err := m.AcceptedNames.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptedNames")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acceptedNames")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
