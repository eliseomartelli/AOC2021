package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Converts boolean to int.
var b2i = map[bool]int{false: 0, true: 1}

func main() {
	numbers, err := arrayFromFile("input")
	if err != nil {
		log.Fatalf("Error reading input file.")
	}
	fmt.Printf("Solution 1: %d\n", solution_1(numbers))
	fmt.Printf("Solution 2: %d\n", solution_2(numbers))
}

func solution_1(numbers []int) int {
	return solution_1_r(numbers, len(numbers)-1)
}

func solution_1_r(numbers []int, i int) int {
	if i == 0 {
		return 0
	}
	// + 1 if previous number is less than current number.
	// + 0 if previous number is equal or greater than current number.
	return solution_1_r(numbers, i-1) + b2i[numbers[i-1] < numbers[i]]
}

func solution_2(numbers []int) int {
	var count = -1 // Starts at -1 to account for first window.
	var prevSum int
	for i := 2; i < len(numbers); i++ { // Starts at third position.
		var current = numbers[i] + numbers[i-1] + numbers[i-2]
		count += b2i[current > prevSum] // 1 if current > prevSum. 0 otherwise.
		prevSum = current
	}
	return count
}

func arrayFromFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return nil, err
		}
		lines = append(lines, int(number))
	}
	return lines, nil
}
