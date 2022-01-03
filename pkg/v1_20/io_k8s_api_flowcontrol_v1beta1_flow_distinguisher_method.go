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

// IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod FlowDistinguisherMethod specifies the method of a flow distinguisher.
//
// swagger:model io.k8s.api.flowcontrol.v1beta1.FlowDistinguisherMethod
type IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod struct {

	// `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api flowcontrol v1beta1 flow distinguisher method
func (m *IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api flowcontrol v1beta1 flow distinguisher method based on context it is used
func (m *IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1beta1FlowDistinguisherMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
