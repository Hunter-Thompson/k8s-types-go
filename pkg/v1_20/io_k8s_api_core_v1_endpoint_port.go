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

// IoK8sAPICoreV1EndpointPort EndpointPort is a tuple that describes a single port.
//
// swagger:model io.k8s.api.core.v1.EndpointPort
type IoK8sAPICoreV1EndpointPort struct {

	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. This is a beta field that is guarded by the ServiceAppProtocol feature gate and enabled by default.
	AppProtocol string `json:"appProtocol,omitempty" json,yaml:"appProtocol,omitempty"`

	// The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
	Name string `json:"name,omitempty" json,yaml:"name,omitempty"`

	// The port number of the endpoint.
	// Required: true
	Port *int32 `json:"port" json,yaml:"port"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol string `json:"protocol,omitempty" json,yaml:"protocol,omitempty"`
}

// Validate validates this io k8s api core v1 endpoint port
func (m *IoK8sAPICoreV1EndpointPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EndpointPort) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api core v1 endpoint port based on context it is used
func (m *IoK8sAPICoreV1EndpointPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EndpointPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EndpointPort) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EndpointPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
