package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IdRanges [][2]int

func main() {
	idRanges := parseInput("input.txt")
	part1(idRanges)
	part2(idRanges)
}

func part1(idRanges IdRanges) {
	sum := 0
	for _, ids := range idRanges {
		start, end := ids[0], ids[1]
		for num := start; num <= end; num++ {
			if repeatedTwice(num) {
				sum += num
			}
		}
	}
	fmt.Println(sum)
}

func part2(idRanges IdRanges) {
	sum := 0
	for _, ids := range idRanges {
		start, end := ids[0], ids[1]
		for num := start; num <= end; num++ {
			if repeatedMultiple(num) {
				sum += num
			}
		}
	}
	fmt.Println(sum)
}

func repeatedTwice(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)

	if n%2 != 0 {
		return false
	}

	left := s[:n/2]
	right := s[n/2:]
	return left == right
}

func repeatedMultiple(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)

	// Let t = chunk size
	for t := 1; t <= n/2; t++ {
		if n%t != 0 {
			continue // chunk size must be divisible by the total len
		}

		key := s[0:t]
		allSame := true

		for i := t; i < n; i += t {
			if s[i:i+t] != key {
				allSame = false
				break
			}
		}

		if allSame {
			return true
		}
	}

	return false
}

func parseInput(filename string) IdRanges {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var idRanges IdRanges
	commaParts := strings.SplitSeq(string(data), ",")
	for parts := range commaParts {
		rangeParts := strings.Split(parts, "-")
		low, _ := strconv.Atoi(rangeParts[0])
		high, _ := strconv.Atoi(rangeParts[1])
		idRanges = append(idRanges, [2]int{low, high})
	}

	return idRanges
}
