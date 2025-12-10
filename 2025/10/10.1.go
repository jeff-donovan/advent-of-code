package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	diagram IndicatorLightDiagram
	buttons []Button
	// joltageRequirements
}

type IndicatorLightDiagram string

type Button []int

func makeMachines(lines []string) []Machine {
	var machines []Machine

	for _, l := range lines {
		machines = append(machines, Machine{makeDiagram(l), makeButtons(l)})
	}

	return machines
}

func makeDiagram(line string) IndicatorLightDiagram {
	leftBracket := strings.Index(line, "[")
	rightBracket := strings.Index(line, "]")
	return IndicatorLightDiagram(line[leftBracket+1 : rightBracket])
}

func makeButtons(line string) []Button {
	var buttons []Button
	lineCopy := strings.Clone(line)

	openParen := strings.Index(lineCopy, "(")
	for openParen != -1 {
		closeParen := openParen + strings.Index(lineCopy[openParen:], ")")
		buttonString := lineCopy[openParen+1 : closeParen]
		buttonAsStrings := strings.Split(buttonString, ",")
		var button Button
		for _, b := range buttonAsStrings {
			bInt, _ := strconv.Atoi(b)
			button = append(button, bInt)
		}
		buttons = append(buttons, button)

		lineCopy = lineCopy[closeParen:]
		openParen = strings.Index(lineCopy, "(")
	}

	return buttons
}

func algorithm10_1(lines []string) int {
	total := 0

	machines := makeMachines(lines)
	for _, m := range machines {
		fmt.Println(m)
	}

	return total
}
