package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	idRanges, ids := parseInput("input.txt")
	slices.SortFunc(idRanges, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})
	idRanges = mergeRanges(idRanges)

	part1(idRanges, ids)
	part2(idRanges)
}

func part1(idRanges [][2]int, ids []int) {
	total := 0
	for _, id := range ids {
		idx := findRange(idRanges, id)
		if idx >= 0 && id <= idRanges[idx][1] {
			total++
		}
	}
	fmt.Println(total)
}

func part2(idRanges [][2]int) {
	total := 0
	for _, ids := range idRanges {
		fmt.Println(ids)
		high := ids[1]
		low := ids[0]
		total += (high - low) + 1
	}
	fmt.Println(total)
}

func mergeRanges(idRanges [][2]int) [][2]int {
	i := 0
	for i < len(idRanges)-1 {
		first := idRanges[i]
		second := idRanges[i+1]

		if second[0] <= first[1] {
			ceiling := max(first[1], second[1])
			idRanges[i] = [2]int{first[0], ceiling}
			idRanges = slices.Delete(idRanges, i+1, i+2)
		} else {
			i++
		}
	}

	return idRanges
}

// Finds the index with the largest starting value <= target
func findRange(idRanges [][2]int, target int) int {
	low := 0
	high := len(idRanges) - 1
	idx := -1

	/*
		Modified binary search

		starts: [3, 10, 16, 25]
		indices: 0   1   2   3
		target: 17

		Only indices 3, 10, 16 are valid
		And index 3 has the largest value (16)
	*/

	for low <= high {
		mid := low + (high-low)/2
		if idRanges[mid][0] <= target {
			idx = mid     // This 'mid' is at least within the floor
			low = mid + 1 // Try to find a better one later by raising the floor
		} else {
			// Any 'mid' that is greater than target is automatically invalid
			high = mid - 1
		}
	}

	return idx
}

func parseInput(filename string) ([][2]int, []int) {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var idRanges [][2]int
	var ids []int

	sections := strings.Split(string(data), "\n\n")

	top := sections[0]
	lines := strings.SplitSeq(top, "\n")
	for line := range lines {
		parts := strings.Split(line, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		idRanges = append(idRanges, [2]int{low, high})
	}

	bot := sections[1]
	lines = strings.SplitSeq(bot, "\n")
	for line := range lines {
		id, _ := strconv.Atoi(line)
		ids = append(ids, id)
	}

	return idRanges, ids
}
