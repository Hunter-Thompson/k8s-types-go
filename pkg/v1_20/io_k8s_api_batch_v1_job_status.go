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

// IoK8sAPIBatchV1JobStatus JobStatus represents the current state of a Job.
//
// swagger:model io.k8s.api.batch.v1.JobStatus
type IoK8sAPIBatchV1JobStatus struct {

	// The number of pending and running pods.
	Active int32 `json:"active,omitempty"`

	// CompletedIndexes holds the completed indexes when .spec.completionMode = "Indexed" in a text format. The indexes are represented as decimal integers separated by commas. The numbers are listed in increasing order. Three or more consecutive numbers are compressed and represented by the first and last element of the series, separated by a hyphen. For example, if the completed indexes are 1, 3, 4, 5 and 7, they are represented as "1,3-5,7".
	CompletedIndexes string `json:"completedIndexes,omitempty"`

	// Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC. The completion time is only set when the job finishes successfully.
	// Format: date-time
	CompletionTime IoK8sApimachineryPkgApisMetaV1Time `json:"completionTime,omitempty"`

	// The latest available observations of an object's current state. When a Job fails, one of the conditions will have type "Failed" and status true. When a Job is suspended, one of the conditions will have type "Suspended" and status true; when the Job is resumed, the status of this condition will become false. When a Job is completed, one of the conditions will have type "Complete" and status true. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Conditions []*IoK8sAPIBatchV1JobCondition `json:"conditions"`

	// The number of pods which reached phase Failed.
	Failed int32 `json:"failed,omitempty"`

	// The number of pods which have a Ready condition.
	//
	// This field is alpha-level. The job controller populates the field when the feature gate JobReadyPods is enabled (disabled by default).
	Ready int32 `json:"ready,omitempty"`

	// Represents time when the job controller started processing a job. When a Job is created in the suspended state, this field is not set until the first time it is resumed. This field is reset every time a Job is resumed from suspension. It is represented in RFC3339 form and is in UTC.
	// Format: date-time
	StartTime IoK8sApimachineryPkgApisMetaV1Time `json:"startTime,omitempty"`

	// The number of pods which reached phase Succeeded.
	Succeeded int32 `json:"succeeded,omitempty"`

	// UncountedTerminatedPods holds the UIDs of Pods that have terminated but the job controller hasn't yet accounted for in the status counters.
	//
	// The job controller creates pods with a finalizer. When a pod terminates (succeeded or failed), the controller does three steps to account for it in the job status: (1) Add the pod UID to the arrays in this field. (2) Remove the pod finalizer. (3) Remove the pod UID from the arrays while increasing the corresponding
	//     counter.
	//
	// This field is beta-level. The job controller only makes use of this field when the feature gate JobTrackingWithFinalizers is enabled (enabled by default). Old jobs might not be tracked using this field, in which case the field remains null.
	UncountedTerminatedPods *IoK8sAPIBatchV1UncountedTerminatedPods `json:"uncountedTerminatedPods,omitempty"`
}

// Validate validates this io k8s api batch v1 job status
func (m *IoK8sAPIBatchV1JobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUncountedTerminatedPods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := m.CompletionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("completionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("completionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) validateConditions(formats strfmt.Registry) error {
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

func (m *IoK8sAPIBatchV1JobStatus) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := m.StartTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("startTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) validateUncountedTerminatedPods(formats strfmt.Registry) error {
	if swag.IsZero(m.UncountedTerminatedPods) { // not required
		return nil
	}

	if m.UncountedTerminatedPods != nil {
		if err := m.UncountedTerminatedPods.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uncountedTerminatedPods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uncountedTerminatedPods")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this io k8s api batch v1 job status based on the context it is used
func (m *IoK8sAPIBatchV1JobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompletionTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUncountedTerminatedPods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) contextValidateCompletionTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CompletionTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("completionTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("completionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IoK8sAPIBatchV1JobStatus) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StartTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("startTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("startTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIBatchV1JobStatus) contextValidateUncountedTerminatedPods(ctx context.Context, formats strfmt.Registry) error {

	if m.UncountedTerminatedPods != nil {
		if err := m.UncountedTerminatedPods.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uncountedTerminatedPods")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uncountedTerminatedPods")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIBatchV1JobStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIBatchV1JobStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIBatchV1JobStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}