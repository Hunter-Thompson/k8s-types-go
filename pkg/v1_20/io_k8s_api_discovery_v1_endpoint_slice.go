// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPIDiscoveryV1EndpointSlice EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
//
// swagger:model io.k8s.api.discovery.v1.EndpointSlice
type IoK8sAPIDiscoveryV1EndpointSlice struct {

	// addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
	//
	// Possible enum values:
	//  - `"FQDN"` represents a FQDN.
	//  - `"IPv4"` represents an IPv4 Address.
	//  - `"IPv6"` represents an IPv6 Address.
	// Required: true
	// Enum: [FQDN IPv4 IPv6]
	AddressType *string `json:"addressType"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
	// Required: true
	Endpoints []*IoK8sAPIDiscoveryV1Endpoint `json:"endpoints"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object's metadata.
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
	Ports []*IoK8sAPIDiscoveryV1EndpointPort `json:"ports"`
}

// Validate validates this io k8s api discovery v1 endpoint slice
func (m *IoK8sAPIDiscoveryV1EndpointSlice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ioK8sApiDiscoveryV1EndpointSliceTypeAddressTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FQDN","IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioK8sApiDiscoveryV1EndpointSliceTypeAddressTypePropEnum = append(ioK8sApiDiscoveryV1EndpointSliceTypeAddressTypePropEnum, v)
	}
}

const (

	// IoK8sAPIDiscoveryV1EndpointSliceAddressTypeFQDN captures enum value "FQDN"
	IoK8sAPIDiscoveryV1EndpointSliceAddressTypeFQDN string = "FQDN"

	// IoK8sAPIDiscoveryV1EndpointSliceAddressTypeIPV4 captures enum value "IPv4"
	IoK8sAPIDiscoveryV1EndpointSliceAddressTypeIPV4 string = "IPv4"

	// IoK8sAPIDiscoveryV1EndpointSliceAddressTypeIPV6 captures enum value "IPv6"
	IoK8sAPIDiscoveryV1EndpointSliceAddressTypeIPV6 string = "IPv6"
)

// prop value enum
func (m *IoK8sAPIDiscoveryV1EndpointSlice) validateAddressTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ioK8sApiDiscoveryV1EndpointSliceTypeAddressTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) validateAddressType(formats strfmt.Registry) error {

	if err := validate.Required("addressType", "body", m.AddressType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAddressTypeEnum("addressType", "body", *m.AddressType); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) validateEndpoints(formats strfmt.Registry) error {

	if err := validate.Required("endpoints", "body", m.Endpoints); err != nil {
		return err
	}

	for i := 0; i < len(m.Endpoints); i++ {
		if swag.IsZero(m.Endpoints[i]) { // not required
			continue
		}

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s api discovery v1 endpoint slice based on the context it is used
func (m *IoK8sAPIDiscoveryV1EndpointSlice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Endpoints); i++ {

		if m.Endpoints[i] != nil {
			if err := m.Endpoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endpoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIDiscoveryV1EndpointSlice) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1EndpointSlice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1EndpointSlice) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIDiscoveryV1EndpointSlice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
