package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

// ReadLines reads a file and returns its content line by line.
//
// It takes the file path as a parameter and returns a slice of strings
// representing each line and an error.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to open file")
	}

	file.Close()
	return lines, nil
}

// Writes the provided data to a JSON file specified by the given path.
//
// Parameters:
// - path: the path to the JSON file.
// - data: the data to be written to the file.
//
// Returns:
// - error: an error if the file creation or JSON encoding fails.
func WriteJson(path string, data any /* any = interface{} */) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create a file")
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert to json")
	}

	file.Close()
	return nil
}
