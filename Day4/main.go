package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	grid := parseInput("input.txt")

	part1(grid)
	part2(grid)
}

func part1(grid [][]rune) {
	coords := findAllCoordinates(grid)
	fmt.Println(len(coords))
}

func part2(grid [][]rune) {
	total := 0

	for {
		coords := findAllCoordinates(grid)
		if len(coords) == 0 {
			break
		}
		total += len(coords)

		// Update the grid
		for _, coord := range coords {
			nr, nc := coord[0], coord[1]
			grid[nr][nc] = '.'
		}
	}

	fmt.Println(total)
}

func findAllCoordinates(grid [][]rune) [][2]int {
	coords := make([][2]int, 0)
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] == '@' {
				adj := numAdjacent(r, c, grid)
				if adj < 4 {
					coords = append(coords, [2]int{r, c})
				}
			}
		}
	}

	return coords
}

func numAdjacent(r, c int, grid [][]rune) int {
	count := 0
	coords := [][2]int{
		{r - 1, c - 1},
		{r - 1, c},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c + 1},
		{r + 1, c - 1},
		{r + 1, c},
		{r + 1, c + 1},
	}

	for _, coord := range coords {
		if grid[coord[0]][coord[1]] == '@' {
			count++
		}
	}

	return count
}

func parseInput(filename string) [][]rune {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var grid [][]rune
	lines := strings.Split(string(data), "\n")
	n := len(lines[0])

	// Add a buffer around all 4 sides for ease of checking adjacency
	grid = append(grid, newBufferLine(n+2, '.'))

	for _, line := range lines {
		extLine := "." + line + "."
		grid = append(grid, []rune(extLine))
	}

	grid = append(grid, newBufferLine(n+2, '.'))
	return grid
}

func newBufferLine(size int, char rune) []rune {
	bufferLine := make([]rune, size)
	for i := 0; i < len(bufferLine); i++ {
		bufferLine[i] = char
	}

	return bufferLine
}
