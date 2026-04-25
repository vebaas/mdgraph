package registry

import "testing"

func TestRegisterAndExists(t *testing.T) {
	r := NewRegistry()
	r.Register("note")

	if !r.IsValid("note") {
		t.Error("expected note to be valid but it was not")
	}

	if r.IsValid("worker") {
		t.Error("expected worker to not be valid but it was")
	}
}
