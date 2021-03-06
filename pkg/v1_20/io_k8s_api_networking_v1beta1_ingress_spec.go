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

// IoK8sAPINetworkingV1beta1IngressSpec IngressSpec describes the Ingress the user wishes to exist.
//
// swagger:model io.k8s.api.networking.v1beta1.IngressSpec
type IoK8sAPINetworkingV1beta1IngressSpec struct {

	// A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	Backend *IoK8sAPINetworkingV1beta1IngressBackend `json:"backend,omitempty" json,yaml:"backend,omitempty"`

	// IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated `kubernetes.io/ingress.class` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.
	IngressClassName string `json:"ingressClassName,omitempty" json,yaml:"ingressClassName,omitempty"`

	// A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.
	Rules []*IoK8sAPINetworkingV1beta1IngressRule `json:"rules" json,yaml:"rules"`

	// TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	TLS []*IoK8sAPINetworkingV1beta1IngressTLS `json:"tls" json,yaml:"tls"`
}

// Validate validates this io k8s api networking v1beta1 ingress spec
func (m *IoK8sAPINetworkingV1beta1IngressSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) validateBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.Backend) { // not required
		return nil
	}

	if m.Backend != nil {
		if err := m.Backend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	for i := 0; i < len(m.TLS); i++ {
		if swag.IsZero(m.TLS[i]) { // not required
			continue
		}

		if m.TLS[i] != nil {
			if err := m.TLS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api networking v1beta1 ingress spec based on the context it is used
func (m *IoK8sAPINetworkingV1beta1IngressSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) contextValidateBackend(ctx context.Context, formats strfmt.Registry) error {

	if m.Backend != nil {
		if err := m.Backend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backend")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rules); i++ {

		if m.Rules[i] != nil {
			if err := m.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPINetworkingV1beta1IngressSpec) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TLS); i++ {

		if m.TLS[i] != nil {
			if err := m.TLS[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1beta1IngressSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1beta1IngressSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
