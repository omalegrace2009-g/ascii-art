package main

import (
	"fmt"
)

// GETS THE CHARACTER'S POSITION IN THE BANNER
func getLine(ban []string, c rune) []string {

	if c < 32 || c > 126 {
		return nil
	}

	begin := (int(c) - 32) * 9
	if begin+9 > len(ban) || begin+1 >= len(ban) {
		return nil
	}
	return ban[begin+1 : begin+9]
}

// PRINTS THE CHARACTER'S ASCII ART FROM THE POSITION IN THE BANNER
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
