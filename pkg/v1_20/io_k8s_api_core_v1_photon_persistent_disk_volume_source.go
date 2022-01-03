// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1PhotonPersistentDiskVolumeSource Represents a Photon Controller persistent disk resource.
//
// swagger:model io.k8s.api.core.v1.PhotonPersistentDiskVolumeSource
type IoK8sAPICoreV1PhotonPersistentDiskVolumeSource struct {

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType string `json:"fsType,omitempty" json,yaml:"fsType,omitempty"`

	// ID that identifies Photon Controller persistent disk
	// Required: true
	PdID *string `json:"pdID" json,yaml:"pdID"`
}

// Validate validates this io k8s api core v1 photon persistent disk volume source
func (m *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePdID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource) validatePdID(formats strfmt.Registry) error {

	if err := validate.Required("pdID", "body", m.PdID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 photon persistent disk volume source based on context it is used
func (m *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PhotonPersistentDiskVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
