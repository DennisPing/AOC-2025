package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	seqs := parseInput("input.txt")
	part1(seqs)
	part2(seqs)
}

func part1(seqs [][]int) {
	sum := 0
	for _, seq := range seqs {
		val := findTwoLargest(seq)
		sum += val
	}
	fmt.Println(sum)
}

func part2(seqs [][]int) {
	sum := 0
	for _, seq := range seqs {
		val := findKLargest(seq, 12)
		sum += val
	}
	fmt.Println(sum)
}

func findTwoLargest(seq []int) int {
	m := 0

	// Find the largest tenth digit
	for i := 0; i < len(seq)-1; i++ {
		if seq[i] > seq[m] {
			m = i
		}
	}

	// Find the largest single digit after index m
	n := m + 1
	for i := m + 1; i < len(seq); i++ {
		if seq[i] > seq[n] {
			n = i
		}
	}

	concatStr := fmt.Sprintf("%d%d", seq[m], seq[n])
	concatInt, _ := strconv.Atoi(concatStr)
	return concatInt
}

func findKLargest(seq []int, k int) int {
	idxs := make([]int, k) // Indexes of the max values
	p := 0                 // Current pointer

	for i := 0; i < k; i++ {
		offset := k - i - 1
		idxs[i] = p // Set the initial max to the current pointer

		for j := p; j < len(seq)-offset; j++ {
			if seq[j] > seq[idxs[i]] {
				idxs[i] = j
			}
		}

		// Update the pointer to the i+1 position
		p = idxs[i] + 1
	}

	concatStr := ""
	for _, idx := range idxs {
		concatStr = fmt.Sprintf("%s%d", concatStr, seq[idx])
	}

	concatInt, _ := strconv.Atoi(concatStr)
	return concatInt
}

func parseInput(filename string) [][]int {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	lines := strings.SplitSeq(string(data), "\n")

	var seqs [][]int
	for line := range lines {
		line = strings.TrimSpace(line)
		var seq []int
		for _, v := range line {
			joltage, _ := strconv.Atoi(string(v))
			seq = append(seq, joltage)
		}

		seqs = append(seqs, seq)
	}

	return seqs
}
