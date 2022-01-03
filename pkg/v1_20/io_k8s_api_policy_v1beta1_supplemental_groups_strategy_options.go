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

// IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
//
// swagger:model io.k8s.api.policy.v1beta1.SupplementalGroupsStrategyOptions
type IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions struct {

	// ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. Required for MustRunAs.
	Ranges []*IoK8sAPIPolicyV1beta1IDRange `json:"ranges"`

	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	Rule string `json:"rule,omitempty"`
}

// Validate validates this io k8s api policy v1beta1 supplemental groups strategy options
func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) validateRanges(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api policy v1beta1 supplemental groups strategy options based on the context it is used
func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) contextValidateRanges(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1SupplementalGroupsStrategyOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
