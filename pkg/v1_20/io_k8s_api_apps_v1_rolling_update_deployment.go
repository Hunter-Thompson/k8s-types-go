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

// IoK8sAPIAppsV1RollingUpdateDeployment Spec to control the desired behavior of rolling update.
//
// swagger:model io.k8s.api.apps.v1.RollingUpdateDeployment
type IoK8sAPIAppsV1RollingUpdateDeployment struct {

	// The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge IoK8sApimachineryPkgUtilIntstrIntOrString `json:"maxSurge,omitempty"`

	// The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.
	MaxUnavailable IoK8sApimachineryPkgUtilIntstrIntOrString `json:"maxUnavailable,omitempty"`
}

// Validate validates this io k8s api apps v1 rolling update deployment
func (m *IoK8sAPIAppsV1RollingUpdateDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxSurge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUnavailable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1RollingUpdateDeployment) validateMaxSurge(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxSurge) { // not required
		return nil
	}

	if err := m.MaxSurge.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maxSurge")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maxSurge")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAppsV1RollingUpdateDeployment) validateMaxUnavailable(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxUnavailable) { // not required
		return nil
	}

	if err := m.MaxUnavailable.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maxUnavailable")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maxUnavailable")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api apps v1 rolling update deployment based on the context it is used
func (m *IoK8sAPIAppsV1RollingUpdateDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaxSurge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxUnavailable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1RollingUpdateDeployment) contextValidateMaxSurge(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxSurge.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maxSurge")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maxSurge")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAppsV1RollingUpdateDeployment) contextValidateMaxUnavailable(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxUnavailable.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("maxUnavailable")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("maxUnavailable")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1RollingUpdateDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1RollingUpdateDeployment) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1RollingUpdateDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
