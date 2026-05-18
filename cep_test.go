package scalarsbr

import (
	"errors"
	"testing"
)

func TestValidateCEP(t *testing.T) {
	valid := []string{"01001000", "01001-000", "20040002", "30130-010"}
	for _, cep := range valid {
		if err := ValidateCEP(cep); err != nil {
			t.Fatalf("ValidateCEP(%q) returned %v", cep, err)
		}
	}

	invalid := []string{"", "0100-1000", "01001 000", "0100100", "010010000", "0100A000"}
	for _, cep := range invalid {
		if err := ValidateCEP(cep); !errors.Is(err, ErrInvalidCEP) {
			t.Fatalf("ValidateCEP(%q) returned %v", cep, err)
		}
	}
}

func TestFormatCEP(t *testing.T) {
	got, err := FormatCEP("CEP 01001-000")
	if err != nil {
		t.Fatalf("FormatCEP returned %v", err)
	}
	if got != "01001-000" {
		t.Fatalf("FormatCEP returned %q", got)
	}

	if _, err := FormatCEP("0100100"); !errors.Is(err, ErrInvalidCEP) {
		t.Fatalf("FormatCEP invalid returned %v", err)
	}
}
