package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
)

// Node
type Junction struct {
	x, y, z int
}

// Edge
type JunctionPair struct {
	a, b     Junction
	distance float64
}

func (me Junction) DistanceTo(other Junction) float64 {
	return math.Sqrt(
		math.Pow(float64(other.x-me.x), 2) +
			math.Pow(float64(other.y-me.y), 2) +
			math.Pow(float64(other.z-me.z), 2))
}

func process(input string) ([]JunctionPair, int) {

	var nodes []Junction
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		coords := strings.Split(line, ",")
		junction := Junction{
			x: helpers.MustAtoi(strings.Trim(coords[0], "\n")),
			y: helpers.MustAtoi(strings.Trim(coords[1], "\n")),
			z: helpers.MustAtoi(strings.Trim(coords[2], "\n")),
		}
		nodes = append(nodes, junction)
	}
	var edges []JunctionPair
	for i, self := range nodes {
		for _, other := range nodes[i+1:] {
			edges = append(edges, JunctionPair{
				a:        self,
				b:        other,
				distance: self.DistanceTo(other),
			})
		}
	}
	slices.SortFunc(edges, func(a JunctionPair, b JunctionPair) int {
		return int(a.distance - b.distance)
	})
	return edges, len(lines)
}

func tryConnectNode(circuit []Junction, edge JunctionPair) (*Junction, bool) {
	connected_a := slices.Contains(circuit, edge.a)
	connected_b := slices.Contains(circuit, edge.b)
	if connected_a && connected_b {
		return nil, true
	} else if connected_a {
		fmt.Printf("added %v -> %v @ %v\n", edge, circuit, edge.b)
		return &edge.b, true
	} else if connected_b {
		fmt.Printf("added %v -> %v @ %v\n", edge, circuit, edge.a)
		return &edge.a, true
	} else {
		return nil, false
	}
}

func mergeCircuits(circuits [][]Junction, added_node *Junction, added_index int) [][]Junction {
	if added_node != nil {
		// check for circuits to merge
		for j, circuit := range circuits {
			// skip self
			if j == added_index {
				continue
			}
			if slices.Contains(circuit, *added_node) {
				fmt.Printf("merge %v %v\n", circuits[added_index], circuits[j])
				circuits[j] = append(circuits[j], circuits[added_index][:len(circuits[added_index])-1]...)
				circuits = slices.Delete(circuits, added_index, added_index+1)
				break
			}
		}
	}
	return circuits
}

func solve(input string, num_connections int) int {
	var circuits [][]Junction

	edges, _ := process(input)

	limit := num_connections
	for i, edge := range edges {
		if i >= limit {
			break
		}
		added_index := -1
		var added_node *Junction
		added := false
		for j, circuit := range circuits {
			added_node, added = tryConnectNode(circuit, edge)
			if added_node != nil {
				added_index = j
				circuits[j] = append(circuits[j], *added_node)
			}
			if added {
				break
			}
		}
		circuits = mergeCircuits(circuits, added_node, added_index)
		if !added {
			// Otherwise add new circuit with single connection
			fmt.Printf("new %v\n", edge)
			circuits = append(circuits, []Junction{edge.a, edge.b})
		}
	}

	fmt.Printf("\n")

	slices.SortFunc(circuits, func(a []Junction, b []Junction) int {
		// sort descending
		return len(b) - len(a)
	})

	circuit_amt := 3
	result := 1
	for i, circuit := range circuits {
		nodes := len(circuit)
		fmt.Printf("%v size: %v\n", circuit, nodes)
		if i < circuit_amt {
			result *= nodes
		}
	}

	return result
}

func solve2(input string) int {
	var circuits [][]Junction

	result := 0

	edges, num_nodes := process(input)

	for _, edge := range edges {

		added_index := -1
		var added_node *Junction
		added := false
		for j, circuit := range circuits {
			added_node, added = tryConnectNode(circuit, edge)
			if added_node != nil {
				added_index = j
				circuits[j] = append(circuits[j], *added_node)
			}
			if added {
				break
			}
		}
		circuits = mergeCircuits(circuits, added_node, added_index)
		if !added {
			// Otherwise add new circuit with single connection
			fmt.Printf("new %v\n", edge)
			circuits = append(circuits, []Junction{edge.a, edge.b})
		}
		if len(circuits[0]) >= num_nodes {
			fmt.Printf("Final connection: %v\n", edge)
			return edge.a.x * edge.b.x
		}
	}
	return result
}

var test_input = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func main() {
	input, err := os.ReadFile("day08/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	var number = "test"
	if len(os.Args) > 1 {
		number = os.Args[1]
	}

	switch number {
	case "1":
		println(solve(string(input), 1000))
	case "2":
		println(solve2(string(input)))
	default:
		println("test 1:")
		println(solve(test_input, 10))
		println("test 2:")
		println(solve2(test_input))
	}
}
