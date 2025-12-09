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
}

func part1(grid []string) {
	start := strings.Index(grid[0], "S")
	count := 0

	beamIdx := make([]int, 0) // Beam indices
	beamIdx = append(beamIdx, start)

	for r := 2; r < len(grid)-1; r++ {
		for c := 0; c < len(grid[0]); c++ {
			// If odd row, then skip
			if r%2 != 0 {
				continue
			}

			splitterIdx := utils.FindAllIndices([]rune(grid[r]), '^')
			for _, splitter := range splitterIdx {
				for j, beam := range beamIdx {
					if splitter == beam {
						// Delete the old beam
						beamIdx = slices.Delete(beamIdx, j, j+1)

						// Add the left beam
						if !slices.Contains(beamIdx, beam-1) {
							beamIdx = append(beamIdx, beam-1)
						}

						// Add the right beam
						if !slices.Contains(beamIdx, beam+1) {
							beamIdx = append(beamIdx, beam+1)
						}

						count++
					}
				}
			}

		}
	}

	fmt.Println(count)
}

func parseInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	return strings.Split(string(data), "\n")
}
