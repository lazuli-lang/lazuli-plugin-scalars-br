package scalarsbr

import "errors"

var ErrInvalidCEP = errors.New("invalid CEP")

// ValidateCEP validates a Brazilian postal code in XXXXXXXX or XXXXX-XXX form.
func ValidateCEP(s string) error {
	if !validCEP(s) {
		return ErrInvalidCEP
	}
	return nil
}

// FormatCEP strips non-digits, validates the value, and returns XXXXX-XXX.
func FormatCEP(raw string) (string, error) {
	digits := stripDigits(raw)
	if len(digits) != 8 {
		return "", ErrInvalidCEP
	}
	return digits[:5] + "-" + digits[5:], nil
}

func validCEP(s string) bool {
	switch len(s) {
	case 8:
		return allDigits(s)
	case 9:
		return s[5] == '-' && allDigits(s[:5]) && allDigits(s[6:])
	default:
		return false
	}
}

func allDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}
