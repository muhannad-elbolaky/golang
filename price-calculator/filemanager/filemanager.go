package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputPath  string
	OutputPath string
}

// ReadLines reads the lines from the input file specified by the FileManager's InputPath field.
//
// It returns a slice of strings containing the lines read from the file and an error if any occurred.
// If the file cannot be opened, it returns an error with the message "failed to open file".
// If there is an error while scanning the file, it returns an error with the message "failed to open file".
// If the file is successfully read, it returns the slice of strings containing the lines and a nil error.
func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)

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

func (fm *FileManager) WriteResult(data any /* any = interface{} */) error {
	file, err := os.Create(fm.OutputPath)

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

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}
