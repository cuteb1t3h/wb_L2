package main

import (
	"github.com/beevik/ntp"
	"testing"
)

func TestNTPTime(t *testing.T) {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		t.Errorf("Error fetching NTP time: %v", err)
	}

	if ntpTime.IsZero() {
		t.Error("NTP time is zero")
	}
}

func TestNTPTimeError(t *testing.T) {
	_, err := ntp.Time("nonexistent.ntp.server")
	if err == nil {
		t.Error("Expected an error but got nil")
	}
}
