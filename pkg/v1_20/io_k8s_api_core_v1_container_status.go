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

// IoK8sAPICoreV1ContainerStatus ContainerStatus contains details for the current status of this container.
//
// swagger:model io.k8s.api.core.v1.ContainerStatus
type IoK8sAPICoreV1ContainerStatus struct {

	// Container's ID in the format 'docker://<container_id>'.
	ContainerID string `json:"containerID,omitempty" json,yaml:"containerID,omitempty"`

	// The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images
	// Required: true
	Image *string `json:"image" json,yaml:"image"`

	// ImageID of the container's image.
	// Required: true
	ImageID *string `json:"imageID" json,yaml:"imageID"`

	// Details about the container's last termination condition.
	LastState *IoK8sAPICoreV1ContainerState `json:"lastState,omitempty" json,yaml:"lastState,omitempty"`

	// This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
	// Required: true
	Name *string `json:"name" json,yaml:"name"`

	// Specifies whether the container has passed its readiness probe.
	// Required: true
	Ready *bool `json:"ready" json,yaml:"ready"`

	// The number of times the container has been restarted, currently based on the number of dead containers that have not yet been removed. Note that this is calculated from dead containers. But those containers are subject to garbage collection. This value will get capped at 5 by GC.
	// Required: true
	RestartCount *int32 `json:"restartCount" json,yaml:"restartCount"`

	// Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always true when no startupProbe is defined.
	Started bool `json:"started,omitempty" json,yaml:"started,omitempty"`

	// Details about the container's current condition.
	State *IoK8sAPICoreV1ContainerState `json:"state,omitempty" json,yaml:"state,omitempty"`
}

// Validate validates this io k8s api core v1 container status
func (m *IoK8sAPICoreV1ContainerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReady(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestartCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateLastState(formats strfmt.Registry) error {
	if swag.IsZero(m.LastState) { // not required
		return nil
	}

	if m.LastState != nil {
		if err := m.LastState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastState")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateReady(formats strfmt.Registry) error {

	if err := validate.Required("ready", "body", m.Ready); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateRestartCount(formats strfmt.Registry) error {

	if err := validate.Required("restartCount", "body", m.RestartCount); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 container status based on the context it is used
func (m *IoK8sAPICoreV1ContainerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) contextValidateLastState(ctx context.Context, formats strfmt.Registry) error {

	if m.LastState != nil {
		if err := m.LastState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastState")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerStatus) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ContainerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
