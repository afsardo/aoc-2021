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
	lastNumber := math.MaxInt

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if lastNumber < number {
			increments++
		}
		lastNumber = number
	}

	fmt.Println("Solution:", increments)
}
