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

// IoK8sAPIAppsV1DeploymentStrategy DeploymentStrategy describes how to replace existing pods with new ones.
//
// swagger:model io.k8s.api.apps.v1.DeploymentStrategy
type IoK8sAPIAppsV1DeploymentStrategy struct {

	// Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
	RollingUpdate *IoK8sAPIAppsV1RollingUpdateDeployment `json:"rollingUpdate,omitempty"`

	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	//
	// Possible enum values:
	//  - `"Recreate"` Kill all existing pods before creating new ones.
	//  - `"RollingUpdate"` Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	// Enum: [Recreate RollingUpdate]
	Type string `json:"type,omitempty"`
}

// Validate validates this io k8s api apps v1 deployment strategy
func (m *IoK8sAPIAppsV1DeploymentStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollingUpdate(formats); err != nil {
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

func (m *IoK8sAPIAppsV1DeploymentStrategy) validateRollingUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.RollingUpdate) { // not required
		return nil
	}

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

var ioK8sApiAppsV1DeploymentStrategyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Recreate","RollingUpdate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioK8sApiAppsV1DeploymentStrategyTypeTypePropEnum = append(ioK8sApiAppsV1DeploymentStrategyTypeTypePropEnum, v)
	}
}

const (

	// IoK8sAPIAppsV1DeploymentStrategyTypeRecreate captures enum value "Recreate"
	IoK8sAPIAppsV1DeploymentStrategyTypeRecreate string = "Recreate"

	// IoK8sAPIAppsV1DeploymentStrategyTypeRollingUpdate captures enum value "RollingUpdate"
	IoK8sAPIAppsV1DeploymentStrategyTypeRollingUpdate string = "RollingUpdate"
)

// prop value enum
func (m *IoK8sAPIAppsV1DeploymentStrategy) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ioK8sApiAppsV1DeploymentStrategyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IoK8sAPIAppsV1DeploymentStrategy) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api apps v1 deployment strategy based on the context it is used
func (m *IoK8sAPIAppsV1DeploymentStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRollingUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DeploymentStrategy) contextValidateRollingUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DeploymentStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DeploymentStrategy) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1DeploymentStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
