package main

import (
	"fmt"
)

// // reads the banner(i.e access the character's storage)
// func readBanner(file string) ([]string, error) {
// 	ban, err := os.ReadFile(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	line := strings.Split(string(ban), "\n")
// 	return line, err
// }

// gets the character's positiion in the banner
func getLine(ban []string, c rune) []string {
	if c < 32 || c > 126 {
		return nil
	}
	begin := (int(c) - 32) * 9
	return ban[begin+1 : begin+9]
}

// prints the characters' ascii art from the position in the banner
func printArt(print []string, ban []string) {
	for _, f := range print {
		if f == "" {
			fmt.Println()
			continue
		}
		for i := 0; i <= 7; i++ {
			for _, g := range f {
				charArt := getLine(ban, g)
				if charArt == nil {
					continue
				}
				fmt.Print(charArt[i])
			}
			fmt.Println()
		}
	}
}
