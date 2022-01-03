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

// IoK8sAPINodeV1beta1Overhead Overhead structure represents the resource overhead associated with running a pod.
//
// swagger:model io.k8s.api.node.v1beta1.Overhead
type IoK8sAPINodeV1beta1Overhead struct {

	// PodFixed represents the fixed resource overhead associated with running a pod.
	PodFixed map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"podFixed,omitempty" json,yaml:"podFixed,omitempty"`
}

// Validate validates this io k8s api node v1beta1 overhead
func (m *IoK8sAPINodeV1beta1Overhead) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodFixed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1beta1Overhead) validatePodFixed(formats strfmt.Registry) error {
	if swag.IsZero(m.PodFixed) { // not required
		return nil
	}

	for k := range m.PodFixed {

		if val, ok := m.PodFixed[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api node v1beta1 overhead based on the context it is used
func (m *IoK8sAPINodeV1beta1Overhead) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePodFixed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1beta1Overhead) contextValidatePodFixed(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PodFixed {

		if val, ok := m.PodFixed[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINodeV1beta1Overhead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINodeV1beta1Overhead) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINodeV1beta1Overhead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
