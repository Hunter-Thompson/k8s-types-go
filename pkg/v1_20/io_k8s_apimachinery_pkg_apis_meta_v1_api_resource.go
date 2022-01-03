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

// IoK8sApimachineryPkgApisMetaV1APIResource APIResource specifies the name of a resource and whether it is namespaced.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.APIResource
type IoK8sApimachineryPkgApisMetaV1APIResource struct {

	// categories is a list of the grouped resources this resource belongs to (e.g. 'all')
	Categories []string `json:"categories"`

	// group is the preferred group of the resource.  Empty implies the group of the containing resource list. For subresources, this may have a different value, for example: Scale".
	Group string `json:"group,omitempty"`

	// kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
	// Required: true
	Kind *string `json:"kind"`

	// name is the plural name of the resource.
	// Required: true
	Name *string `json:"name"`

	// namespaced indicates if a resource is namespaced or not.
	// Required: true
	Namespaced *bool `json:"namespaced"`

	// shortNames is a list of suggested short names of the resource.
	ShortNames []string `json:"shortNames"`

	// singularName is the singular name of the resource.  This allows clients to handle plural and singular opaquely. The singularName is more correct for reporting status on a single item and both singular and plural are allowed from the kubectl CLI interface.
	// Required: true
	SingularName *string `json:"singularName"`

	// The hash value of the storage version, the version this resource is converted to when written to the data store. Value must be treated as opaque by clients. Only equality comparison on the value is valid. This is an alpha feature and may change or be removed in the future. The field is populated by the apiserver only if the StorageVersionHash feature gate is enabled. This field will remain optional even if it graduates.
	StorageVersionHash string `json:"storageVersionHash,omitempty"`

	// verbs is a list of supported kube verbs (this includes get, list, watch, create, update, patch, delete, deletecollection, and proxy)
	// Required: true
	Verbs []string `json:"verbs"`

	// version is the preferred version of the resource.  Empty implies the version of the containing resource list For subresources, this may have a different value, for example: v1 (while inside a v1beta1 version of the core resource's group)".
	Version string `json:"version,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 API resource
func (m *IoK8sApimachineryPkgApisMetaV1APIResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaced(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingularName(formats); err != nil {
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

func (m *IoK8sApimachineryPkgApisMetaV1APIResource) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIResource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIResource) validateNamespaced(formats strfmt.Registry) error {

	if err := validate.Required("namespaced", "body", m.Namespaced); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIResource) validateSingularName(formats strfmt.Registry) error {

	if err := validate.Required("singularName", "body", m.SingularName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1APIResource) validateVerbs(formats strfmt.Registry) error {

	if err := validate.Required("verbs", "body", m.Verbs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 API resource based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1APIResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1APIResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1APIResource) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1APIResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
