package scalarsbr

import "errors"

var ErrInvalidCNPJ = errors.New("invalid CNPJ")

var (
	cnpjFirstWeights  = [...]int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	cnpjSecondWeights = [...]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
)

// ValidateCNPJ validates a 14-digit Brazilian company taxpayer number.
func ValidateCNPJ(s string) error {
	digits, ok := collectDigits(s, "./- ")
	if !ok || !validCNPJDigits(digits) {
		return ErrInvalidCNPJ
	}
	return nil
}

// FormatCNPJ strips non-digits, validates the value, and returns XX.XXX.XXX/XXXX-XX.
func FormatCNPJ(raw string) (string, error) {
	digits := stripDigits(raw)
	if !validCNPJDigits(digits) {
		return "", ErrInvalidCNPJ
	}
	return digits[:2] + "." + digits[2:5] + "." + digits[5:8] + "/" + digits[8:12] + "-" + digits[12:], nil
}

func validCNPJDigits(digits string) bool {
	if len(digits) != 14 || allSameDigits(digits) {
		return false
	}
	return int(digits[12]-'0') == cnpjDigit(digits[:12], cnpjFirstWeights[:]) &&
		int(digits[13]-'0') == cnpjDigit(digits[:13], cnpjSecondWeights[:])
}

func cnpjDigit(digits string, weights []int) int {
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += int(digits[i]-'0') * weights[i]
	}
	digit := 11 - (sum % 11)
	if digit >= 10 {
		return 0
	}
	return digit
}
