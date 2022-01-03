// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPINodeV1beta1Scheduling Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
//
// swagger:model io.k8s.api.node.v1beta1.Scheduling
type IoK8sAPINodeV1beta1Scheduling struct {

	// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	NodeSelector map[string]string `json:"nodeSelector,omitempty" json,yaml:"nodeSelector,omitempty"`

	// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
	Tolerations []*IoK8sAPICoreV1Toleration `json:"tolerations" json,yaml:"tolerations"`
}

// Validate validates this io k8s api node v1beta1 scheduling
func (m *IoK8sAPINodeV1beta1Scheduling) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTolerations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1beta1Scheduling) validateTolerations(formats strfmt.Registry) error {
	if swag.IsZero(m.Tolerations) { // not required
		return nil
	}

	for i := 0; i < len(m.Tolerations); i++ {
		if swag.IsZero(m.Tolerations[i]) { // not required
			continue
		}

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api node v1beta1 scheduling based on the context it is used
func (m *IoK8sAPINodeV1beta1Scheduling) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTolerations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1beta1Scheduling) contextValidateTolerations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tolerations); i++ {

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINodeV1beta1Scheduling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINodeV1beta1Scheduling) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINodeV1beta1Scheduling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
