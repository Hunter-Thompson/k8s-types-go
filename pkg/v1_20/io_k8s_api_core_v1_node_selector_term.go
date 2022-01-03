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

// IoK8sAPICoreV1NodeSelectorTerm A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
//
// swagger:model io.k8s.api.core.v1.NodeSelectorTerm
type IoK8sAPICoreV1NodeSelectorTerm struct {

	// A list of node selector requirements by node's labels.
	MatchExpressions []*IoK8sAPICoreV1NodeSelectorRequirement `json:"matchExpressions" json,yaml:"matchExpressions"`

	// A list of node selector requirements by node's fields.
	MatchFields []*IoK8sAPICoreV1NodeSelectorRequirement `json:"matchFields" json,yaml:"matchFields"`
}

// Validate validates this io k8s api core v1 node selector term
func (m *IoK8sAPICoreV1NodeSelectorTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchExpressions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeSelectorTerm) validateMatchExpressions(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchExpressions); i++ {
		if swag.IsZero(m.MatchExpressions[i]) { // not required
			continue
		}

		if m.MatchExpressions[i] != nil {
			if err := m.MatchExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSelectorTerm) validateMatchFields(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchFields) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchFields); i++ {
		if swag.IsZero(m.MatchFields[i]) { // not required
			continue
		}

		if m.MatchFields[i] != nil {
			if err := m.MatchFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api core v1 node selector term based on the context it is used
func (m *IoK8sAPICoreV1NodeSelectorTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchExpressions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatchFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeSelectorTerm) contextValidateMatchExpressions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchExpressions); i++ {

		if m.MatchExpressions[i] != nil {
			if err := m.MatchExpressions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSelectorTerm) contextValidateMatchFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MatchFields); i++ {

		if m.MatchFields[i] != nil {
			if err := m.MatchFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("matchFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSelectorTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSelectorTerm) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1NodeSelectorTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
