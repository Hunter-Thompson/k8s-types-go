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

// IoK8sAPICoreV1NFSVolumeSource Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.NFSVolumeSource
type IoK8sAPICoreV1NFSVolumeSource struct {

	// Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// Required: true
	Path *string `json:"path"`

	// ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	ReadOnly bool `json:"readOnly,omitempty"`

	// Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// Required: true
	Server *string `json:"server"`
}

// Validate validates this io k8s api core v1 n f s volume source
func (m *IoK8sAPICoreV1NFSVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NFSVolumeSource) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NFSVolumeSource) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 n f s volume source based on context it is used
func (m *IoK8sAPICoreV1NFSVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1NFSVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1NFSVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1NFSVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}