package s3n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreprocessToList(t *testing.T) {
	want := 12
	result, err := preprocessToList("756123456789")

	assert.NoError(t, err)
	assert.Len(t, result, want, "Parsed number contain 12 elements")

	result, err = preprocessToList("756.1234.5678.9")

	assert.NoError(t, err)
	assert.Len(t, result, want, "Parsed dotted number contain 12 elements")

	result, err = preprocessToList("756.asdf.ghjk.q")
	assert.Error(t, err, "Parsing should error when the string contains characters that are not digits")
	assert.Nil(t, result)
}

func TestCheckSum(t *testing.T) {
	// 756.1234.5678.97
	number := "756123456789"
	digits, _ := preprocessToList(number)
	want := 7
	result := checksum(digits)

	assert.Equal(t, result, want, "Checksum should be computed correctly")

	// 756.1234.5678.97
	number = "756.1234.5678.9"
	digits, _ = preprocessToList(number)
	want = 7
	result = checksum(digits)

	assert.Equal(t, result, want, "Checksum should be computed correctly")

	// 756.3047.5009.62
	number = "756.3047.5009.6"
	digits, _ = preprocessToList(number)
	want = 2
	result = checksum(digits)

	assert.Equal(t, result, want, "Checksum should be computed correctly")
}

func TestCheckValidity(t *testing.T) {
	assert.True(t, IsValid("756.9217.0769.85"))

	assert.True(t, IsValid("756.3047.5009.62"))

	assert.True(t, IsValid("756.1234.5678.97"))

	assert.False(t, IsValid("756.9217.0769.83"))

	assert.False(t, IsValid("718.9217.0769.83"))

	assert.False(t, IsValid("756.9217asf.a"))
}

func TestIsValidWithoutDotted(t *testing.T) {
	dotted := "756.9217.0769.85"
	dottedValid := IsValid(dotted)

	number := "7569217076985"
	valid := IsValid(number)

	assert.Equal(t, dottedValid, valid, "Number should be valid with or without dots")
}