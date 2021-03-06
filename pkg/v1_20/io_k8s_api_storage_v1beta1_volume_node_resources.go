// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIStorageV1beta1VolumeNodeResources VolumeNodeResources is a set of resource limits for scheduling of volumes.
//
// swagger:model io.k8s.api.storage.v1beta1.VolumeNodeResources
type IoK8sAPIStorageV1beta1VolumeNodeResources struct {

	// Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is nil, then the supported number of volumes on this node is unbounded.
	Count int32 `json:"count,omitempty" json,yaml:"count,omitempty"`
}

// Validate validates this io k8s api storage v1beta1 volume node resources
func (m *IoK8sAPIStorageV1beta1VolumeNodeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api storage v1beta1 volume node resources based on context it is used
func (m *IoK8sAPIStorageV1beta1VolumeNodeResources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeNodeResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeNodeResources) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1beta1VolumeNodeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
