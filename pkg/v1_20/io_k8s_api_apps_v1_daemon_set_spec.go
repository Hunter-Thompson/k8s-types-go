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

// IoK8sAPIAppsV1DaemonSetSpec DaemonSetSpec is the specification of a daemon set.
//
// swagger:model io.k8s.api.apps.v1.DaemonSetSpec
type IoK8sAPIAppsV1DaemonSetSpec struct {

	// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
	MinReadySeconds int32 `json:"minReadySeconds,omitempty"`

	// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
	RevisionHistoryLimit int32 `json:"revisionHistoryLimit,omitempty"`

	// A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	// Required: true
	Selector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"selector"`

	// An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
	// Required: true
	Template *IoK8sAPICoreV1PodTemplateSpec `json:"template"`

	// An update strategy to replace existing DaemonSet pods with new pods.
	UpdateStrategy *IoK8sAPIAppsV1DaemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
}

// Validate validates this io k8s api apps v1 daemon set spec
func (m *IoK8sAPIAppsV1DaemonSetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) validateSelector(formats strfmt.Registry) error {

	if err := validate.Required("selector", "body", m.Selector); err != nil {
		return err
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) validateUpdateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateStrategy) { // not required
		return nil
	}

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api apps v1 daemon set spec based on the context it is used
func (m *IoK8sAPIAppsV1DaemonSetSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.Selector != nil {
		if err := m.Selector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAppsV1DaemonSetSpec) contextValidateUpdateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdateStrategy != nil {
		if err := m.UpdateStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateStrategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateStrategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DaemonSetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DaemonSetSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1DaemonSetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
