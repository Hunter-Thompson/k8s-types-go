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

// IoK8sApimachineryPkgApisMetaV1WatchEvent Event represents a single event to a watched resource.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent
type IoK8sApimachineryPkgApisMetaV1WatchEvent struct {

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before deletion.
	//  * If Type is Error: *Status is recommended; other types may make sense
	//    depending on context.
	// Required: true
	Object IoK8sApimachineryPkgRuntimeRawExtension `json:"object" json,yaml:"object"`

	// type
	// Required: true
	Type *string `json:"type" json,yaml:"type"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 watch event
func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) validateObject(formats strfmt.Registry) error {

	if m.Object == nil {
		return errors.Required("object", "body", nil)
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 watch event based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1WatchEvent) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1WatchEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
