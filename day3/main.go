package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var total int

func main() {
	file, err := os.Open("corruptedMemory.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// go line by line and use regex to find all cases of
	// mul(X, Y) and then run based on the X and Y.
	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}
		// stringifiedLine := string(line)

		regex := regexp.MustCompile(`mul\((([0-9]+),([0-9]+))\)`)

		match := regex.FindAllStringSubmatch(string(line), -1)

		evaluateAllMuls(match)
	}

	log.Println(total)
}

func mul(a int, b int) int {
	return a * b
}

func evaluateAllMuls(allMuls [][]string) {
	i := 0
	for i < len(allMuls) {
		x, err := strconv.Atoi(allMuls[i][2])

		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(allMuls[i][3])

		if err != nil {
			log.Fatal(err)
		}

		total += mul(x, y)
		i++
	}
}
