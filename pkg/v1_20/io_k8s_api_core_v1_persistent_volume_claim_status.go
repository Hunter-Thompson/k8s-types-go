// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1PersistentVolumeClaimStatus PersistentVolumeClaimStatus is the current status of a persistent volume claim.
//
// swagger:model io.k8s.api.core.v1.PersistentVolumeClaimStatus
type IoK8sAPICoreV1PersistentVolumeClaimStatus struct {

	// AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes"`

	// The storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
	AllocatedResources map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"allocatedResources,omitempty"`

	// Represents the actual resources of the underlying volume.
	Capacity map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"capacity,omitempty"`

	// Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
	Conditions []*IoK8sAPICoreV1PersistentVolumeClaimCondition `json:"conditions"`

	// Phase represents the current phase of PersistentVolumeClaim.
	//
	// Possible enum values:
	//  - `"Bound"` used for PersistentVolumeClaims that are bound
	//  - `"Lost"` used for PersistentVolumeClaims that lost their underlying PersistentVolume. The claim was bound to a PersistentVolume and this volume does not exist any longer and all data on it was lost.
	//  - `"Pending"` used for PersistentVolumeClaims that are not yet bound
	// Enum: [Bound Lost Pending]
	Phase string `json:"phase,omitempty"`

	// ResizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
	ResizeStatus string `json:"resizeStatus,omitempty"`
}

// Validate validates this io k8s api core v1 persistent volume claim status
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validateAllocatedResources(formats strfmt.Registry) error {
	if swag.IsZero(m.AllocatedResources) { // not required
		return nil
	}

	for k := range m.AllocatedResources {

		if val, ok := m.AllocatedResources[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validateConditions(formats strfmt.Registry) error {
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

var ioK8sApiCoreV1PersistentVolumeClaimStatusTypePhasePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Bound","Lost","Pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioK8sApiCoreV1PersistentVolumeClaimStatusTypePhasePropEnum = append(ioK8sApiCoreV1PersistentVolumeClaimStatusTypePhasePropEnum, v)
	}
}

const (

	// IoK8sAPICoreV1PersistentVolumeClaimStatusPhaseBound captures enum value "Bound"
	IoK8sAPICoreV1PersistentVolumeClaimStatusPhaseBound string = "Bound"

	// IoK8sAPICoreV1PersistentVolumeClaimStatusPhaseLost captures enum value "Lost"
	IoK8sAPICoreV1PersistentVolumeClaimStatusPhaseLost string = "Lost"

	// IoK8sAPICoreV1PersistentVolumeClaimStatusPhasePending captures enum value "Pending"
	IoK8sAPICoreV1PersistentVolumeClaimStatusPhasePending string = "Pending"
)

// prop value enum
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validatePhaseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ioK8sApiCoreV1PersistentVolumeClaimStatusTypePhasePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 persistent volume claim status based on the context it is used
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocatedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) contextValidateAllocatedResources(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.AllocatedResources {

		if val, ok := m.AllocatedResources[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolumeClaimStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}