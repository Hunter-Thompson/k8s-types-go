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
	"github.com/go-openapi/validate"
)

// IoK8sAPIAppsV1ReplicaSetStatus ReplicaSetStatus represents the current status of a ReplicaSet.
//
// swagger:model io.k8s.api.apps.v1.ReplicaSetStatus
type IoK8sAPIAppsV1ReplicaSetStatus struct {

	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	AvailableReplicas int32 `json:"availableReplicas,omitempty" json,yaml:"availableReplicas,omitempty"`

	// Represents the latest available observations of a replica set's current state.
	Conditions []*IoK8sAPIAppsV1ReplicaSetCondition `json:"conditions" json,yaml:"conditions"`

	// The number of pods that have labels matching the labels of the pod template of the replicaset.
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty" json,yaml:"fullyLabeledReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	ObservedGeneration int64 `json:"observedGeneration,omitempty" json,yaml:"observedGeneration,omitempty"`

	// The number of ready replicas for this replica set.
	ReadyReplicas int32 `json:"readyReplicas,omitempty" json,yaml:"readyReplicas,omitempty"`

	// Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	// Required: true
	Replicas *int32 `json:"replicas" json,yaml:"replicas"`
}

// Validate validates this io k8s api apps v1 replica set status
func (m *IoK8sAPIAppsV1ReplicaSetStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1ReplicaSetStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *IoK8sAPIAppsV1ReplicaSetStatus) validateReplicas(formats strfmt.Registry) error {

	if err := validate.Required("replicas", "body", m.Replicas); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api apps v1 replica set status based on the context it is used
func (m *IoK8sAPIAppsV1ReplicaSetStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1ReplicaSetStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPIAppsV1ReplicaSetStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1ReplicaSetStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1ReplicaSetStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
