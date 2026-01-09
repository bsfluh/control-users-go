package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("error reading: %w", err)
		}
		return "", fmt.Errorf("invalid input")
	}
	text := strings.TrimSpace(scanner.Text())
	return text, nil
}

func ReadBool(scanner *bufio.Scanner) (bool, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return false, fmt.Errorf("error reading: %w", err)
		}
		return false, fmt.Errorf("invalid input")
	}
	resultStr := strings.TrimSpace(scanner.Text())
	result, err := strconv.ParseBool(resultStr)
	if err != nil {
		return false, fmt.Errorf("invalid parse status")
	}
	return result, nil

}
func ReadInt(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, fmt.Errorf("error reading: %w", err)
		}
		return 0, fmt.Errorf("invalid input")
	}
	numStr := strings.TrimSpace(scanner.Text())
	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		return 0, fmt.Errorf("invalid parse status")
	}
	return num, nil

}
