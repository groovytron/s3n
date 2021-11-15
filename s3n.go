// Package s3n provides swiss social security numbers validation
package s3n

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const dottedFormatString = "%d%d%d.%d%d%d%d.%d%d%d%d.%d%d"
const dotlessFormatString = "%d%d%d%d%d%d%d%d%d%d%d%d%d"
const invalidNumberErrorMessage = "Social security number is invalid"

func preprocessToList(number string) ([]int, error) {
	parts := strings.Split(number, ".")

	result := make([]int, 0)

	for _, segment := range parts {
		for _, c := range segment {

			buf := make([]byte, 1)
			_ = utf8.EncodeRune(buf, c)

			value, err := strconv.Atoi(string(buf))

			if err != nil {
				return nil, errors.New("Invalid character found")
			}

			result = append(result, value)
		}
	}

	return result, nil
}

func checksum(digits []int) int {
	sum := 0

	for index, value := range digits {
		if index%2 != 0 {
			sum += 3 * value
		} else {
			sum += value
		}
	}

	rest := sum % 10

	if rest != 0 {
		rest = 10 - rest
	}

	return rest
}

// Validates a given social security number. The numbers can be dotted or not.
func IsValid(number string) bool {
	if !strings.HasPrefix(number, "756") {
		return false
	}

	digits, err := preprocessToList(number)

	if err != nil {
		return false
	}

	givenChecksum := digits[len(digits)-1]
	digits = digits[:len(digits)-1]

	checksum := checksum(digits)

	if err != nil || givenChecksum != checksum {
		return false
	}

	return true
}

// Formats a given social security number with dots.
// "7569217076985" becomes "756.9217.0769.85".
// Please not that the number needs to be valid to be properly formatted.
func DottedFormat(number string) (string, error) {
	if !IsValid(number) {
		return "", errors.New(invalidNumberErrorMessage)
	}

	digits, _ := preprocessToList(number)

	return fmt.Sprintf(
			dottedFormatString,
			digits[0],
			digits[1],
			digits[2],
			digits[3],
			digits[4],
			digits[5],
			digits[6],
			digits[7],
			digits[8],
			digits[9],
			digits[10],
			digits[11],
			digits[12],
		),
		nil
}

// Formats a given social security number without dots.
// "756.9217.0769.85" becomes "7569217076985".
// Please not that the number needs to be valid to be properly formatted.
func DotlessFormat(number string) (string, error) {
	if !IsValid(number) {
		return "", errors.New(invalidNumberErrorMessage)
	}

	digits, _ := preprocessToList(number)

	return fmt.Sprintf(
			dotlessFormatString,
			digits[0],
			digits[1],
			digits[2],
			digits[3],
			digits[4],
			digits[5],
			digits[6],
			digits[7],
			digits[8],
			digits[9],
			digits[10],
			digits[11],
			digits[12],
		),
		nil
}
