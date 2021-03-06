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

// IoK8sAPICertificatesV1CertificateSigningRequestStatus CertificateSigningRequestStatus contains conditions used to indicate approved/denied/failed status of the request, and the issued certificate.
//
// swagger:model io.k8s.api.certificates.v1.CertificateSigningRequestStatus
type IoK8sAPICertificatesV1CertificateSigningRequestStatus struct {

	// certificate is populated with an issued certificate by the signer after an Approved condition is present. This field is set via the /status subresource. Once populated, this field is immutable.
	//
	// If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.
	//
	// Validation requirements:
	//  1. certificate must contain one or more PEM blocks.
	//  2. All PEM blocks must have the "CERTIFICATE" label, contain no headers, and the encoded data
	//   must be a BER-encoded ASN.1 Certificate structure as described in section 4 of RFC5280.
	//  3. Non-PEM content may appear before or after the "CERTIFICATE" PEM blocks and is unvalidated,
	//   to allow for explanatory text as described in section 5.2 of RFC7468.
	//
	// If more than one PEM block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.
	//
	// The certificate is encoded in PEM format.
	//
	// When serialized as JSON or YAML, the data is additionally base64-encoded, so it consists of:
	//
	//     base64(
	//     -----BEGIN CERTIFICATE-----
	//     ...
	//     -----END CERTIFICATE-----
	//     )
	// Format: byte
	Certificate strfmt.Base64 `json:"certificate,omitempty" json,yaml:"certificate,omitempty"`

	// conditions applied to the request. Known conditions are "Approved", "Denied", and "Failed".
	Conditions []*IoK8sAPICertificatesV1CertificateSigningRequestCondition `json:"conditions" json,yaml:"conditions"`
}

// Validate validates this io k8s api certificates v1 certificate signing request status
func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) validateConditions(formats strfmt.Registry) error {
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

// ContextValidate validate this io k8s api certificates v1 certificate signing request status based on the context it is used
func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1CertificateSigningRequestStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1CertificateSigningRequestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
