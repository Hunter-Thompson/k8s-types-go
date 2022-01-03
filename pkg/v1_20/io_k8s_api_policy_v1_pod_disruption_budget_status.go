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

// IoK8sAPIPolicyV1PodDisruptionBudgetStatus PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
//
// swagger:model io.k8s.api.policy.v1.PodDisruptionBudgetStatus
type IoK8sAPIPolicyV1PodDisruptionBudgetStatus struct {

	// Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute
	//               the number of allowed disruptions. Therefore no disruptions are
	//               allowed and the status of the condition will be False.
	// - InsufficientPods: The number of pods are either at or below the number
	//                     required by the PodDisruptionBudget. No disruptions are
	//                     allowed and the status of the condition will be False.
	// - SufficientPods: There are more pods than required by the PodDisruptionBudget.
	//                   The condition will be True, and the number of allowed
	//                   disruptions are provided by the disruptionsAllowed property.
	Conditions []*IoK8sApimachineryPkgApisMetaV1Condition `json:"conditions"`

	// current number of healthy pods
	// Required: true
	CurrentHealthy *int32 `json:"currentHealthy"`

	// minimum desired number of healthy pods
	// Required: true
	DesiredHealthy *int32 `json:"desiredHealthy"`

	// DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	DisruptedPods map[string]IoK8sApimachineryPkgApisMetaV1Time `json:"disruptedPods,omitempty"`

	// Number of pod disruptions that are currently allowed.
	// Required: true
	DisruptionsAllowed *int32 `json:"disruptionsAllowed"`

	// total number of pods counted by this disruption budget
	// Required: true
	ExpectedPods *int32 `json:"expectedPods"`

	// Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// Validate validates this io k8s api policy v1 pod disruption budget status
func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruptedPods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruptionsAllowed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedPods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateCurrentHealthy(formats strfmt.Registry) error {

	if err := validate.Required("currentHealthy", "body", m.CurrentHealthy); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateDesiredHealthy(formats strfmt.Registry) error {

	if err := validate.Required("desiredHealthy", "body", m.DesiredHealthy); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateDisruptedPods(formats strfmt.Registry) error {
	if swag.IsZero(m.DisruptedPods) { // not required
		return nil
	}

	for k := range m.DisruptedPods {

		if val, ok := m.DisruptedPods[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateDisruptionsAllowed(formats strfmt.Registry) error {

	if err := validate.Required("disruptionsAllowed", "body", m.DisruptionsAllowed); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) validateExpectedPods(formats strfmt.Registry) error {

	if err := validate.Required("expectedPods", "body", m.ExpectedPods); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api policy v1 pod disruption budget status based on the context it is used
func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisruptedPods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) contextValidateDisruptedPods(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.DisruptedPods {

		if val, ok := m.DisruptedPods[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1PodDisruptionBudgetStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1PodDisruptionBudgetStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
