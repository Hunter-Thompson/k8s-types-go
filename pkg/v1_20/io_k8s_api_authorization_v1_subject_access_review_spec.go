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

// IoK8sAPIAuthorizationV1SubjectAccessReviewSpec SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set
//
// swagger:model io.k8s.api.authorization.v1.SubjectAccessReviewSpec
type IoK8sAPIAuthorizationV1SubjectAccessReviewSpec struct {

	// Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.
	Extra map[string][]string `json:"extra,omitempty" json,yaml:"extra,omitempty"`

	// Groups is the groups you're testing for.
	Groups []string `json:"groups" json,yaml:"groups"`

	// NonResourceAttributes describes information for a non-resource access request
	NonResourceAttributes *IoK8sAPIAuthorizationV1NonResourceAttributes `json:"nonResourceAttributes,omitempty" json,yaml:"nonResourceAttributes,omitempty"`

	// ResourceAuthorizationAttributes describes information for a resource access request
	ResourceAttributes *IoK8sAPIAuthorizationV1ResourceAttributes `json:"resourceAttributes,omitempty" json,yaml:"resourceAttributes,omitempty"`

	// UID information about the requesting user.
	UID string `json:"uid,omitempty" json,yaml:"uid,omitempty"`

	// User is the user you're testing for. If you specify "User" but not "Groups", then is it interpreted as "What if User were not a member of any groups
	User string `json:"user,omitempty" json,yaml:"user,omitempty"`
}

// Validate validates this io k8s api authorization v1 subject access review spec
func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNonResourceAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) validateNonResourceAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.NonResourceAttributes) { // not required
		return nil
	}

	if m.NonResourceAttributes != nil {
		if err := m.NonResourceAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonResourceAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonResourceAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) validateResourceAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceAttributes) { // not required
		return nil
	}

	if m.ResourceAttributes != nil {
		if err := m.ResourceAttributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceAttributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api authorization v1 subject access review spec based on the context it is used
func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNonResourceAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) contextValidateNonResourceAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.NonResourceAttributes != nil {
		if err := m.NonResourceAttributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonResourceAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonResourceAttributes")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) contextValidateResourceAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceAttributes != nil {
		if err := m.ResourceAttributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceAttributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceAttributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1SubjectAccessReviewSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthorizationV1SubjectAccessReviewSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
