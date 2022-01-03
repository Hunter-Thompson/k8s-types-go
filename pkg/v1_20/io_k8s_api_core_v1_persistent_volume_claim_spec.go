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

// IoK8sAPICoreV1PersistentVolumeClaimSpec PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes
//
// swagger:model io.k8s.api.core.v1.PersistentVolumeClaimSpec
type IoK8sAPICoreV1PersistentVolumeClaimSpec struct {

	// AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes" json,yaml:"accessModes"`

	// This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source.
	DataSource *IoK8sAPICoreV1TypedLocalObjectReference `json:"dataSource,omitempty" json,yaml:"dataSource,omitempty"`

	// Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	Resources *IoK8sAPICoreV1ResourceRequirements `json:"resources,omitempty" json,yaml:"resources,omitempty"`

	// A label query over volumes to consider for binding.
	Selector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"selector,omitempty" json,yaml:"selector,omitempty"`

	// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName string `json:"storageClassName,omitempty" json,yaml:"storageClassName,omitempty"`

	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode string `json:"volumeMode,omitempty" json,yaml:"volumeMode,omitempty"`

	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName string `json:"volumeName,omitempty" json,yaml:"volumeName,omitempty"`
}

// Validate validates this io k8s api core v1 persistent volume claim spec
func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) validateDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSource) { // not required
		return nil
	}

	if m.DataSource != nil {
		if err := m.DataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 persistent volume claim spec based on the context it is used
func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) contextValidateDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSource != nil {
		if err := m.DataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataSource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataSource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {
		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeClaimSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolumeClaimSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
