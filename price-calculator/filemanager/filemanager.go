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

func (fm *FileManager) WriteJson(data any /* any = interface{} */) error {
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
