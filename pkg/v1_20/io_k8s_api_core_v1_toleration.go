// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1Toleration The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.
//
// swagger:model io.k8s.api.core.v1.Toleration
type IoK8sAPICoreV1Toleration struct {

	// Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
	Effect string `json:"effect,omitempty" json,yaml:"effect,omitempty"`

	// Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
	Key string `json:"key,omitempty" json,yaml:"key,omitempty"`

	// Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
	Operator string `json:"operator,omitempty" json,yaml:"operator,omitempty"`

	// TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
	TolerationSeconds int64 `json:"tolerationSeconds,omitempty" json,yaml:"tolerationSeconds,omitempty"`

	// Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
	Value string `json:"value,omitempty" json,yaml:"value,omitempty"`
}

// Validate validates this io k8s api core v1 toleration
func (m *IoK8sAPICoreV1Toleration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s api core v1 toleration based on context it is used
func (m *IoK8sAPICoreV1Toleration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1Toleration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1Toleration) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1Toleration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
