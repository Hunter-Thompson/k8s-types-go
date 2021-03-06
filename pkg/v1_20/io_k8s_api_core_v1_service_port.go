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

// IoK8sAPICoreV1ServicePort ServicePort contains information on service's port.
//
// swagger:model io.k8s.api.core.v1.ServicePort
type IoK8sAPICoreV1ServicePort struct {

	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. This is a beta field that is guarded by the ServiceAppProtocol feature gate and enabled by default.
	AppProtocol string `json:"appProtocol,omitempty" json,yaml:"appProtocol,omitempty"`

	// The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.
	Name string `json:"name,omitempty" json,yaml:"name,omitempty"`

	// The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	NodePort int32 `json:"nodePort,omitempty" json,yaml:"nodePort,omitempty"`

	// The port that will be exposed by this service.
	// Required: true
	Port *int32 `json:"port" json,yaml:"port"`

	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default is TCP.
	Protocol string `json:"protocol,omitempty" json,yaml:"protocol,omitempty"`

	// Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	TargetPort IoK8sApimachineryPkgUtilIntstrIntOrString `json:"targetPort,omitempty" json,yaml:"targetPort,omitempty"`
}

// Validate validates this io k8s api core v1 service port
func (m *IoK8sAPICoreV1ServicePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ServicePort) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1ServicePort) validateTargetPort(formats strfmt.Registry) error {
	if swag.IsZero(m.TargetPort) { // not required
		return nil
	}

	if err := m.TargetPort.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetPort")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("targetPort")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api core v1 service port based on the context it is used
func (m *IoK8sAPICoreV1ServicePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTargetPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ServicePort) contextValidateTargetPort(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TargetPort.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetPort")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("targetPort")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ServicePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ServicePort) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ServicePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
