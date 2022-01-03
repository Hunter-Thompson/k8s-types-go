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

// IoK8sAPICoreV1EnvVarSource EnvVarSource represents a source for the value of an EnvVar.
//
// swagger:model io.k8s.api.core.v1.EnvVarSource
type IoK8sAPICoreV1EnvVarSource struct {

	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *IoK8sAPICoreV1ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`

	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	FieldRef *IoK8sAPICoreV1ObjectFieldSelector `json:"fieldRef,omitempty"`

	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *IoK8sAPICoreV1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`

	// Selects a key of a secret in the pod's namespace
	SecretKeyRef *IoK8sAPICoreV1SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// Validate validates this io k8s api core v1 env var source
func (m *IoK8sAPICoreV1EnvVarSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigMapKeyRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceFieldRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKeyRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) validateConfigMapKeyRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigMapKeyRef) { // not required
		return nil
	}

	if m.ConfigMapKeyRef != nil {
		if err := m.ConfigMapKeyRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMapKeyRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMapKeyRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) validateFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldRef) { // not required
		return nil
	}

	if m.FieldRef != nil {
		if err := m.FieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) validateResourceFieldRef(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceFieldRef) { // not required
		return nil
	}

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) validateSecretKeyRef(formats strfmt.Registry) error {
	if swag.IsZero(m.SecretKeyRef) { // not required
		return nil
	}

	if m.SecretKeyRef != nil {
		if err := m.SecretKeyRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretKeyRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretKeyRef")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 env var source based on the context it is used
func (m *IoK8sAPICoreV1EnvVarSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigMapKeyRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceFieldRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecretKeyRef(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) contextValidateConfigMapKeyRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigMapKeyRef != nil {
		if err := m.ConfigMapKeyRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configMapKeyRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configMapKeyRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) contextValidateFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.FieldRef != nil {
		if err := m.FieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) contextValidateResourceFieldRef(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceFieldRef != nil {
		if err := m.ResourceFieldRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceFieldRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceFieldRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EnvVarSource) contextValidateSecretKeyRef(ctx context.Context, formats strfmt.Registry) error {

	if m.SecretKeyRef != nil {
		if err := m.SecretKeyRef.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretKeyRef")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secretKeyRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EnvVarSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EnvVarSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EnvVarSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
