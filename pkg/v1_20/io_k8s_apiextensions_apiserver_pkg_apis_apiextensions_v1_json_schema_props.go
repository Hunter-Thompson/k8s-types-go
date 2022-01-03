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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.JSONSchemaProps
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps struct {

	// dollar ref
	DollarRef string `json:"$ref,omitempty"`

	// dollar schema
	DollarSchema string `json:"$schema,omitempty"`

	// additional items
	AdditionalItems IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsOrBool `json:"additionalItems,omitempty"`

	// additional properties
	AdditionalProperties IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsOrBool `json:"additionalProperties,omitempty"`

	// all of
	AllOf []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"allOf"`

	// any of
	AnyOf []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"anyOf"`

	// default is a default value for undefined object fields. Defaulting is a beta feature under the CustomResourceDefaulting feature gate. Defaulting requires spec.preserveUnknownFields to be false.
	Default IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSON `json:"default,omitempty"`

	// definitions
	Definitions map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"definitions,omitempty"`

	// dependencies
	Dependencies map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsOrStringArray `json:"dependencies,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enum
	Enum []IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSON `json:"enum"`

	// example
	Example IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSON `json:"example,omitempty"`

	// exclusive maximum
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`

	// exclusive minimum
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`

	// external docs
	ExternalDocs *IoK8sApiextensionsApiserverPkgApisApiextensionsV1ExternalDocumentation `json:"externalDocs,omitempty"`

	// format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated:
	//
	// - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like "0321751043" or "978-0321751041" - isbn10: an ISBN10 number string like "0321751043" - isbn13: an ISBN13 number string like "978-0321751041" - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^\d{3}[- ]?\d{2}[- ]?\d{4}$ - hexcolor: an hexadecimal color code like "#FFFFFF: following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like "rgb(255,255,2559" - byte: base64 encoded binary data - password: any kind of string - date: a date string like "2006-01-02" as defined by full-date in RFC3339 - duration: a duration string like "22 ns" as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like "2014-12-15T19:30:20.000Z" as defined by date-time in RFC3339.
	Format string `json:"format,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// items
	Items IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaPropsOrArray `json:"items,omitempty"`

	// max items
	MaxItems int64 `json:"maxItems,omitempty"`

	// max length
	MaxLength int64 `json:"maxLength,omitempty"`

	// max properties
	MaxProperties int64 `json:"maxProperties,omitempty"`

	// maximum
	Maximum float64 `json:"maximum,omitempty"`

	// min items
	MinItems int64 `json:"minItems,omitempty"`

	// min length
	MinLength int64 `json:"minLength,omitempty"`

	// min properties
	MinProperties int64 `json:"minProperties,omitempty"`

	// minimum
	Minimum float64 `json:"minimum,omitempty"`

	// multiple of
	MultipleOf float64 `json:"multipleOf,omitempty"`

	// not
	Not *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"not,omitempty"`

	// nullable
	Nullable bool `json:"nullable,omitempty"`

	// one of
	OneOf []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"oneOf"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// pattern properties
	PatternProperties map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"patternProperties,omitempty"`

	// properties
	Properties map[string]IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps `json:"properties,omitempty"`

	// required
	Required []string `json:"required"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unique items
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// x-kubernetes-embedded-resource defines that the value is an embedded Kubernetes runtime.Object, with TypeMeta and ObjectMeta. The type must be object. It is allowed to further restrict the embedded object. kind, apiVersion and metadata are validated automatically. x-kubernetes-preserve-unknown-fields is allowed to be true, but does not have to be if the object is fully specified (up to kind, apiVersion, metadata).
	XKubernetesEmbeddedResource bool `json:"x-kubernetes-embedded-resource,omitempty"`

	// x-kubernetes-int-or-string specifies that this value is either an integer or a string. If this is true, an empty type is allowed and type as child of anyOf is permitted if following one of the following patterns:
	//
	// 1) anyOf:
	//    - type: integer
	//    - type: string
	// 2) allOf:
	//    - anyOf:
	//      - type: integer
	//      - type: string
	//    - ... zero or more
	XKubernetesIntOrString bool `json:"x-kubernetes-int-or-string,omitempty"`

	// x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type `map` by specifying the keys used as the index of the map.
	//
	// This tag MUST only be used on lists that have the "x-kubernetes-list-type" extension set to "map". Also, the values specified for this attribute must be a scalar typed field of the child structure (no nesting is supported).
	//
	// The properties specified must either be required or have a default value, to ensure those properties are present for all list items.
	XKubernetesListMapKeys []string `json:"x-kubernetes-list-map-keys"`

	// x-kubernetes-list-type annotates an array to further describe its topology. This extension must only be used on lists and may have 3 possible values:
	//
	// 1) `atomic`: the list is treated as a single entity, like a scalar.
	//      Atomic lists will be entirely replaced when updated. This extension
	//      may be used on any type of list (struct, scalar, ...).
	// 2) `set`:
	//      Sets are lists that must not have multiple items with the same value. Each
	//      value must be a scalar, an object with x-kubernetes-map-type `atomic` or an
	//      array with x-kubernetes-list-type `atomic`.
	// 3) `map`:
	//      These lists are like maps in that their elements have a non-index key
	//      used to identify them. Order is preserved upon merge. The map tag
	//      must only be used on a list with elements of type object.
	// Defaults to atomic for arrays.
	XKubernetesListType string `json:"x-kubernetes-list-type,omitempty"`

	// x-kubernetes-map-type annotates an object to further describe its topology. This extension must only be used when type is object and may have 2 possible values:
	//
	// 1) `granular`:
	//      These maps are actual maps (key-value pairs) and each fields are independent
	//      from each other (they can each be manipulated by separate actors). This is
	//      the default behaviour for all maps.
	// 2) `atomic`: the list is treated as a single entity, like a scalar.
	//      Atomic maps will be entirely replaced when updated.
	XKubernetesMapType string `json:"x-kubernetes-map-type,omitempty"`

	// x-kubernetes-preserve-unknown-fields stops the API server decoding step from pruning fields which are not specified in the validation schema. This affects fields recursively, but switches back to normal pruning behaviour if nested properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden.
	XKubernetesPreserveUnknownFields bool `json:"x-kubernetes-preserve-unknown-fields,omitempty"`

	// x-kubernetes-validations describes a list of validation rules written in the CEL expression language. This field is an alpha-level. Using this field requires the feature gate `CustomResourceValidationExpressions` to be enabled.
	XKubernetesValidations []*IoK8sApiextensionsApiserverPkgApisApiextensionsV1ValidationRule `json:"x-kubernetes-validations"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1 JSON schema props
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnyOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDocs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatternProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXKubernetesValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateAllOf(formats strfmt.Registry) error {
	if swag.IsZero(m.AllOf) { // not required
		return nil
	}

	for i := 0; i < len(m.AllOf); i++ {
		if swag.IsZero(m.AllOf[i]) { // not required
			continue
		}

		if m.AllOf[i] != nil {
			if err := m.AllOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateAnyOf(formats strfmt.Registry) error {
	if swag.IsZero(m.AnyOf) { // not required
		return nil
	}

	for i := 0; i < len(m.AnyOf); i++ {
		if swag.IsZero(m.AnyOf[i]) { // not required
			continue
		}

		if m.AnyOf[i] != nil {
			if err := m.AnyOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("anyOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("anyOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateDefinitions(formats strfmt.Registry) error {
	if swag.IsZero(m.Definitions) { // not required
		return nil
	}

	for k := range m.Definitions {

		if err := validate.Required("definitions"+"."+k, "body", m.Definitions[k]); err != nil {
			return err
		}
		if val, ok := m.Definitions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateExternalDocs(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalDocs) { // not required
		return nil
	}

	if m.ExternalDocs != nil {
		if err := m.ExternalDocs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDocs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalDocs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateNot(formats strfmt.Registry) error {
	if swag.IsZero(m.Not) { // not required
		return nil
	}

	if m.Not != nil {
		if err := m.Not.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("not")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateOneOf(formats strfmt.Registry) error {
	if swag.IsZero(m.OneOf) { // not required
		return nil
	}

	for i := 0; i < len(m.OneOf); i++ {
		if swag.IsZero(m.OneOf[i]) { // not required
			continue
		}

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validatePatternProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.PatternProperties) { // not required
		return nil
	}

	for k := range m.PatternProperties {

		if err := validate.Required("patternProperties"+"."+k, "body", m.PatternProperties[k]); err != nil {
			return err
		}
		if val, ok := m.PatternProperties[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}
		if val, ok := m.Properties[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) validateXKubernetesValidations(formats strfmt.Registry) error {
	if swag.IsZero(m.XKubernetesValidations) { // not required
		return nil
	}

	for i := 0; i < len(m.XKubernetesValidations); i++ {
		if swag.IsZero(m.XKubernetesValidations[i]) { // not required
			continue
		}

		if m.XKubernetesValidations[i] != nil {
			if err := m.XKubernetesValidations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("x-kubernetes-validations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("x-kubernetes-validations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this io k8s apiextensions apiserver pkg apis apiextensions v1 JSON schema props based on the context it is used
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnyOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefinitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalDocs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePatternProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateXKubernetesValidations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateAllOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllOf); i++ {

		if m.AllOf[i] != nil {
			if err := m.AllOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateAnyOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AnyOf); i++ {

		if m.AnyOf[i] != nil {
			if err := m.AnyOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("anyOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("anyOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateDefinitions(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Definitions {

		if val, ok := m.Definitions[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateExternalDocs(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalDocs != nil {
		if err := m.ExternalDocs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDocs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalDocs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateNot(ctx context.Context, formats strfmt.Registry) error {

	if m.Not != nil {
		if err := m.Not.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("not")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateOneOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OneOf); i++ {

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidatePatternProperties(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.PatternProperties {

		if val, ok := m.PatternProperties[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Properties {

		if val, ok := m.Properties[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) contextValidateXKubernetesValidations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.XKubernetesValidations); i++ {

		if m.XKubernetesValidations[i] != nil {
			if err := m.XKubernetesValidations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("x-kubernetes-validations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("x-kubernetes-validations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1JSONSchemaProps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
