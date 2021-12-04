package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers, err := readFile("input")
	if err != nil {
		log.Fatalf("Error reading input file.")
	}
	fmt.Printf("Solution 1: %d\n", solution_1(numbers))
	fmt.Printf("Solution 2: %d\n", solution_2(numbers))
}

func solution_1(numbers []string) int64 {
	var buckets [12]int
	for _, n := range numbers {
		for i, c := range n {
			if c == '1' {
				buckets[i]++
			}
		}
	}
	nlength := len(numbers)
	var s1 string
	for _, b := range buckets {
		if b >= nlength/2 {
			s1 += "1"
		} else {
			s1 += "0"
		}
	}
	i1, err := strconv.ParseInt(s1, 2, 64)
	if err != nil {
		return -1
	}
	i2 := 4095 - i1 // [1111 1111 1111]2 - i1. Reverse of the original number.
	return i1 * i2
}

func filter(numbers []string, inverse bool) int64 {
	// To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.
	// To find CO2 scrubber rating, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.

	strLength := len(numbers[0])

	var filtered []string
	toProcess := numbers

	for i := 0; i < strLength; i++ {
		count := 0
		var toTake = "0"
		for _, sequence := range toProcess {
			if strings.Split(sequence, "")[i] == "0" {
				count++
			}
		}

		predicate := count > len(toProcess) / 2

		if !predicate && inverse || predicate && !inverse {
			toTake = "0"
		} else {
			toTake = "1"
		}

		for _, sequence := range toProcess {
			if strings.Split(sequence, "")[i] == toTake {
				filtered = append(filtered, sequence)
			}
		}
		if len(filtered) >= 1 {
			toProcess = filtered
			filtered = make([]string, 0)
		}
	}

	i1, err := strconv.ParseInt(toProcess[0], 2, 64);
	if err != nil {
		return -1;
	}
	return i1;
}

func solution_2(numbers []string) int64 {
	return filter(numbers, false) * filter(numbers, true)
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
