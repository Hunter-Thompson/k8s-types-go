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

// IoK8sApimachineryPkgApisMetaV1LabelSelector A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
type IoK8sApimachineryPkgApisMetaV1LabelSelector struct {

	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions []*IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement `json:"matchExpressions"`

	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 label selector
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchExpressions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) validateMatchExpressions(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s apimachinery pkg apis meta v1 label selector based on the context it is used
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchExpressions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) contextValidateMatchExpressions(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelector) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1LabelSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}