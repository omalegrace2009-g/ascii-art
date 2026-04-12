package main

import (
	"bufio"
	"os"
)

func LoadBanner(filename string) ([]string, error) {
	bannerf, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer bannerf.Close()
	scanner := bufio.NewScanner(bannerf)
	var fileLines []string

	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileLines, nil
}