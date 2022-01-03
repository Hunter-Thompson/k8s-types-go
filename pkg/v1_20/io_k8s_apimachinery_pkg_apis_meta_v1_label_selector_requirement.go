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

// IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
type IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement struct {

	// key is the label key that the selector applies to.
	// Required: true
	Key *string `json:"key" json,yaml:"key"`

	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	// Required: true
	Operator *string `json:"operator" json,yaml:"operator"`

	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values []string `json:"values" json,yaml:"values"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 label selector requirement
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 label selector requirement based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
