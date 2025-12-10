package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/DennisPing/AOC-2025/utils"
)

func main() {
	grid := parseInput("input.txt")
	part1(grid)
	part2(grid)
}

func part1(grid []string) {
	start := strings.Index(grid[0], "S")
	count := 0

	beamIdx := make([]int, 0) // Beam indices
	beamIdx = append(beamIdx, start)

	// Only even rows matter, skip odd rows
	for r := 2; r < len(grid)-1; r += 2 {
		splitterIdx := utils.FindAllIndices([]rune(grid[r]), '^')

		for _, splitter := range splitterIdx {
			for j, beam := range beamIdx {
				if splitter == beam {
					// Delete the old beam
					beamIdx = slices.Delete(beamIdx, j, j+1)

					// Add the left beam
					if beam-1 >= 0 && !slices.Contains(beamIdx, beam-1) {
						beamIdx = append(beamIdx, beam-1)
					}

					// Add the right beam
					if beam+1 < len(grid[0]) && !slices.Contains(beamIdx, beam+1) {
						beamIdx = append(beamIdx, beam+1)
					}

					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func part2(grid []string) {
	// Don't build a polytree and do DFS. Too complex.
	// Instead, track how many beams currently exist in each column as we move downwards.
	// Each row is a time step. The # of beams in each column == # of possibilities.

	length := len(grid)
	width := len(grid[0])
	start := strings.Index(grid[0], "S")

	beams := make([]int, width) // Array of beam counts in each position
	beams[start] = 1

	for r := 2; r < length; r += 2 {
		row := grid[r]
		next := make([]int, width) // Array of next beam counts in each position

		for c := 0; c < width; c++ {
			k := beams[c]
			if k == 0 {
				continue // No beams to propagate down
			}

			if c < len(row) && row[c] == '^' {
				// Add count k to left column
				if c-1 >= 0 {
					next[c-1] += k
				}

				// Add count k to right column
				if c+1 < width {
					next[c+1] += k
				}
			} else {
				// Propagate straight down
				next[c] += k
			}
		}

		beams = next
	}

	totalPaths := 0
	for _, k := range beams {
		totalPaths += k
	}

	fmt.Println(totalPaths)
}

func parseInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	return strings.Split(string(data), "\n")
}
