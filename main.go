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
	input := os.Args

	// Edge cases for no argument.
	if len(input) != 2 {
		return
	}

	// Split input on \n to get seperate rows
	splitIntput := strings.Split(input[1], "\n")

	valid, err := validateInput(splitIntput)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(valid)

}
