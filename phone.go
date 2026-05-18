package scalarsbr

import "errors"

var ErrInvalidPhone = errors.New("invalid Brazilian phone")

var validDDDs = map[string]bool{
	"11": true, "12": true, "13": true, "14": true, "15": true, "16": true, "17": true, "18": true, "19": true,
	"21": true, "22": true, "24": true, "27": true, "28": true,
	"31": true, "32": true, "33": true, "34": true, "35": true, "37": true, "38": true,
	"41": true, "42": true, "43": true, "44": true, "45": true, "46": true, "47": true, "48": true, "49": true,
	"51": true, "53": true, "54": true, "55": true,
	"61": true, "62": true, "63": true, "64": true, "65": true, "66": true, "67": true, "68": true, "69": true,
	"71": true, "73": true, "74": true, "75": true, "77": true, "79": true,
	"81": true, "82": true, "83": true, "84": true, "85": true, "86": true, "87": true, "88": true, "89": true,
	"91": true, "92": true, "93": true, "94": true, "95": true, "96": true, "97": true, "98": true, "99": true,
}

// ValidatePhone validates a Brazilian DDD plus 8- or 9-digit local number.
func ValidatePhone(s string) error {
	if !validPhoneDigits(stripDigits(s)) {
		return ErrInvalidPhone
	}
	return nil
}

// FormatPhone strips non-digits, applies dddDefault to local-only input, and formats the number.
func FormatPhone(raw, dddDefault string) (string, error) {
	digits := stripDigits(raw)
	if len(digits) == 8 || len(digits) == 9 {
		ddd := stripDigits(dddDefault)
		if len(ddd) != 2 || !validDDDs[ddd] {
			return "", ErrInvalidPhone
		}
		digits = ddd + digits
	}
	if !validPhoneDigits(digits) {
		return "", ErrInvalidPhone
	}
	if len(digits) == 10 {
		return "(" + digits[:2] + ") " + digits[2:6] + "-" + digits[6:], nil
	}
	return "(" + digits[:2] + ") " + digits[2:7] + "-" + digits[7:], nil
}

func validPhoneDigits(digits string) bool {
	return (len(digits) == 10 || len(digits) == 11) && validDDDs[digits[:2]]
}
