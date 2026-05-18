package scalarsbr

import (
	"errors"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	valid := []string{
		"11999990000",
		"(11) 99999-0000",
		"21 3333-4444",
		"(85) 98888-7777",
		"97912345678",
	}
	for _, phone := range valid {
		if err := ValidatePhone(phone); err != nil {
			t.Fatalf("ValidatePhone(%q) returned %v", phone, err)
		}
	}

	invalid := []string{
		"",
		"20999990000",
		"119999000",
		"119999900000",
		"2312345678",
		"00999990000",
	}
	for _, phone := range invalid {
		if err := ValidatePhone(phone); !errors.Is(err, ErrInvalidPhone) {
			t.Fatalf("ValidatePhone(%q) returned %v", phone, err)
		}
	}
}

func TestFormatPhone(t *testing.T) {
	tests := map[string]string{
		"11999990000":    "(11) 99999-0000",
		"(11) 3333-4444": "(11) 3333-4444",
		"99999-0000":     "(11) 99999-0000",
		"3333-4444":      "(11) 3333-4444",
		"2199998888":     "(21) 9999-8888",
	}
	for raw, want := range tests {
		got, err := FormatPhone(raw, "11")
		if err != nil {
			t.Fatalf("FormatPhone(%q) returned %v", raw, err)
		}
		if got != want {
			t.Fatalf("FormatPhone(%q) returned %q, want %q", raw, got, want)
		}
	}

	if _, err := FormatPhone("99999-0000", "20"); !errors.Is(err, ErrInvalidPhone) {
		t.Fatalf("FormatPhone invalid default returned %v", err)
	}
}
