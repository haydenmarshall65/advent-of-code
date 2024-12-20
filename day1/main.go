package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var leftSide []string
var rightSide []string

func main() {
	file, err := os.Open("locationIDs.txt")

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
		if len(line) > 0 {
			data := strings.Split(string(line), "   ")

			leftSide = append(leftSide, data[0])
			rightSide = append(rightSide, data[1])
		}
	}

	mergeSort(&leftSide, 0, len(leftSide))
	mergeSort(&rightSide, 0, len(rightSide))
	// now they are sorted on each side, so now find the total distance between the variables on each side.

	total := 0

	for i := 0; i < len(leftSide); i++ {
		left, err := strconv.Atoi(leftSide[i])

		if err != nil {
			log.Fatal(err)
		}

		right, err := strconv.Atoi(rightSide[i])

		if err != nil {
			log.Fatal(err)
		}

		difference := math.Abs(float64(left - right))

		total += int(difference)
	}

	log.Println(total)

	// part 2, not right, needs to be fixed

	i := 0
	j := 0
	var listOfTotals map[int]int = make(map[int]int)
	var totalPart2 int
	var totalCount int = 0
	for i < len(leftSide) && j < len(rightSide) {
		numToCount, err := strconv.Atoi(leftSide[i])

		if err != nil {
			log.Fatal(err)
			break
		}
		numToCheck, err := strconv.Atoi(rightSide[j])

		if err != nil {
			log.Fatal(err)
			break
		}

		if slices.Contains(maps.Keys(listOfTotals), numToCount) {
			totalPart2 += numToCount * listOfTotals[numToCount]
			i++
			continue
		} else if numToCheck < numToCount {
			j++ //keep moving, we don't need to count this
		} else if numToCheck == numToCount {
			totalCount++
			j++
		} else {
			listOfTotals[numToCount] = totalCount
			totalPart2 += numToCount * totalCount
			totalCount = 0
			i++
		}
	}

	log.Println(totalPart2)
}

func mergeSort(arr *[]string, left int, right int) {
	if right-left <= 1 {
		return
	}

	mid := (left + right) / 2

	mergeSort(arr, left, mid)
	mergeSort(arr, mid, right)

	merge(arr, left, mid, right)
}

func merge(arr *[]string, left int, mid int, right int) {
	leftArr := make([]string, mid-left)
	rightArr := make([]string, right-mid)

	copy(leftArr, (*arr)[left:mid])
	copy(rightArr, (*arr)[mid:right])

	numLeft := len(leftArr)
	numRight := len(rightArr)

	i := 0
	j := 0
	k := left

	for i < len(leftArr) && j < len(rightArr) {
		leftArrInt, err := strconv.Atoi(leftArr[i])

		if err != nil {
			log.Fatal(err)
		}

		rightArrInt, err := strconv.Atoi(rightArr[j])

		if err != nil {
			log.Fatal(err)
		}

		if leftArrInt <= rightArrInt {
			(*arr)[k] = leftArr[i]
			i++
			k++
		} else {
			(*arr)[k] = rightArr[j]
			j++
			k++
		}
	}

	for i < numLeft {
		(*arr)[k] = leftArr[i]
		i++
		k++
	}

	for j < numRight {
		(*arr)[k] = rightArr[j]
		j++
		k++
	}
}
