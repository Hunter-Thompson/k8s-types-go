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

// IoK8sAPIAutoscalingV2MetricStatus MetricStatus describes the last-read state of a single metric.
//
// swagger:model io.k8s.api.autoscaling.v2.MetricStatus
type IoK8sAPIAutoscalingV2MetricStatus struct {

	// container resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	ContainerResource *IoK8sAPIAutoscalingV2ContainerResourceMetricStatus `json:"containerResource,omitempty"`

	// external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
	External *IoK8sAPIAutoscalingV2ExternalMetricStatus `json:"external,omitempty"`

	// object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
	Object *IoK8sAPIAutoscalingV2ObjectMetricStatus `json:"object,omitempty"`

	// pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
	Pods *IoK8sAPIAutoscalingV2PodsMetricStatus `json:"pods,omitempty"`

	// resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
	Resource *IoK8sAPIAutoscalingV2ResourceMetricStatus `json:"resource,omitempty"`

	// type is the type of metric source.  It will be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each corresponds to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api autoscaling v2 metric status
func (m *IoK8sAPIAutoscalingV2MetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validateContainerResource(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerResource) { // not required
		return nil
	}

	if m.ContainerResource != nil {
		if err := m.ContainerResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerResource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validateExternal(formats strfmt.Registry) error {
	if swag.IsZero(m.External) { // not required
		return nil
	}

	if m.External != nil {
		if err := m.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validateObject(formats strfmt.Registry) error {
	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validatePods(formats strfmt.Registry) error {
	if swag.IsZero(m.Pods) { // not required
		return nil
	}

	if m.Pods != nil {
		if err := m.Pods.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api autoscaling v2 metric status based on the context it is used
func (m *IoK8sAPIAutoscalingV2MetricStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) contextValidateContainerResource(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerResource != nil {
		if err := m.ContainerResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerResource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerResource")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) contextValidateExternal(ctx context.Context, formats strfmt.Registry) error {

	if m.External != nil {
		if err := m.External.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) contextValidateObject(ctx context.Context, formats strfmt.Registry) error {

	if m.Object != nil {
		if err := m.Object.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) contextValidatePods(ctx context.Context, formats strfmt.Registry) error {

	if m.Pods != nil {
		if err := m.Pods.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pods")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2MetricStatus) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2MetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2MetricStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2MetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
