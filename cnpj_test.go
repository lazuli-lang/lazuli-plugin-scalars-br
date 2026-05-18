package scalarsbr

import (
	"errors"
	"testing"
)

func TestValidateCNPJValid(t *testing.T) {
	valid := []string{
		"11222333000181",
		"12.345.678/0001-95",
		"04252011000110",
		"40.688.134/0001-61",
		"71506168000111",
		"45.723.174/0001-10",
		"33214056000106",
		"27.865.757/0001-02",
		"60701190000104",
		"19.131.243/0001-97",
	}
	for _, cnpj := range valid {
		if err := ValidateCNPJ(cnpj); err != nil {
			t.Fatalf("ValidateCNPJ(%q) returned %v", cnpj, err)
		}
	}
}

func TestValidateCNPJInvalid(t *testing.T) {
	invalid := []string{
		"",
		"00000000000000",
		"11.111.111/1111-11",
		"11222333000180",
		"12345678000194",
		"04252011000111",
		"1122233300018",
		"112223330001811",
		"11.222.333/0001-8a",
		"40.688.134/0001-62",
	}
	for _, cnpj := range invalid {
		if err := ValidateCNPJ(cnpj); !errors.Is(err, ErrInvalidCNPJ) {
			t.Fatalf("ValidateCNPJ(%q) returned %v", cnpj, err)
		}
	}
}

func TestFormatCNPJ(t *testing.T) {
	got, err := FormatCNPJ("cnpj=11222333000181")
	if err != nil {
		t.Fatalf("FormatCNPJ returned %v", err)
	}
	if got != "11.222.333/0001-81" {
		t.Fatalf("FormatCNPJ returned %q", got)
	}

	if _, err := FormatCNPJ("00000000000000"); !errors.Is(err, ErrInvalidCNPJ) {
		t.Fatalf("FormatCNPJ invalid returned %v", err)
	}
}
