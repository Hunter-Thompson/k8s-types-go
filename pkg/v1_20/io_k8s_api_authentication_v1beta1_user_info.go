// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAuthenticationV1beta1UserInfo UserInfo holds the information about the user needed to implement the user.Info interface.
//
// swagger:model io.k8s.api.authentication.v1beta1.UserInfo
type IoK8sAPIAuthenticationV1beta1UserInfo struct {

	// Any additional information provided by the authenticator.
	Extra map[string][]string `json:"extra,omitempty" json,yaml:"extra,omitempty"`

	// The names of groups this user is a part of.
	Groups []string `json:"groups" json,yaml:"groups"`

	// A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
	UID string `json:"uid,omitempty" json,yaml:"uid,omitempty"`

	// The name that uniquely identifies this user among all active users.
	Username string `json:"username,omitempty" json,yaml:"username,omitempty"`
}

// Validate validates this io k8s api authentication v1beta1 user info
func (m *IoK8sAPIAuthenticationV1beta1UserInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api authentication v1beta1 user info based on context it is used
func (m *IoK8sAPIAuthenticationV1beta1UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1UserInfo) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthenticationV1beta1UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
