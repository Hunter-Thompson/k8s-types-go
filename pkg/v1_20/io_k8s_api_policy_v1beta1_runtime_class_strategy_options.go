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

// IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.
//
// swagger:model io.k8s.api.policy.v1beta1.RuntimeClassStrategyOptions
type IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions struct {

	// allowedRuntimeClassNames is an allowlist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
	// Required: true
	AllowedRuntimeClassNames []string `json:"allowedRuntimeClassNames" json,yaml:"allowedRuntimeClassNames"`

	// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
	DefaultRuntimeClassName string `json:"defaultRuntimeClassName,omitempty" json,yaml:"defaultRuntimeClassName,omitempty"`
}

// Validate validates this io k8s api policy v1beta1 runtime class strategy options
func (m *IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedRuntimeClassNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions) validateAllowedRuntimeClassNames(formats strfmt.Registry) error {

	if err := validate.Required("allowedRuntimeClassNames", "body", m.AllowedRuntimeClassNames); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api policy v1beta1 runtime class strategy options based on context it is used
func (m *IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1RuntimeClassStrategyOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
