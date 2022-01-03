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

// IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject ServiceAccountSubject holds detailed information for service-account-kind subject.
//
// swagger:model io.k8s.api.flowcontrol.v1beta1.ServiceAccountSubject
type IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject struct {

	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	// Required: true
	Name *string `json:"name" json,yaml:"name"`

	// `namespace` is the namespace of matching ServiceAccount objects. Required.
	// Required: true
	Namespace *string `json:"namespace" json,yaml:"namespace"`
}

// Validate validates this io k8s api flowcontrol v1beta1 service account subject
func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api flowcontrol v1beta1 service account subject based on context it is used
func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1beta1ServiceAccountSubject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
