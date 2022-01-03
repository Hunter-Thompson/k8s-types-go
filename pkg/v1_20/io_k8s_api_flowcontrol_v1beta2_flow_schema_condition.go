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

// IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition FlowSchemaCondition describes conditions for a FlowSchema.
//
// swagger:model io.k8s.api.flowcontrol.v1beta2.FlowSchemaCondition
type IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition struct {

	// `lastTransitionTime` is the last time the condition transitioned from one status to another.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// `message` is a human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// `status` is the status of the condition. Can be True, False, Unknown. Required.
	Status string `json:"status,omitempty"`

	// `type` is the type of the condition. Required.
	Type string `json:"type,omitempty"`
}

// Validate validates this io k8s api flowcontrol v1beta2 flow schema condition
func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) validateLastTransitionTime(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api flowcontrol v1beta2 flow schema condition based on the context it is used
func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1beta2FlowSchemaCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
