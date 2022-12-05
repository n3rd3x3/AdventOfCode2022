package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("Day1.txt")
	cat := strings.Split(string(contents), "\n\n")
	if err != nil {
		return
	}
	for i := 0; i < len(cat); i++ {
		spl := strings.Split(cat[i], "\n")
		count := 0
		for j := 0; j < len(spl); j++ {
			aj, err := strconv.Atoi(spl[j])
			if err != nil {
				panic(err)
			}
			count += aj
		}
		fmt.Println(count)
	}
}
