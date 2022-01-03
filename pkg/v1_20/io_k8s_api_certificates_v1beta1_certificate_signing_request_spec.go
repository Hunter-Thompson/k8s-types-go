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

// IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
//
// swagger:model io.k8s.api.certificates.v1beta1.CertificateSigningRequestSpec
type IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec struct {

	// Extra information about the requesting user. See user.Info interface for details.
	Extra map[string][]string `json:"extra,omitempty" json,yaml:"extra,omitempty"`

	// Group information about the requesting user. See user.Info interface for details.
	Groups []string `json:"groups" json,yaml:"groups"`

	// Base64-encoded PKCS#10 CSR data
	// Required: true
	// Format: byte
	Request *strfmt.Base64 `json:"request" json,yaml:"request"`

	// Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
	//  1. If it's a kubelet client certificate, it is assigned
	//     "kubernetes.io/kube-apiserver-client-kubelet".
	//  2. If it's a kubelet serving certificate, it is assigned
	//     "kubernetes.io/kubelet-serving".
	//  3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
	// Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
	SignerName string `json:"signerName,omitempty" json,yaml:"signerName,omitempty"`

	// UID information about the requesting user. See user.Info interface for details.
	UID string `json:"uid,omitempty" json,yaml:"uid,omitempty"`

	// allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
	//      https://tools.ietf.org/html/rfc5280#section-4.2.1.12
	// Valid values are:
	//  "signing",
	//  "digital signature",
	//  "content commitment",
	//  "key encipherment",
	//  "key agreement",
	//  "data encipherment",
	//  "cert sign",
	//  "crl sign",
	//  "encipher only",
	//  "decipher only",
	//  "any",
	//  "server auth",
	//  "client auth",
	//  "code signing",
	//  "email protection",
	//  "s/mime",
	//  "ipsec end system",
	//  "ipsec tunnel",
	//  "ipsec user",
	//  "timestamping",
	//  "ocsp signing",
	//  "microsoft sgc",
	//  "netscape sgc"
	Usages []string `json:"usages" json,yaml:"usages"`

	// Information about the requesting user. See user.Info interface for details.
	Username string `json:"username,omitempty" json,yaml:"username,omitempty"`
}

// Validate validates this io k8s api certificates v1beta1 certificate signing request spec
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec) validateRequest(formats strfmt.Registry) error {

	if err := validate.Required("request", "body", m.Request); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api certificates v1beta1 certificate signing request spec based on context it is used
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1beta1CertificateSigningRequestSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}