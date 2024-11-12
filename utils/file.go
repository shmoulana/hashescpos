package utils

import (
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
