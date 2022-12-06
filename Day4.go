package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("Day4.txt")
	group := strings.Split(string(contents), "\n")
	if err != nil {
		return
	}
	points := 0

	for i := 0; i < len(group); i++ {
		elf := strings.Split(group[i], ",")

		input1 := strings.Split(elf[0], "-")
		input2 := strings.Split(elf[1], "-")

		i1, err := strconv.Atoi(input1[0])
		i12, err := strconv.Atoi(input1[1])
		i2, err := strconv.Atoi(input2[0])
		i21, err := strconv.Atoi(input2[1])
		if err != nil {
			return
		}

		if (i1 > i2 && i12 > i21 && i12 < i2) || (i2 > i1 && i21 > i12 && i21 < i1) {
			points++
		}
	}
	fmt.Print(points)
}
