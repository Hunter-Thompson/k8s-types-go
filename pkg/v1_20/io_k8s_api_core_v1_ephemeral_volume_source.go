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

// IoK8sAPICoreV1EphemeralVolumeSource Represents an ephemeral volume that is handled by a normal storage driver.
//
// swagger:model io.k8s.api.core.v1.EphemeralVolumeSource
type IoK8sAPICoreV1EphemeralVolumeSource struct {

	// Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be `<pod name>-<volume name>` where `<volume name>` is the name from the `PodSpec.Volumes` array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).
	//
	// An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.
	//
	// This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.
	//
	// Required, must not be nil.
	VolumeClaimTemplate *IoK8sAPICoreV1PersistentVolumeClaimTemplate `json:"volumeClaimTemplate,omitempty"`
}

// Validate validates this io k8s api core v1 ephemeral volume source
func (m *IoK8sAPICoreV1EphemeralVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeClaimTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EphemeralVolumeSource) validateVolumeClaimTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeClaimTemplate) { // not required
		return nil
	}

	if m.VolumeClaimTemplate != nil {
		if err := m.VolumeClaimTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeClaimTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeClaimTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 ephemeral volume source based on the context it is used
func (m *IoK8sAPICoreV1EphemeralVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVolumeClaimTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EphemeralVolumeSource) contextValidateVolumeClaimTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.VolumeClaimTemplate != nil {
		if err := m.VolumeClaimTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeClaimTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeClaimTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EphemeralVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EphemeralVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EphemeralVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
