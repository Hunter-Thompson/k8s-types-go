// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAutoscalingV1ScaleSpec ScaleSpec describes the attributes of a scale subresource.
//
// swagger:model io.k8s.api.autoscaling.v1.ScaleSpec
type IoK8sAPIAutoscalingV1ScaleSpec struct {

	// desired number of instances for the scaled object.
	Replicas int32 `json:"replicas,omitempty" json,yaml:"replicas,omitempty"`
}

// Validate validates this io k8s api autoscaling v1 scale spec
func (m *IoK8sAPIAutoscalingV1ScaleSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api autoscaling v1 scale spec based on context it is used
func (m *IoK8sAPIAutoscalingV1ScaleSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV1ScaleSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV1ScaleSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV1ScaleSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
