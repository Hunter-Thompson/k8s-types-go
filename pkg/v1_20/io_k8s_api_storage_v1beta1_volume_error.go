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

// IoK8sAPIStorageV1beta1VolumeError VolumeError captures an error encountered during a volume operation.
//
// swagger:model io.k8s.api.storage.v1beta1.VolumeError
type IoK8sAPIStorageV1beta1VolumeError struct {

	// String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
	Message string `json:"message,omitempty" json,yaml:"message,omitempty"`

	// Time the error was encountered.
	// Format: date-time
	Time IoK8sApimachineryPkgApisMetaV1Time `json:"time,omitempty" json,yaml:"time,omitempty"`
}

// Validate validates this io k8s api storage v1beta1 volume error
func (m *IoK8sAPIStorageV1beta1VolumeError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeError) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("time")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api storage v1beta1 volume error based on the context it is used
func (m *IoK8sAPIStorageV1beta1VolumeError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeError) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Time.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("time")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeError) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1beta1VolumeError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
