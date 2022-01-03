// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1PodCondition PodCondition contains details for the current condition of this pod.
//
// swagger:model io.k8s.api.core.v1.PodCondition
type IoK8sAPICoreV1PodCondition struct {

	// Last time we probed the condition.
	// Format: date-time
	LastProbeTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastProbeTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	// Required: true
	Status *string `json:"status"`

	// Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	//
	// Possible enum values:
	//  - `"ContainersReady"` indicates whether all containers in the pod are ready.
	//  - `"Initialized"` means that all init containers in the pod have started successfully.
	//  - `"PodScheduled"` represents status of the scheduling process for this pod.
	//  - `"Ready"` means the pod is able to service requests and should be added to the load balancing pools of all matching services.
	// Required: true
	// Enum: [ContainersReady Initialized PodScheduled Ready]
	Type *string `json:"type"`
}

// Validate validates this io k8s api core v1 pod condition
func (m *IoK8sAPICoreV1PodCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastProbeTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
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

func (m *IoK8sAPICoreV1PodCondition) validateLastProbeTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastProbeTime) { // not required
		return nil
	}

	if err := m.LastProbeTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PodCondition) validateLastTransitionTime(formats strfmt.Registry) error {
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

func (m *IoK8sAPICoreV1PodCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var ioK8sApiCoreV1PodConditionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ContainersReady","Initialized","PodScheduled","Ready"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioK8sApiCoreV1PodConditionTypeTypePropEnum = append(ioK8sApiCoreV1PodConditionTypeTypePropEnum, v)
	}
}

const (

	// IoK8sAPICoreV1PodConditionTypeContainersReady captures enum value "ContainersReady"
	IoK8sAPICoreV1PodConditionTypeContainersReady string = "ContainersReady"

	// IoK8sAPICoreV1PodConditionTypeInitialized captures enum value "Initialized"
	IoK8sAPICoreV1PodConditionTypeInitialized string = "Initialized"

	// IoK8sAPICoreV1PodConditionTypePodScheduled captures enum value "PodScheduled"
	IoK8sAPICoreV1PodConditionTypePodScheduled string = "PodScheduled"

	// IoK8sAPICoreV1PodConditionTypeReady captures enum value "Ready"
	IoK8sAPICoreV1PodConditionTypeReady string = "Ready"
)

// prop value enum
func (m *IoK8sAPICoreV1PodCondition) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ioK8sApiCoreV1PodConditionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IoK8sAPICoreV1PodCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 pod condition based on the context it is used
func (m *IoK8sAPICoreV1PodCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastProbeTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastTransitionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodCondition) contextValidateLastProbeTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastProbeTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastProbeTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastProbeTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1PodCondition) contextValidateLastTransitionTime(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICoreV1PodCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
