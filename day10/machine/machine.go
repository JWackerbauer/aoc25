package machine

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
	"gonum.org/v1/gonum/stat/combin"
)

type Line struct {
	lights  []bool
	buttons [][]int
	correct []bool
}

func New(input string) *Line {
	string_sl := strings.SplitN(input, " ", 2)
	correct_lights := createLights(string_sl[0])
	string_sl = strings.SplitN(string_sl[1], " {", 2)
	string_sl = strings.Split(string_sl[0], ") (")
	buttons := [][]int{}
	for _, button := range string_sl {
		button := strings.Trim(button, "()")
		wires := []int{}
		for wire := range strings.SplitSeq(button, ",") {
			wires = append(wires, helpers.MustAtoi(wire))
		}
		buttons = append(buttons, wires)
	}
	return &Line{
		lights:  make([]bool, len(correct_lights)),
		correct: correct_lights,
		buttons: buttons,
	}
}

func createLights(input string) []bool {
	input = strings.Trim(input, "[]")
	var lights []bool
	for _, light := range input {
		switch string(light) {
		case ".":
			lights = append(lights, false)
		case "#":
			lights = append(lights, true)
		default:
			panic(fmt.Sprintf("unexpected character %c", light))
		}
	}
	return lights
}

func (line *Line) Solve(max int) (int, error) {
	for n := 1; n <= max; n++ {
		// loop through the combinations for n button presses
		for _, c := range combin.Combinations(len(line.buttons), n) {
			initial_state := make([]bool, len(line.lights))
			copy(initial_state, line.lights)
			// try the combination
			for _, button := range c {
				line.pressButton(button)
			}
			// check the outcome
			if reflect.DeepEqual(line.lights, line.correct) {
				return n, nil
			}
			// reset
			line.lights = initial_state
		}
	}
	return max, fmt.Errorf("no solution found in %v button presses", max)
}

func (line *Line) pressButton(i int) {
	for _, wire := range line.buttons[i] {
		line.lights[wire] = !line.lights[wire]
	}
}
