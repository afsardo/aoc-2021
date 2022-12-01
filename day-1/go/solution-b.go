package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	increments := 0
	lastWindow := math.MaxInt

	for i := 0; i < len(lines); i++ {
		if len(lines) < i+3 {
			break
		}
		window := 0
		for _, line := range lines[i : i+3] {
			number, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			window += number
		}
		if lastWindow < window {
			increments++
		}
		lastWindow = window
	}

	fmt.Println("Solution:", increments)
}
