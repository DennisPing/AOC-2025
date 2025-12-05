package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sequence []int // positive number is R, negative number is L

func main() {
	seq := parseInput("input.txt")

	part1(seq)
	part2(seq)

}

func part1(seq Sequence) {
	idx := 50
	count := 0
	for _, click := range seq {
		idx = rotate1(idx, click)
		if idx == 0 {
			count++
		}
	}

	fmt.Println(count)
}

func part2(seq Sequence) {
	idx := 50
	count := 0
	for _, click := range seq {
		newIdx, loops := rotate2(idx, click)
		count += loops
		idx = newIdx
	}

	fmt.Println(count)
}

// Rotate the position n number of clicks to get the new position
func rotate1(pos, clicks int) int {
	// Non-negative modulo
	return ((pos+clicks)%100 + 100) % 100
}

// Rotate the position n number of clicks to get the new position and number of loops
func rotate2(pos, clicks int) (int, int) {
	newPos := ((pos+clicks)%100 + 100) % 100

	loops := 0
	if clicks > 0 {
		// Positive direction
		loops = (pos + clicks) / 100
	} else {
		// Negative direction
		absClicks := -clicks

		// [-100] <----- pos ----- [0]
		distToZero := pos

		// Edge case: if input position is 0, then treat it as 100
		if distToZero == 0 {
			distToZero = 100
		}

		if absClicks >= distToZero {
			// We have at least 1 loop plus whatever integer division gives us
			loops = 1 + (absClicks-distToZero)/100
		}
	}

	return newPos, loops
}

func parseInput(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var seq Sequence
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		dir := string(line[0])
		if dir == "R" {
			clicks, _ := strconv.Atoi(line[1:])
			seq = append(seq, clicks)
		} else if dir == "L" {
			clicks, _ := strconv.Atoi(line[1:])
			seq = append(seq, -clicks)
		} else {
			log.Fatalln("invalid line:", line)
		}
	}

	return seq
}
