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

// IoK8sAPINodeV1RuntimeClass RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://kubernetes.io/docs/concepts/containers/runtime-class/
//
// swagger:model io.k8s.api.node.v1.RuntimeClass
type IoK8sAPINodeV1RuntimeClass struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Handler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The Handler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	// Required: true
	Handler *string `json:"handler"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see
	//  https://kubernetes.io/docs/concepts/scheduling-eviction/pod-overhead/
	// This field is in beta starting v1.18 and is only honored by servers that enable the PodOverhead feature.
	Overhead *IoK8sAPINodeV1Overhead `json:"overhead,omitempty"`

	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling *IoK8sAPINodeV1Scheduling `json:"scheduling,omitempty"`
}

// Validate validates this io k8s api node v1 runtime class
func (m *IoK8sAPINodeV1RuntimeClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHandler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverhead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduling(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1RuntimeClass) validateHandler(formats strfmt.Registry) error {

	if err := validate.Required("handler", "body", m.Handler); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPINodeV1RuntimeClass) validateMetadata(formats strfmt.Registry) error {
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

func (m *IoK8sAPINodeV1RuntimeClass) validateOverhead(formats strfmt.Registry) error {
	if swag.IsZero(m.Overhead) { // not required
		return nil
	}

	if m.Overhead != nil {
		if err := m.Overhead.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overhead")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overhead")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINodeV1RuntimeClass) validateScheduling(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheduling) { // not required
		return nil
	}

	if m.Scheduling != nil {
		if err := m.Scheduling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api node v1 runtime class based on the context it is used
func (m *IoK8sAPINodeV1RuntimeClass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOverhead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheduling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1RuntimeClass) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoK8sAPINodeV1RuntimeClass) contextValidateOverhead(ctx context.Context, formats strfmt.Registry) error {

	if m.Overhead != nil {
		if err := m.Overhead.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overhead")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overhead")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINodeV1RuntimeClass) contextValidateScheduling(ctx context.Context, formats strfmt.Registry) error {

	if m.Scheduling != nil {
		if err := m.Scheduling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINodeV1RuntimeClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINodeV1RuntimeClass) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINodeV1RuntimeClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
