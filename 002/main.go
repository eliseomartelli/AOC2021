package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	dx int
	dy int
}

var movements = map[string]Movement {
	"forward": {1,0},
	"down": {0,1},
	"up": {0,-1},
};

func main() {
	movements, err := parseFile("input")
	if err != nil {
		log.Fatalf("Error reading input file.")
	}
	fmt.Printf("Solution 1: %d\n", solution_1(movements))
	fmt.Printf("Solution 2: %d\n", solution_2(movements))
}

func solution_1(movements []Movement) int {
	var res Movement;
	for _, m := range movements {
		res.dx += m.dx;
		res.dy += m.dy;
	}
	return res.dx * res.dy;
}

func solution_2(movements []Movement) int {
	var aim int;
	var res Movement;
	for _, m := range movements {
		if m.dx == 0 {
			aim += m.dy;
		} else {
			res.dx += m.dx;
			res.dy += m.dx * aim;
		}
	}
	return res.dx * res.dy;
}

func parseFile(path string) ([]Movement, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []Movement
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		distance, err := strconv.ParseInt(s[1], 10, 0);
		if err != nil {
			return nil, err
		}
		direction := movements[s[0]]
		movement := Movement{
			dx: direction.dx * int(distance),
			dy: direction.dy * int(distance),
		}
		lines = append(lines, movement)
	}
	return lines, nil
}
