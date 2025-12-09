package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type MathProblem struct {
	values []int
	op     rune
}

type OpPosition struct {
	op    rune
	index int
}

func main() {
	mathProblems := parseInput1("input.txt")
	calculate(mathProblems)

	mathProblems = parseInput2("input.txt")
	calculate(mathProblems)
}

func calculate(mps []MathProblem) {
	total := 0
	for _, mp := range mps {
		var subtotal int
		switch mp.op {
		case '+':
			subtotal = 0
			for _, value := range mp.values {
				subtotal += value
			}
		case '*':
			subtotal = 1
			for _, value := range mp.values {
				subtotal *= value
			}
		default:
			log.Fatalln("unknown operator:", mp.op)
		}
		total += subtotal
	}

	fmt.Println(total)
}

func parseInput1(filename string) []MathProblem {
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

func parseInput2(filename string) []MathProblem {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	lines := strings.Split(string(data), "\n")

	// Extract the operator line to figure out the start and end of each column
	opLine := lines[len(lines)-1]
	opPos := make([]OpPosition, 0)
	for i, char := range opLine {
		if char != ' ' {
			switch char {
			case '*', '+':
				opPos = append(opPos, OpPosition{char, i})
			default:
				log.Fatalln("unknown operator:", char)
			}
		}
	}

	// Append a terminator (.) to the end of opLine so we can loop easily
	opPos = append(opPos, OpPosition{'.', len(opLine)})

	// Find the longest line
	maxLength := 0
	for _, line := range lines[:len(lines)-1] {
		maxLength = max(maxLength, len(line))
	}

	// Pad right with spaces to align the last column
	for i := 0; i < len(lines)-1; i++ {
		diff := maxLength - len(lines[i])
		padding := string(slices.Repeat([]rune{' '}, diff))
		lines[i] += padding
	}

	numProblems := len(opPos) - 1
	mathProblems := make([]MathProblem, numProblems)

	for i := 0; i < len(opPos); i++ {
		if opPos[i].op == '.' {
			break // terminate
		}
		start := opPos[i].index
		end := opPos[i+1].index

		// Look down each column and build the number
		for c := start; c < end; c++ {
			word := ""
			for r := 0; r < len(lines)-1; r++ {
				if lines[r][c] != ' ' {
					word += string(lines[r][c])
				}
			}
			// Convert the word to integer and append to list of values
			if len(word) > 0 {
				value, _ := strconv.Atoi(word)
				mathProblems[i].values = append(mathProblems[i].values, value)
			}
		}
		mathProblems[i].op = opPos[i].op
	}

	return mathProblems
}
