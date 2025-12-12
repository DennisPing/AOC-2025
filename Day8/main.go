package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point3D struct {
	X, Y, Z float64
}

type Edge struct {
	I, J     int // The indices of the edge relative to the given points
	Distance float64
}

func main() {
	points := parseInput("input.txt")

	edges := buildAllEdges(points)

	// Sort from least to greatest edge distance
	slices.SortFunc(edges, func(a, b Edge) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	part1(edges)
	part2(points, edges)
}

func part1(edges []Edge) {
	// Find chains of the closest N edges
	chains := findChains(edges[:1000])
	//chains := findChains(edges[:10]) // Use this for "test.txt"

	// Sort from greatest to least
	slices.SortFunc(chains, func(a, b []int) int {
		return -cmp.Compare(len(a), len(b))
	})

	ans := 1
	for _, chain := range chains[:3] {
		ans *= len(chain)
	}

	fmt.Println(ans)
}

func part2(points []Point3D, edges []Edge) {
	// Track which chain each edge belongs to
	// At first, each point belongs to its own chain
	chainIds := make([]int, len(points))
	for i := 0; i < len(points); i++ {
		chainIds[i] = i
	}

	numChains := len(points)
	var lastMerge Edge

	for _, edge := range edges {
		chainI := chainIds[edge.I]
		chainJ := chainIds[edge.J]

		if chainI == chainJ {
			continue
		}

		// Merge the two ID's into one
		for k := 0; k < len(points); k++ {
			if chainIds[k] == chainJ {
				chainIds[k] = chainI
			}
		}

		numChains--
		lastMerge = edge

		if numChains == 1 {
			break
		}
	}

	p1 := points[lastMerge.I]
	p2 := points[lastMerge.J]

	ans := int(p1.X * p2.X)
	fmt.Println(ans)
}

// Euclidean distance
func dist3d(p1, p2 Point3D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z

	sumSq := dx*dx + dy*dy + dz*dz
	return math.Sqrt(sumSq)
}

// Build all possible edges from i to j
func buildAllEdges(points []Point3D) []Edge {
	n := len(points)
	edges := make([]Edge, 0, n*(n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := dist3d(points[i], points[j])
			edges = append(edges, Edge{
				I:        i,
				J:        j,
				Distance: dist,
			})
		}
	}

	return edges
}

// Find the connected chains
func findChains(edges []Edge) [][]int {
	// Build an adjacency list
	adj := make(map[int][]int)
	for _, edge := range edges {
		i, j := edge.I, edge.J
		adj[i] = append(adj[i], j)
		adj[j] = append(adj[j], i)
	}

	visited := make(map[int]bool)
	chains := make([][]int, 0)

	// DFS
	for k := 0; k < len(adj); k++ {
		if visited[k] {
			continue
		}

		// Start a new chain from k
		stack := []int{k}
		visited[k] = true
		var component []int

		for len(stack) > 0 {
			idx := stack[len(stack)-1]   // Pop from top
			stack = stack[:len(stack)-1] // Remove top

			component = append(component, idx)

			// Traverse through all neighbors
			for _, nb := range adj[idx] {
				if !visited[nb] {
					visited[nb] = true
					stack = append(stack, nb)
				}
			}
		}

		chains = append(chains, component)
	}

	return chains
}

func parseInput(filename string) []Point3D {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		log.Fatalln("unable to read file:", filename)
	}

	var points []Point3D
	for line := range strings.SplitSeq(string(data), "\n") {
		fields := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(fields[0], 64)
		y, _ := strconv.ParseFloat(fields[1], 64)
		z, _ := strconv.ParseFloat(fields[2], 64)

		points = append(points, Point3D{x, y, z})
	}

	return points
}
