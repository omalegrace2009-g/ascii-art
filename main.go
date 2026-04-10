package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func validateInput(args []string) ([]string, error) {
	result := []string{}
	for _, row := range args {
		for _, char := range row {
			if !(char >= 32 && char <= 126 || char == '\n') {
				return nil, errors.New("bad character found")
			}

		}
		result = append(result, string(row))

	}
	return result, nil

}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [string]")
		return
	}

	input := os.Args[1]

	// Edge cases for no argument.
	if len(input) != 2 {
		return
	}

	// Split input on \n to get seperate rows
	splitIntput := strings.Split(input, "\n")

	valid, err := validateInput(splitIntput)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fit, err := LoadBanner("banner/standard.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	printArt(splitIntput, fit)

	fmt.Println(valid)

}
