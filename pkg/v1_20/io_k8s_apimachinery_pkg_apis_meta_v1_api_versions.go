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

// IoK8sApimachineryPkgApisMetaV1APIVersions APIVersions lists the versions that are available, to allow clients to discover the API at /api, which is the root path of the legacy v1 API.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.APIVersions
type IoK8sApimachineryPkgApisMetaV1APIVersions struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// a map of client CIDR to server address that is serving this group. This is to help clients reach servers in the most network-efficient way possible. Clients can use the appropriate server address as per the CIDR that they match. In case of multiple matches, clients should use the longest matching CIDR. The server returns only those CIDRs that it thinks that the client can match. For example: the master will return an internal IP CIDR only, if the client reaches the server using an internal IP. Server looks at X-Forwarded-For header or X-Real-Ip header or request.RemoteAddr (in that order) to get the client IP.
	// Required: true
	ServerAddressByClientCIDRs []*IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR `json:"serverAddressByClientCIDRs"`

	// versions are the api versions that are available.
	// Required: true
	Versions []string `json:"versions"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 API versions
func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerAddressByClientCIDRs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) validateServerAddressByClientCIDRs(formats strfmt.Registry) error {

	if err := validate.Required("serverAddressByClientCIDRs", "body", m.ServerAddressByClientCIDRs); err != nil {
		return err
	}

	for i := 0; i < len(m.ServerAddressByClientCIDRs); i++ {
		if swag.IsZero(m.ServerAddressByClientCIDRs[i]) { // not required
			continue
		}

		if m.ServerAddressByClientCIDRs[i] != nil {
			if err := m.ServerAddressByClientCIDRs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverAddressByClientCIDRs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverAddressByClientCIDRs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) validateVersions(formats strfmt.Registry) error {

	if err := validate.Required("versions", "body", m.Versions); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this io k8s apimachinery pkg apis meta v1 API versions based on the context it is used
func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerAddressByClientCIDRs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) contextValidateServerAddressByClientCIDRs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerAddressByClientCIDRs); i++ {

		if m.ServerAddressByClientCIDRs[i] != nil {
			if err := m.ServerAddressByClientCIDRs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverAddressByClientCIDRs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverAddressByClientCIDRs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1APIVersions) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1APIVersions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}