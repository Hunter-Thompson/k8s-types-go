// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIDiscoveryV1beta1EndpointPort EndpointPort represents a Port used by an EndpointSlice
//
// swagger:model io.k8s.api.discovery.v1beta1.EndpointPort
type IoK8sAPIDiscoveryV1beta1EndpointPort struct {

	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
	AppProtocol string `json:"appProtocol,omitempty" json,yaml:"appProtocol,omitempty"`

	// The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
	Name string `json:"name,omitempty" json,yaml:"name,omitempty"`

	// The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
	Port int32 `json:"port,omitempty" json,yaml:"port,omitempty"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol string `json:"protocol,omitempty" json,yaml:"protocol,omitempty"`
}

// Validate validates this io k8s api discovery v1beta1 endpoint port
func (m *IoK8sAPIDiscoveryV1beta1EndpointPort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api discovery v1beta1 endpoint port based on context it is used
func (m *IoK8sAPIDiscoveryV1beta1EndpointPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1EndpointPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1EndpointPort) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIDiscoveryV1beta1EndpointPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
