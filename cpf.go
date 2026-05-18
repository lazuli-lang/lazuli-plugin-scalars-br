// Package scalarsbr validates and formats Brazilian scalar values.
package scalarsbr

import "errors"

var ErrInvalidCPF = errors.New("invalid CPF")

// ValidateCPF validates an 11-digit Brazilian individual taxpayer number.
func ValidateCPF(s string) error {
	digits, ok := collectDigits(s, ".- ")
	if !ok || !validCPFDigits(digits) {
		return ErrInvalidCPF
	}
	return nil
}

// FormatCPF strips non-digits, validates the value, and returns XXX.XXX.XXX-XX.
func FormatCPF(raw string) (string, error) {
	digits := stripDigits(raw)
	if !validCPFDigits(digits) {
		return "", ErrInvalidCPF
	}
	return digits[:3] + "." + digits[3:6] + "." + digits[6:9] + "-" + digits[9:], nil
}

func validCPFDigits(digits string) bool {
	if len(digits) != 11 || allSameDigits(digits) {
		return false
	}
	return int(digits[9]-'0') == cpfDigit(digits[:9], 10) &&
		int(digits[10]-'0') == cpfDigit(digits[:10], 11)
}

func cpfDigit(digits string, weight int) int {
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i]-'0') * (weight - i)
	}
	digit := 11 - (sum % 11)
	if digit >= 10 {
		return 0
	}
	return digit
}

func collectDigits(s, allowed string) (string, bool) {
	digits := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			digits = append(digits, c)
		case containsByte(allowed, c):
		default:
			return "", false
		}
	}
	return string(digits), true
}

func stripDigits(s string) string {
	digits := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits = append(digits, s[i])
		}
	}
	return string(digits)
}

func allSameDigits(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return len(s) > 0
}

func containsByte(s string, c byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
