package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
)

type Tile struct {
	x, y int
}

type Rectangle struct {
	a, b Tile
	area int
}

func process(input string) []Tile {
	var tiles []Tile
	for line := range strings.Lines(input) {
		coords := strings.Split(line, ",")
		tiles = append(tiles, Tile{
			x: helpers.MustAtoi(strings.Trim(coords[0], "\n")),
			y: helpers.MustAtoi(strings.Trim(coords[1], "\n")),
		})
	}
	return tiles
}

func tileDistance(a, b int) int {
	return 1 + int(math.Abs(float64(a-b)))
}

func area(a, b Tile) int {
	return tileDistance(a.x, b.x) *
		tileDistance(a.y, b.y)
}

func sortedRectangles(tiles []Tile) []Rectangle {
	var rectangles []Rectangle

	for i, a := range tiles {
		for _, b := range tiles[i+1:] {
			rectangles = append(rectangles, Rectangle{
				a:    a,
				b:    b,
				area: area(a, b),
			})
		}
	}

	slices.SortFunc(rectangles, func(a, b Rectangle) int {
		// Sort desc
		return a.area - b.area
	})

	return rectangles
}

func solve(input string) int {

	tiles := process(input)
	rectangles := sortedRectangles(tiles)

	return rectangles[len(rectangles)-1].area
}

type Edge struct {
	a, b Tile
}

func (me Edge) vertical() bool {
	return me.a.x == me.b.x
}

func (me Edge) horizontal() bool {
	return me.a.y == me.b.y
}

func perpendicular(e1, e2 Edge) bool {
	if e1.horizontal() && e2.horizontal() {
		return false
	}
	if e1.vertical() && e2.vertical() {
		return false
	}
	return true
}

func intersect(e1, e2 Edge) bool {
	if !perpendicular(e1, e2) {
		return false
	}
	if e1.horizontal() {
		// -> e2 is vertical
		ax_e2_bx := e1.a.x < e2.a.x && e2.a.x < e1.b.x
		bx_e2_ax := e1.b.x < e2.a.x && e2.a.x < e1.a.x
		ay_e1_by := e2.a.y < e1.a.y && e1.a.y < e2.b.y
		by_e1_ay := e2.b.y < e1.a.y && e1.a.y < e2.a.y
		return (ax_e2_bx || bx_e2_ax) && (ay_e1_by || by_e1_ay)

	} else if e1.vertical() {
		// -> e2 is horizontal
		// -> e2 is vertical
		ay_e2_by := e1.a.y < e2.a.y && e2.a.y < e1.b.y
		by_e2_ay := e1.b.y < e2.a.y && e2.a.y < e1.a.y
		ax_e1_bx := e2.a.x < e1.a.x && e1.a.x < e2.b.x
		bx_e1_ax := e2.b.x < e1.a.x && e1.a.x < e2.a.x
		return (ay_e2_by || by_e2_ay) && (ax_e1_bx || bx_e1_ax)
	} else {
		panic(fmt.Sprintf("diagonal line!? %v %v", e1, e2))
	}
}

func projectCorner(e1, e2 Edge) (Edge, Edge, error) {
	if !perpendicular(e1, e2) {
		return Edge{}, Edge{}, fmt.Errorf("Not a corner: %v %v are not perpendicular", e1, e2)
	}
	if e1.b != e2.a {
		return Edge{}, Edge{}, fmt.Errorf("Not a corner: %v %v are not the same", e1.b, e2.a)
	}
	//Project corner e1.b (= e2.a)
	var proj_corner Tile
	if e1.horizontal() {
		proj_corner = Tile{
			x: e1.a.x,
			y: e2.b.y,
		}
	} else {
		proj_corner = Tile{
			x: e2.b.x,
			y: e1.a.y,
		}
	}
	e1_p := Edge{
		e2.b,
		proj_corner,
	}
	e2_p := Edge{
		proj_corner,
		e1.a,
	}
	return e1_p, e2_p, nil
}

func solve2(input string) int {
	var result = 0

	tiles := process(input)
	edges := []Edge{}
	for i, tile := range tiles {
		next := i + 1
		if next >= len(tiles) {
			next = 0
		}
		edges = append(edges, Edge{tile, tiles[next]})
	}
	//Iterate over corners
	for i, e1 := range edges {
		next := i + 1
		if next >= len(tiles) {
			next = 0
		}
		e2 := edges[next]
		e1_p, e2_p, err := projectCorner(e1, e2)
		if err != nil {
			// Not a corner
			continue
		}
		area := area(e1.a, e2.b)
		// Be greedy
		if area > result {
			fmt.Printf("looking at %v%v%v%v area:%v", e1, e2, e1_p, e2_p, area)
			// Check if the projected edges intersect any of the existing edges.
			// If they intersect, then they are not inside the green tiles.
			contained := true
			for _, other := range edges {
				if intersect(e1_p, other) || intersect(e2_p, other) {
					contained = false
					fmt.Printf(" intersects %v", other)
					break
				}
			}
			if contained {
				result = area
			}
			fmt.Printf("\n")
		}
	}

	return result
}

var test_input = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func main() {
	input, err := os.ReadFile("day09/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	var number = "test"
	if len(os.Args) > 1 {
		number = os.Args[1]
	}

	switch number {
	case "1":
		println(solve(string(input)))
	case "2":
		println(solve2(string(input)))
	default:
		println("test 1:")
		println(solve(test_input))
		println("test 2:")
		println(solve2(test_input))
	}
}
