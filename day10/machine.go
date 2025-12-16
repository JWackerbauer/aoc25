package main

import (
	"fmt"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
)

type Machine struct {
	lights  []bool
	buttons [][]int
	correct []bool
}

func CreateMachine(input string) Machine {
	string_sl := strings.SplitN(input, " ", 2)
	correct_lights := CreateLights(string_sl[0])
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
	machine := Machine{
		lights:  make([]bool, len(correct_lights)),
		correct: correct_lights,
		buttons: buttons,
	}
	fmt.Printf("%v\n", machine)
	return machine
}

func CreateLights(input string) []bool {
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
