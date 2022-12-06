package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("Day6.txt")
	letter := strings.Split(string(contents), "")
	if err != nil {
		return
	}
	points := 0

	for i := 0; i < len(letter); i++ {
		if letter[i] != letter[i+1] && letter[i+1] != letter[i+2] && letter[i+2] != letter[i+3] {
			i = len(letter)
		}
		points++
	}
	fmt.Println(points)
}
