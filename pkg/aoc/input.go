package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Get(day int) (input []string, err error) {
	inputPath, exists := getInputPath(1)

	if exists {
		input, err = readInput(inputPath)
	} else {
		input, err = getFromServer(day)
		if err == nil {
			err = writeInput(inputPath, input)
		}
	}
	return input, err
}

func getInputPath(day int) (path string, exists bool) {
	// Get the working directory where main.go is
	wd, err := os.Getwd()
	if err != nil {
		return "", false
	}

	// Construct path relative to project root
	path = filepath.Join(wd, "..", "..", "days", fmt.Sprintf("day%02d", day), "input.txt")
	_, err = os.Stat(path)
	exists = !os.IsNotExist(err)

	return path, exists
}

func writeInput(path string, lines []string) error {
	// Create the directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	// Join lines with newlines and write to file
	content := strings.Join(lines, "\n")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func readInput(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	// Split by newlines and trim whitespace
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	return lines, nil
}

func getFromServer(day int) ([]string, error) {
	var client http.Client
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	session := os.Getenv("SESSION")

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	body := string(bytes)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Server returned %d, response body: %s", res.StatusCode, body)

	}
	return strings.Split(body, "\n"), nil
}
