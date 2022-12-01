package main

import (
	"fmt"
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

	positionX := 0
	positionY := 0
	aim := 0

	for _, line := range lines {
		action := strings.Split(line, " ")[0]
		value, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic(err)
		}

		if action == "forward" {
			positionX += value
			positionY += aim * value
		}

		if action == "down" {
			aim += value
		}

		if action == "up" {
			aim -= value
		}
	}

	fmt.Println("Solution:", positionX*positionY)
}
