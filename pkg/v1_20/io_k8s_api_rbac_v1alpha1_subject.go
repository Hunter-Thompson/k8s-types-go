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

// IoK8sAPIRbacV1alpha1Subject Subject contains a reference to the object or user identities a role binding applies to.  This can either hold a direct API object reference, or a value for non-objects such as user and group names.
//
// swagger:model io.k8s.api.rbac.v1alpha1.Subject
type IoK8sAPIRbacV1alpha1Subject struct {

	// APIVersion holds the API group and version of the referenced subject. Defaults to "v1" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io/v1alpha1" for User and Group subjects.
	APIVersion string `json:"apiVersion,omitempty" json,yaml:"apiVersion,omitempty"`

	// Kind of object being referenced. Values defined by this API group are "User", "Group", and "ServiceAccount". If the Authorizer does not recognized the kind value, the Authorizer should report an error.
	// Required: true
	Kind *string `json:"kind" json,yaml:"kind"`

	// Name of the object being referenced.
	// Required: true
	Name *string `json:"name" json,yaml:"name"`

	// Namespace of the referenced object.  If the object kind is non-namespace, such as "User" or "Group", and this value is not empty the Authorizer should report an error.
	Namespace string `json:"namespace,omitempty" json,yaml:"namespace,omitempty"`
}

// Validate validates this io k8s api rbac v1alpha1 subject
func (m *IoK8sAPIRbacV1alpha1Subject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIRbacV1alpha1Subject) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIRbacV1alpha1Subject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api rbac v1alpha1 subject based on context it is used
func (m *IoK8sAPIRbacV1alpha1Subject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIRbacV1alpha1Subject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIRbacV1alpha1Subject) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIRbacV1alpha1Subject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
