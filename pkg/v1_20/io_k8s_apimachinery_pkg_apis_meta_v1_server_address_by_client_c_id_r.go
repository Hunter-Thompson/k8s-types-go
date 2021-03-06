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

// IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.ServerAddressByClientCIDR
type IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR struct {

	// The CIDR with which clients can match their IP to figure out the server address that they should use.
	// Required: true
	ClientCIDR *string `json:"clientCIDR" json,yaml:"clientCIDR"`

	// Address of this server, suitable for a client that matches the above CIDR. This can be a hostname, hostname:port, IP or IP:port.
	// Required: true
	ServerAddress *string `json:"serverAddress" json,yaml:"serverAddress"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 server address by client c ID r
func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientCIDR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) validateClientCIDR(formats strfmt.Registry) error {

	if err := validate.Required("clientCIDR", "body", m.ClientCIDR); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) validateServerAddress(formats strfmt.Registry) error {

	if err := validate.Required("serverAddress", "body", m.ServerAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 server address by client c ID r based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
