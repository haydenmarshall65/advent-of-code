package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var safeReactors int = 0

// only considered stable if all values in line are:
// 1. increasing together or decreasing together
// 2. adjacent values are at least one and at most 3 apart.
func main() {
	file, err := os.Open("reactorLevels.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		reactorLevels := strings.Split(string(line), " ")

		increasing := false
		decreasing := false
		reactorIsSafe := true

		for i := 0; i < len(reactorLevels)-1; i++ {
			if i+1 > len(reactorLevels)+1 {
				continue // don't go out of bounds
			}

			currentLevel, err := strconv.Atoi(reactorLevels[i])

			if err != nil {
				log.Fatal(err)
			}

			nextLevel, err := strconv.Atoi(reactorLevels[i+1])

			if err != nil {
				log.Fatal(err)
			}

			if currentLevel < nextLevel {
				increasing = true
			} else if currentLevel > nextLevel {
				decreasing = true
			}

			if increasing && decreasing {
				reactorIsSafe = false
				break
			}

			if !adjacentLevelsAreSafe(currentLevel, nextLevel) {
				reactorIsSafe = false
				break
			}
		}

		if reactorIsSafe {
			safeReactors++
		}
	}

	log.Println(safeReactors)
}

// check if adjacent levels increase or decrease by 1-3
func adjacentLevelsAreSafe(currentLevel int, nextLevel int) bool {
	passes := false

	difference := math.Abs(float64(currentLevel - nextLevel))

	if difference > 0 && difference < 4 {
		passes = true
	}

	return passes
}
