package main

import (
	"bufio"
	"log"
	"os"
)

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
		log.Println(line)
		//todo regex on line to find all instances of mul(X, Y)
	}
}

func mul(a int, b int) int {
	return a * b
}
