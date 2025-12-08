package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MathProblem struct {
	values []int
	op     rune
}

func main() {
	mps := parseInput("input.txt")
	part1(mps)
}

func part1(mps []MathProblem) {
	total := 0
	for _, mp := range mps {
		var subtotal int
		if mp.op == '+' {
			subtotal = 0
			for _, value := range mp.values {
				subtotal += value
			}
		} else if mp.op == '*' {
			subtotal = 1
			for _, value := range mp.values {
				subtotal *= value
			}
		}
		total += subtotal
	}

	fmt.Println(total)
}

func parseInput(filename string) []MathProblem {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	lines := strings.Split(string(data), "\n")

	// Find the size of each problem M x N
	parts := strings.Fields(lines[0])
	numProblems := len(parts)   // M
	numValues := len(lines) - 1 // N

	mathProblems := make([]MathProblem, numProblems)
	for i := range mathProblems {
		mathProblems[i].values = make([]int, numValues)
	}

	for i, line := range lines[:len(lines)-1] {
		for j, word := range strings.Fields(line) {
			value, _ := strconv.Atoi(word)
			mathProblems[j].values[i] = value
		}
	}

	opsLine := lines[len(lines)-1]
	for k, op := range strings.Fields(opsLine) {
		mathProblems[k].op = rune(op[0]) // The operator is always 1 char long
	}

	return mathProblems
}
