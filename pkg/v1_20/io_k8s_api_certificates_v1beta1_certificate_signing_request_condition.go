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

// IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition io k8s api certificates v1beta1 certificate signing request condition
//
// swagger:model io.k8s.api.certificates.v1beta1.CertificateSigningRequestCondition
type IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition struct {

	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty" json,yaml:"lastTransitionTime,omitempty"`

	// timestamp for the last update to this condition
	// Format: date-time
	LastUpdateTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastUpdateTime,omitempty" json,yaml:"lastUpdateTime,omitempty"`

	// human readable message with details about the request state
	Message string `json:"message,omitempty" json,yaml:"message,omitempty"`

	// brief reason for the request state
	Reason string `json:"reason,omitempty" json,yaml:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown". Defaults to "True". If unset, should be treated as "True".
	Status string `json:"status,omitempty" json,yaml:"status,omitempty"`

	// type of the condition. Known conditions include "Approved", "Denied", and "Failed".
	// Required: true
	Type *string `json:"type" json,yaml:"type"`
}

// Validate validates this io k8s api certificates v1beta1 certificate signing request condition
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateLastUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdateTime) { // not required
		return nil
	}

	if err := m.LastUpdateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastUpdateTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastUpdateTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api certificates v1beta1 certificate signing request condition based on the context it is used
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) contextValidateLastUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastUpdateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastUpdateTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastUpdateTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
