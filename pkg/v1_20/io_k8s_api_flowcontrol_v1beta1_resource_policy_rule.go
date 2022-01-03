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

// IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule ResourcePolicyRule is a predicate that matches some resource requests, testing the request's verb and the target resource. A ResourcePolicyRule matches a resource request if and only if: (a) at least one member of verbs matches the request, (b) at least one member of apiGroups matches the request, (c) at least one member of resources matches the request, and (d) least one member of namespaces matches the request.
//
// swagger:model io.k8s.api.flowcontrol.v1beta1.ResourcePolicyRule
type IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule struct {

	// `apiGroups` is a list of matching API groups and may not be empty. "*" matches all API groups and, if present, must be the only entry. Required.
	// Required: true
	APIGroups []string `json:"apiGroups" json,yaml:"apiGroups"`

	// `clusterScope` indicates whether to match requests that do not specify a namespace (which happens either because the resource is not namespaced or the request targets all namespaces). If this field is omitted or false then the `namespaces` field must contain a non-empty list.
	ClusterScope bool `json:"clusterScope,omitempty" json,yaml:"clusterScope,omitempty"`

	// `namespaces` is a list of target namespaces that restricts matches.  A request that specifies a target namespace matches only if either (a) this list contains that target namespace or (b) this list contains "*".  Note that "*" matches any specified namespace but does not match a request that _does not specify_ a namespace (see the `clusterScope` field for that). This list may be empty, but only if `clusterScope` is true.
	Namespaces []string `json:"namespaces" json,yaml:"namespaces"`

	// `resources` is a list of matching resources (i.e., lowercase and plural) with, if desired, subresource.  For example, [ "services", "nodes/status" ].  This list may not be empty. "*" matches all resources and, if present, must be the only entry. Required.
	// Required: true
	Resources []string `json:"resources" json,yaml:"resources"`

	// `verbs` is a list of matching verbs and may not be empty. "*" matches all verbs and, if present, must be the only entry. Required.
	// Required: true
	Verbs []string `json:"verbs" json,yaml:"verbs"`
}

// Validate validates this io k8s api flowcontrol v1beta1 resource policy rule
func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerbs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) validateAPIGroups(formats strfmt.Registry) error {

	if err := validate.Required("apiGroups", "body", m.APIGroups); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) validateVerbs(formats strfmt.Registry) error {

	if err := validate.Required("verbs", "body", m.Verbs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s api flowcontrol v1beta1 resource policy rule based on context it is used
func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1beta1ResourcePolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
