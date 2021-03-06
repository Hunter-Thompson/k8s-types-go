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

// IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus APIServiceStatus contains derived information about an API server
//
// swagger:model io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIServiceStatus
type IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus struct {

	// Current service state of apiService.
	Conditions []*IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition `json:"conditions" json,yaml:"conditions"`
}

// Validate validates this io k8s kube aggregator pkg apis apiregistration v1beta1 API service status
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s kube aggregator pkg apis apiregistration v1beta1 API service status based on the context it is used
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
