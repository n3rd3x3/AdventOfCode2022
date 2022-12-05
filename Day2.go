package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("Day2.txt")
	game := strings.Split(string(contents), "\n")
	if err != nil {
		return
	}
	points := 0
	for i := 0; i < len(game); i++ {
		spl := strings.Split(game[i], " ")
		if spl[0] == "A" {
			if spl[1] == "X" {
				points += 4
			}
			if spl[1] == "Y" {
				points += 8
			}
			if spl[1] == "Z" {
				points += 3
			}
		}

		if spl[0] == "B" {
			if spl[1] == "X" {
				points++
			}
			if spl[1] == "Y" {
				points += 5
			}
			if spl[1] == "Z" {
				points += 9
			}
		}

		if spl[0] == "C" {
			if spl[1] == "X" {
				points += 7
			}
			if spl[1] == "Y" {
				points += 2
			}
			if spl[1] == "Z" {
				points += 6
			}
		}
	}
	fmt.Println(points)
}
