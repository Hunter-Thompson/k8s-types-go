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

// IoK8sAPIAutoscalingV2beta1ResourceMetricSource ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
//
// swagger:model io.k8s.api.autoscaling.v2beta1.ResourceMetricSource
type IoK8sAPIAutoscalingV2beta1ResourceMetricSource struct {

	// name is the name of the resource in question.
	// Required: true
	Name *string `json:"name" json,yaml:"name"`

	// targetAverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	TargetAverageUtilization int32 `json:"targetAverageUtilization,omitempty" json,yaml:"targetAverageUtilization,omitempty"`

	// targetAverageValue is the target value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type.
	TargetAverageValue IoK8sApimachineryPkgAPIResourceQuantity `json:"targetAverageValue,omitempty" json,yaml:"targetAverageValue,omitempty"`
}

// Validate validates this io k8s api autoscaling v2beta1 resource metric source
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) validateTargetAverageValue(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetAverageValue) { // not required
		return nil
	}

	if err := m.TargetAverageValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetAverageValue")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("targetAverageValue")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api autoscaling v2beta1 resource metric source based on the context it is used
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetAverageValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) contextValidateTargetAverageValue(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TargetAverageValue.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetAverageValue")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("targetAverageValue")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta1ResourceMetricSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
