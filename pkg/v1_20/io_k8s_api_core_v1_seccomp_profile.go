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

// IoK8sAPICoreV1SeccompProfile SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.
//
// swagger:model io.k8s.api.core.v1.SeccompProfile
type IoK8sAPICoreV1SeccompProfile struct {

	// localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is "Localhost".
	LocalhostProfile string `json:"localhostProfile,omitempty"`

	// type indicates which kind of seccomp profile will be applied. Valid options are:
	//
	// Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.
	//
	// Possible enum values:
	//  - `"Localhost"` indicates a profile defined in a file on the node should be used. The file's location relative to <kubelet-root-dir>/seccomp.
	//  - `"RuntimeDefault"` represents the default container runtime seccomp profile.
	//  - `"Unconfined"` indicates no seccomp profile is applied (A.K.A. unconfined).
	// Required: true
	// Enum: [Localhost RuntimeDefault Unconfined]
	Type *string `json:"type"`
}

// Validate validates this io k8s api core v1 seccomp profile
func (m *IoK8sAPICoreV1SeccompProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ioK8sApiCoreV1SeccompProfileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Localhost","RuntimeDefault","Unconfined"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioK8sApiCoreV1SeccompProfileTypeTypePropEnum = append(ioK8sApiCoreV1SeccompProfileTypeTypePropEnum, v)
	}
}

const (

	// IoK8sAPICoreV1SeccompProfileTypeLocalhost captures enum value "Localhost"
	IoK8sAPICoreV1SeccompProfileTypeLocalhost string = "Localhost"

	// IoK8sAPICoreV1SeccompProfileTypeRuntimeDefault captures enum value "RuntimeDefault"
	IoK8sAPICoreV1SeccompProfileTypeRuntimeDefault string = "RuntimeDefault"

	// IoK8sAPICoreV1SeccompProfileTypeUnconfined captures enum value "Unconfined"
	IoK8sAPICoreV1SeccompProfileTypeUnconfined string = "Unconfined"
)

// prop value enum
func (m *IoK8sAPICoreV1SeccompProfile) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ioK8sApiCoreV1SeccompProfileTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IoK8sAPICoreV1SeccompProfile) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 seccomp profile based on context it is used
func (m *IoK8sAPICoreV1SeccompProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1SeccompProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1SeccompProfile) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1SeccompProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
