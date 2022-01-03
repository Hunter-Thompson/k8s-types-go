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

// IoK8sAPICoreV1Probe Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
//
// swagger:model io.k8s.api.core.v1.Probe
type IoK8sAPICoreV1Probe struct {

	// Exec specifies the action to take.
	Exec *IoK8sAPICoreV1ExecAction `json:"exec,omitempty"`

	// Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
	FailureThreshold int32 `json:"failureThreshold,omitempty"`

	// GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.
	Grpc *IoK8sAPICoreV1GRPCAction `json:"grpc,omitempty"`

	// HTTPGet specifies the http request to perform.
	HTTPGet *IoK8sAPICoreV1HTTPGetAction `json:"httpGet,omitempty"`

	// Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`

	// How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
	PeriodSeconds int32 `json:"periodSeconds,omitempty"`

	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold int32 `json:"successThreshold,omitempty"`

	// TCPSocket specifies an action involving a TCP port.
	TCPSocket *IoK8sAPICoreV1TCPSocketAction `json:"tcpSocket,omitempty"`

	// Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

// Validate validates this io k8s api core v1 probe
func (m *IoK8sAPICoreV1Probe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrpc(formats); err != nil {
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

func (m *IoK8sAPICoreV1Probe) validateExec(formats strfmt.Registry) error {
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

func (m *IoK8sAPICoreV1Probe) validateGrpc(formats strfmt.Registry) error {
	if swag.IsZero(m.Grpc) { // not required
		return nil
	}

	if m.Grpc != nil {
		if err := m.Grpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grpc")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Probe) validateHTTPGet(formats strfmt.Registry) error {
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

func (m *IoK8sAPICoreV1Probe) validateTCPSocket(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api core v1 probe based on the context it is used
func (m *IoK8sAPICoreV1Probe) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrpc(ctx, formats); err != nil {
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

func (m *IoK8sAPICoreV1Probe) contextValidateExec(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoK8sAPICoreV1Probe) contextValidateGrpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Grpc != nil {
		if err := m.Grpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grpc")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1Probe) contextValidateHTTPGet(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoK8sAPICoreV1Probe) contextValidateTCPSocket(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICoreV1Probe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1Probe) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1Probe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}