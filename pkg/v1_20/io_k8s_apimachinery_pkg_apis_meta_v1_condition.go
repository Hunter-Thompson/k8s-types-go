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

// IoK8sApimachineryPkgApisMetaV1Condition Condition contains details for one aspect of the current state of this API Resource.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.Condition
type IoK8sApimachineryPkgApisMetaV1Condition struct {

	// lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	// Required: true
	// Format: date-time
	LastTransitionTime *IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime" json,yaml:"lastTransitionTime"`

	// message is a human readable message indicating details about the transition. This may be an empty string.
	// Required: true
	Message *string `json:"message" json,yaml:"message"`

	// observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.
	ObservedGeneration int64 `json:"observedGeneration,omitempty" json,yaml:"observedGeneration,omitempty"`

	// reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.
	// Required: true
	Reason *string `json:"reason" json,yaml:"reason"`

	// status of the condition, one of True, False, Unknown.
	// Required: true
	Status *string `json:"status" json,yaml:"status"`

	// type of condition in CamelCase or in foo.example.com/CamelCase.
	// Required: true
	Type *string `json:"type" json,yaml:"type"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 condition
func (m *IoK8sApimachineryPkgApisMetaV1Condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *IoK8sApimachineryPkgApisMetaV1Condition) validateLastTransitionTime(formats strfmt.Registry) error {

	if err := validate.Required("lastTransitionTime", "body", m.LastTransitionTime); err != nil {
		return err
	}

	if err := validate.Required("lastTransitionTime", "body", m.LastTransitionTime); err != nil {
		return err
	}

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1Condition) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1Condition) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1Condition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1Condition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s apimachinery pkg apis meta v1 condition based on the context it is used
func (m *IoK8sApimachineryPkgApisMetaV1Condition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1Condition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

	if m.LastTransitionTime != nil {
		if err := m.LastTransitionTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastTransitionTime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastTransitionTime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1Condition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1Condition) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1Condition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
