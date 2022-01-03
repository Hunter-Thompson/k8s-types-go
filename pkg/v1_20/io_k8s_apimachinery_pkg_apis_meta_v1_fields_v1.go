// Code generated by go-swagger; DO NOT EDIT.

package v1_20

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// IoK8sApimachineryPkgApisMetaV1FieldsV1 FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.
//
// Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.
//
// The exact format is defined in sigs.k8s.io/structured-merge-diff
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
type IoK8sApimachineryPkgApisMetaV1FieldsV1 interface{}
