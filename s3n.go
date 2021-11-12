package s3n

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

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
