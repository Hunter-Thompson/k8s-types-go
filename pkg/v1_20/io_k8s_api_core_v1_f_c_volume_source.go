// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1FCVolumeSource Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.FCVolumeSource
type IoK8sAPICoreV1FCVolumeSource struct {

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType string `json:"fsType,omitempty" json,yaml:"fsType,omitempty"`

	// Optional: FC target lun number
	Lun int32 `json:"lun,omitempty" json,yaml:"lun,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty" json,yaml:"readOnly,omitempty"`

	// Optional: FC target worldwide names (WWNs)
	TargetWWNs []string `json:"targetWWNs" json,yaml:"targetWWNs"`

	// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids []string `json:"wwids" json,yaml:"wwids"`
}

// Validate validates this io k8s api core v1 f c volume source
func (m *IoK8sAPICoreV1FCVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 f c volume source based on context it is used
func (m *IoK8sAPICoreV1FCVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1FCVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1FCVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1FCVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
