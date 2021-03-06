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
	"github.com/go-openapi/validate"
)

// IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions RunAsGroupStrategyOptions defines the strategy type and any options used to create the strategy.
//
// swagger:model io.k8s.api.policy.v1beta1.RunAsGroupStrategyOptions
type IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions struct {

	// ranges are the allowed ranges of gids that may be used. If you would like to force a single gid then supply a single range with the same start and end. Required for MustRunAs.
	Ranges []*IoK8sAPIPolicyV1beta1IDRange `json:"ranges" json,yaml:"ranges"`

	// rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
	// Required: true
	Rule *string `json:"rule" json,yaml:"rule"`
}

// Validate validates this io k8s api policy v1beta1 run as group strategy options
func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) validateRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	for i := 0; i < len(m.Ranges); i++ {
		if swag.IsZero(m.Ranges[i]) { // not required
			continue
		}

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) validateRule(formats strfmt.Registry) error {

	if err := validate.Required("rule", "body", m.Rule); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api policy v1beta1 run as group strategy options based on the context it is used
func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ranges); i++ {

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1RunAsGroupStrategyOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
