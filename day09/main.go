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
		return b.area - a.area
	})

	return rectangles
}

func solve(input string) int {

	tiles := process(input)
	rectangles := sortedRectangles(tiles)

	return rectangles[0].area
}

type Edge struct {
	a, b Tile
}

func (e Edge) normalize() Edge {
	// Check if e.a is "larger" than e.b
	if e.a.x > e.b.x || (e.a.x == e.b.x && e.a.y > e.b.y) {
		return Edge{a: e.b, b: e.a}
	}
	return e
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

func between(a, b, between int) bool {
	return a < between && between < b ||
		b < between && between < a
}

func between_incl(a, b, between int) bool {
	return a <= between && between <= b ||
		b <= between && between <= a
}

func between_lower_incl(a, b, between int) bool {
	return a <= between && between < b ||
		b <= between && between < a
}

func edgeContains(e Edge, t Tile) bool {
	var cont_x, cont_y bool
	if e.horizontal() {
		cont_x = between_incl(e.a.x, e.b.x, t.x)
		cont_y = t.y == e.a.y
	} else {
		cont_x = t.x == e.a.x
		cont_y = between_incl(e.a.y, e.b.y, t.y)
	}
	return cont_x && cont_y
}

func intersect_incl(e1, e2 Edge) bool {
	if !perpendicular(e1, e2) {
		return false
	}
	if e1.horizontal() {
		// -> e2 is vertical
		return between(e1.a.x, e1.b.x, e2.a.x) && between_lower_incl(e2.a.y, e2.b.y, e1.a.y)
	} else if e1.vertical() {
		// -> e2 is horizontal
		// -> e2 is vertical
		return between(e1.a.y, e1.b.y, e2.a.y) && between_lower_incl(e2.a.x, e2.b.x, e1.a.x)
	} else {
		panic(fmt.Sprintf("diagonal line!? %v %v", e1, e2))
	}
}

func oppositeCorners(t1, t3 Tile) (Tile, Tile) {
	t2 := Tile{x: t1.x, y: t3.y}
	t4 := Tile{x: t3.x, y: t1.y}
	return t2, t4
}

func edgesOfCorners(t1, t2, t3, t4 Tile) (Edge, Edge, Edge, Edge) {
	return Edge{t1, t2}, Edge{t2, t3}, Edge{t3, t4}, Edge{t4, t1}
}

var greenTiles = make(map[Tile]bool)

func isInsideEdges(t Tile, edges []Edge) bool {
	memo, ok := greenTiles[t]
	if ok {
		return memo
	}
	t_edge := Edge{Tile{x: 0, y: t.y}, t}
	intersections := 0
	for _, e := range edges {
		if edgeContains(e, t) {
			greenTiles[t] = true
			return true
		}
		if intersect_incl(t_edge, e) {
			intersections++
		}
	}
	inside := intersections%2 != 0
	//greenTiles[t] = inside
	return inside
}

var greenEdges = make(map[Edge]bool)

func edgeIsInsideEdges(e Edge, edges []Edge) bool {
	e = e.normalize()
	memo, ok := greenEdges[e]
	if ok {
		return memo
	}
	if e.horizontal() {
		x1, x2 := e.a.x, e.b.x
		y := e.a.y
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			if !isInsideEdges(Tile{x: x, y: y}, edges) {
				greenEdges[e] = false
				return false
			}
		}
	} else {
		y1, y2 := e.a.y, e.b.y
		x := e.a.x
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			if !isInsideEdges(Tile{x: x, y: y}, edges) {
				greenEdges[e] = false
				return false
			}
		}
	}
	greenEdges[e] = true
	return true
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
		e := Edge{tile, tiles[next]}

		edges = append(edges, e)

	}
	rectangles := sortedRectangles(tiles)
	//Iterate over corners
	for _, r := range rectangles {
		t2, t4 := oppositeCorners(r.a, r.b)
		e1, e2, e3, e4 := edgesOfCorners(r.a, t2, r.b, t4)
		// Check if the opposite corners are inside the green tiles
		if !isInsideEdges(t2, edges) {
			continue
		}
		if !isInsideEdges(t4, edges) {
			continue
		}
		// Check if all the edges are fully inside the green tiles
		if edgeIsInsideEdges(e1, edges) &&
			edgeIsInsideEdges(e2, edges) &&
			edgeIsInsideEdges(e3, edges) &&
			edgeIsInsideEdges(e4, edges) {
			result = r.area
			break
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
