package domain

import (
	"errors"
	"testing"
	"time"
)

func TestNewAppointmentRejectsNonPositiveDuration(t *testing.T) {
	now := time.Date(2026, 10, 20, 10, 0, 0, 0, time.UTC)
	_, err := NewAppointment("a", "t", "b", "r", "c", now, now)
	if !errors.Is(err, ErrInvalidAppointment) {
		t.Fatalf("expected invalid appointment error, got %v", err)
	}
}

func TestAppointmentLifecycleIsForwardOnly(t *testing.T) {
	now := time.Date(2026, 10, 20, 10, 0, 0, 0, time.UTC)
	a, err := NewAppointment("a", "t", "b", "r", "c", now, now.Add(time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	if err := a.Cancel(); err != nil {
		t.Fatal(err)
	}
	if err := a.Complete(); !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("expected invalid transition, got %v", err)
	}
}
