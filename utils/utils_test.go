package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToDate(t *testing.T) {
	t.Run("Valid date format", func(t *testing.T) {
		date := "2022-01-01T12:34:56"
		fname := "date"
		expected := time.Date(2022, time.January, 1, 12, 34, 56, 0, time.UTC)

		result, err := ParseStringToDate(date, fname)

		assert.Nil(t, err, "Expected no error, but got an error")
		assert.Equal(t, expected, result, "Expected result to be equal to expected")
	})

	// Caso de prueba: cadena de fecha vacía
	t.Run("Empty date string", func(t *testing.T) {
		date := ""
		fname := "date"
		expected := time.Time{}

		result, err := ParseStringToDate(date, fname)

		assert.Nil(t, err, "Expected no error, but got an error")
		assert.Equal(t, expected, result, "Expected result to be equal to expected")
	})

	// Caso de prueba: formato de fecha inválido
	t.Run("Invalid date format", func(t *testing.T) {
		date := "invalid"
		fname := "date"
		expectedError := "bad_request"

		_, err := ParseStringToDate(date, fname)

		assert.NotNil(t, err, "Expected an error, but got no error")
		assert.EqualError(t, err, expectedError)
	})
}
