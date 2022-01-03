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

// IoK8sAPIAutoscalingV2beta2ResourceMetricStatus ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
//
// swagger:model io.k8s.api.autoscaling.v2beta2.ResourceMetricStatus
type IoK8sAPIAutoscalingV2beta2ResourceMetricStatus struct {

	// current contains the current value for the given metric
	// Required: true
	Current *IoK8sAPIAutoscalingV2beta2MetricValueStatus `json:"current"`

	// Name is the name of the resource in question.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this io k8s api autoscaling v2beta2 resource metric status
func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api autoscaling v2beta2 resource metric status based on the context it is used
func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) contextValidateCurrent(ctx context.Context, formats strfmt.Registry) error {

	if m.Current != nil {
		if err := m.Current.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ResourceMetricStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2ResourceMetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
