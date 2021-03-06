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
)

// IoK8sAPIAppsV1DeploymentStatus DeploymentStatus is the most recently observed status of the Deployment.
//
// swagger:model io.k8s.api.apps.v1.DeploymentStatus
type IoK8sAPIAppsV1DeploymentStatus struct {

	// Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
	AvailableReplicas int32 `json:"availableReplicas,omitempty" json,yaml:"availableReplicas,omitempty"`

	// Count of hash collisions for the Deployment. The Deployment controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ReplicaSet.
	CollisionCount int32 `json:"collisionCount,omitempty" json,yaml:"collisionCount,omitempty"`

	// Represents the latest available observations of a deployment's current state.
	Conditions []*IoK8sAPIAppsV1DeploymentCondition `json:"conditions" json,yaml:"conditions"`

	// The generation observed by the deployment controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty" json,yaml:"observedGeneration,omitempty"`

	// Total number of ready pods targeted by this deployment.
	ReadyReplicas int32 `json:"readyReplicas,omitempty" json,yaml:"readyReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	Replicas int32 `json:"replicas,omitempty" json,yaml:"replicas,omitempty"`

	// Total number of unavailable pods targeted by this deployment. This is the total number of pods that are still required for the deployment to have 100% available capacity. They may either be pods that are running but not yet available or pods that still have not been created.
	UnavailableReplicas int32 `json:"unavailableReplicas,omitempty" json,yaml:"unavailableReplicas,omitempty"`

	// Total number of non-terminated pods targeted by this deployment that have the desired template spec.
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty" json,yaml:"updatedReplicas,omitempty"`
}

// Validate validates this io k8s api apps v1 deployment status
func (m *IoK8sAPIAppsV1DeploymentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DeploymentStatus) validateConditions(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api apps v1 deployment status based on the context it is used
func (m *IoK8sAPIAppsV1DeploymentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DeploymentStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPIAppsV1DeploymentStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DeploymentStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1DeploymentStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
