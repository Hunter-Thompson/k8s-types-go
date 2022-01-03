// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1ContainerStateWaiting ContainerStateWaiting is a waiting state of a container.
//
// swagger:model io.k8s.api.core.v1.ContainerStateWaiting
type IoK8sAPICoreV1ContainerStateWaiting struct {

	// Message regarding why the container is not yet running.
	Message string `json:"message,omitempty"`

	// (brief) reason the container is not yet running.
	Reason string `json:"reason,omitempty"`
}

// Validate validates this io k8s api core v1 container state waiting
func (m *IoK8sAPICoreV1ContainerStateWaiting) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 container state waiting based on context it is used
func (m *IoK8sAPICoreV1ContainerStateWaiting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerStateWaiting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerStateWaiting) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ContainerStateWaiting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
