package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	contents, err := os.ReadFile("Day3.txt")
	rucksack := strings.Split(string(contents), "\n")
	if err != nil {
		return
	}

	points := 0

	for i := 0; i < len(rucksack); i++ {
		container := strings.SplitAfterN(rucksack[i], "", len([]rune(rucksack[i])) / 2)
		var str = strings.Split(container[0], "")
		var str2 = strings.Split(container[1], "")
		for j := 0; j < len([]rune(container[0])); j++ {
			if str[j] == str2[j] {
				points += int(unicode.ToLower(str[j])-96)
				if unicode.IsUpper([]rune(str[j])){
					points += 26
				}
			}
		}
	}
	fmt.Println(points)
}