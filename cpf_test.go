package scalarsbr

import (
	"errors"
	"testing"
)

func TestValidateCPFValid(t *testing.T) {
	valid := []string{
		"52998224725",
		"111.444.777-35",
		"12345678909",
		"987.654.321-00",
		"93541134780",
		"390.533.447-05",
		"16899535009",
		"033.000.000-41",
		"04397328080",
		"652.509.920-05",
	}
	for _, cpf := range valid {
		if err := ValidateCPF(cpf); err != nil {
			t.Fatalf("ValidateCPF(%q) returned %v", cpf, err)
		}
	}
}

func TestValidateCPFInvalid(t *testing.T) {
	invalid := []string{
		"",
		"00000000000",
		"111.111.111-11",
		"12345678900",
		"52998224724",
		"93541134781",
		"1234567890",
		"123456789091",
		"123.456.789-0a",
		"987.654.321-01",
	}
	for _, cpf := range invalid {
		if err := ValidateCPF(cpf); !errors.Is(err, ErrInvalidCPF) {
			t.Fatalf("ValidateCPF(%q) returned %v", cpf, err)
		}
	}
}

func TestFormatCPF(t *testing.T) {
	got, err := FormatCPF("cpf=529.982.247-25")
	if err != nil {
		t.Fatalf("FormatCPF returned %v", err)
	}
	if got != "529.982.247-25" {
		t.Fatalf("FormatCPF returned %q", got)
	}

	if _, err := FormatCPF("11111111111"); !errors.Is(err, ErrInvalidCPF) {
		t.Fatalf("FormatCPF invalid returned %v", err)
	}
}
