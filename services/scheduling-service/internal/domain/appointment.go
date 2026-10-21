package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidAppointment = errors.New("invalid appointment")
	ErrInvalidTransition  = errors.New("invalid appointment status transition")
)

type AppointmentStatus string

const (
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCancelled AppointmentStatus = "cancelled"
	StatusCompleted AppointmentStatus = "completed"
	StatusNoShow    AppointmentStatus = "no_show"
)

type Appointment struct {
	ID         string
	TenantID   string
	BranchID   string
	ResourceID string
	CustomerID string
	StartsAt   time.Time
	EndsAt     time.Time
	Status     AppointmentStatus
}

func NewAppointment(id, tenantID, branchID, resourceID, customerID string, startsAt, endsAt time.Time) (Appointment, error) {
	if id == "" || tenantID == "" || branchID == "" || resourceID == "" || customerID == "" || !endsAt.After(startsAt) {
		return Appointment{}, ErrInvalidAppointment
	}

	return Appointment{
		ID: id, TenantID: tenantID, BranchID: branchID,
		ResourceID: resourceID, CustomerID: customerID,
		StartsAt: startsAt, EndsAt: endsAt, Status: StatusConfirmed,
	}, nil
}

func (a *Appointment) Cancel() error {
	if a.Status != StatusConfirmed {
		return ErrInvalidTransition
	}
	a.Status = StatusCancelled
	return nil
}

func (a *Appointment) Complete() error {
	if a.Status != StatusConfirmed {
		return ErrInvalidTransition
	}
	a.Status = StatusCompleted
	return nil
}

func (a *Appointment) MarkNoShow() error {
	if a.Status != StatusConfirmed {
		return ErrInvalidTransition
	}
	a.Status = StatusNoShow
	return nil
}
