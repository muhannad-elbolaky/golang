package conversion

import (
	"errors"
	"strconv"
)

// StringsToFloat converts a slice of strings to a slice of floats.
//
// Parameters:
// - strings: a slice of strings to be converted.
//
// Returns:
// - []float64: a slice of floats representing the converted values.
// - error: an error if any of the strings cannot be converted to floats.
func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, string := range strings {
		floatValue, err := strconv.ParseFloat(string, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floats = append(floats, floatValue)
	}

	return floats, nil
}
