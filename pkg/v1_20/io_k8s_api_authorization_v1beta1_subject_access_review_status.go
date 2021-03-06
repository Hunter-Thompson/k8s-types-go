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

// IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus SubjectAccessReviewStatus
//
// swagger:model io.k8s.api.authorization.v1beta1.SubjectAccessReviewStatus
type IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus struct {

	// Allowed is required. True if the action would be allowed, false otherwise.
	// Required: true
	Allowed *bool `json:"allowed" json,yaml:"allowed"`

	// Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.
	Denied bool `json:"denied,omitempty" json,yaml:"denied,omitempty"`

	// EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.
	EvaluationError string `json:"evaluationError,omitempty" json,yaml:"evaluationError,omitempty"`

	// Reason is optional.  It indicates why a request was allowed or denied.
	Reason string `json:"reason,omitempty" json,yaml:"reason,omitempty"`
}

// Validate validates this io k8s api authorization v1beta1 subject access review status
func (m *IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus) validateAllowed(formats strfmt.Registry) error {

	if err := validate.Required("allowed", "body", m.Allowed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api authorization v1beta1 subject access review status based on context it is used
func (m *IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthorizationV1beta1SubjectAccessReviewStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
