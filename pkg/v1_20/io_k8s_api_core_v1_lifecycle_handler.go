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

// IoK8sAPICoreV1LifecycleHandler LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
//
// swagger:model io.k8s.api.core.v1.LifecycleHandler
type IoK8sAPICoreV1LifecycleHandler struct {

	// Exec specifies the action to take.
	Exec *IoK8sAPICoreV1ExecAction `json:"exec,omitempty"`

	// HTTPGet specifies the http request to perform.
	HTTPGet *IoK8sAPICoreV1HTTPGetAction `json:"httpGet,omitempty"`

	// Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.
	TCPSocket *IoK8sAPICoreV1TCPSocketAction `json:"tcpSocket,omitempty"`
}

// Validate validates this io k8s api core v1 lifecycle handler
func (m *IoK8sAPICoreV1LifecycleHandler) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPGet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPSocket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) validateExec(formats strfmt.Registry) error {
	if swag.IsZero(m.Exec) { // not required
		return nil
	}

	if m.Exec != nil {
		if err := m.Exec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) validateHTTPGet(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPGet) { // not required
		return nil
	}

	if m.HTTPGet != nil {
		if err := m.HTTPGet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) validateTCPSocket(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPSocket) { // not required
		return nil
	}

	if m.TCPSocket != nil {
		if err := m.TCPSocket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 lifecycle handler based on the context it is used
func (m *IoK8sAPICoreV1LifecycleHandler) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHTTPGet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPSocket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) contextValidateExec(ctx context.Context, formats strfmt.Registry) error {

	if m.Exec != nil {
		if err := m.Exec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) contextValidateHTTPGet(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPGet != nil {
		if err := m.HTTPGet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("httpGet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("httpGet")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1LifecycleHandler) contextValidateTCPSocket(ctx context.Context, formats strfmt.Registry) error {

	if m.TCPSocket != nil {
		if err := m.TCPSocket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tcpSocket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tcpSocket")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1LifecycleHandler) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1LifecycleHandler) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1LifecycleHandler
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}