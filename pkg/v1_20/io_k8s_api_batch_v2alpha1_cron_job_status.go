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

// IoK8sAPIBatchV2alpha1CronJobStatus CronJobStatus represents the current state of a cron job.
//
// swagger:model io.k8s.api.batch.v2alpha1.CronJobStatus
type IoK8sAPIBatchV2alpha1CronJobStatus struct {

	// A list of pointers to currently running jobs.
	Active []*IoK8sAPICoreV1ObjectReference `json:"active" json,yaml:"active"`

	// Information when was the last time the job was successfully scheduled.
	// Format: date-time
	LastScheduleTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastScheduleTime,omitempty" json,yaml:"lastScheduleTime,omitempty"`
}

// Validate validates this io k8s api batch v2alpha1 cron job status
func (m *IoK8sAPIBatchV2alpha1CronJobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScheduleTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobStatus) validateActive(formats strfmt.Registry) error {
	if swag.IsZero(m.Active) { // not required
		return nil
	}

	for i := 0; i < len(m.Active); i++ {
		if swag.IsZero(m.Active[i]) { // not required
			continue
		}

		if m.Active[i] != nil {
			if err := m.Active[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobStatus) validateLastScheduleTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastScheduleTime) { // not required
		return nil
	}

	if err := m.LastScheduleTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastScheduleTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastScheduleTime")
		}
		return err
	}

	return nil
}

// ContextValidate validate this io k8s api batch v2alpha1 cron job status based on the context it is used
func (m *IoK8sAPIBatchV2alpha1CronJobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastScheduleTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobStatus) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Active); i++ {

		if m.Active[i] != nil {
			if err := m.Active[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIBatchV2alpha1CronJobStatus) contextValidateLastScheduleTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastScheduleTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastScheduleTime")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("lastScheduleTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIBatchV2alpha1CronJobStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIBatchV2alpha1CronJobStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIBatchV2alpha1CronJobStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
