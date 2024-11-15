package utils

import (
	"bufio"
	"os"
)

func WriteToFile(filePath string, data string) error {
	// Create or open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(filePath string) (string, error) {
	// Read the entire file into a byte slice
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert byte slice to string and return
	return string(data), nil
}

func ReadFromFileToLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}
