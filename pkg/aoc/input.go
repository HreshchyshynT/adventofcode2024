package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Get(day int) ([]string, error) {
	return getFromServer(day)
}

func getFromServer(day int) ([]string, error) {
	var client http.Client
	// TODO: change year to 2024
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)

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
