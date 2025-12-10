package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	diagram             IndicatorLightDiagram
	buttons             []Button
	requirements        JoltageRequirement
	isImpossiblePathMap map[string]struct{}
}

type IndicatorLightDiagram []bool

type Button []int

type JoltageRequirement []int

func calculateFewestButtonClicks(machine Machine) int {
	n := 1
	for {
		permutations := generatePermutations(machine.buttons, n)
		for _, p := range permutations {
			result := makeEndResult(machine, p)
			if areDiagramsEqual(machine.diagram, result) {
				return n
			}
		}
		n++
	}
}

func makeEndResult(machine Machine, clicks []Button) IndicatorLightDiagram {
	// TODO: add memoization
	generatedDiagram := make(IndicatorLightDiagram, len(machine.diagram))
	// fmt.Println("generatedDiagram before: ", generatedDiagram)
	// fmt.Println("clicks: ", clicks)
	for _, click := range clicks {
		for _, i := range click {
			generatedDiagram[i] = !generatedDiagram[i]
		}
	}
	// fmt.Println("generatedDiagram after: ", generatedDiagram)
	return generatedDiagram
}

func areDiagramsEqual(a, b IndicatorLightDiagram) bool {
	if len(a) != len(b) {
		fmt.Println("unexpectedly different lengths!")
		return false
	}

	for i, _ := range a {
		aValue := a[i]
		bValue := b[i]
		if aValue != bValue {
			return false
		}
	}

	return true
}

func generatePermutations(buttons []Button, n int) [][]Button {
	var permutations [][]Button
	if n == 1 {
		for _, b := range buttons {
			permutations = append(permutations, []Button{b})
		}
		return permutations
	}

	previousPermutations := generatePermutations(buttons, n-1)
	for _, p := range previousPermutations {
		for _, b := range buttons {
			newPermutation := slices.Concat(p, []Button{b})
			permutations = append(permutations, newPermutation)
		}
	}

	return permutations
}

func makeMachines(lines []string) []Machine {
	var machines []Machine

	for _, l := range lines {
		machines = append(machines, Machine{makeDiagram(l), makeButtons(l), makeJoltageRequirements(l), make(map[string]struct{})})
	}

	return machines
}

func makeDiagram(line string) IndicatorLightDiagram {
	var diagram []bool

	leftBracket := strings.Index(line, "[")
	rightBracket := strings.Index(line, "]")
	diagramString := line[leftBracket+1 : rightBracket]
	for _, c := range diagramString {
		if c == '.' {
			diagram = append(diagram, false)
		}
		if c == '#' {
			diagram = append(diagram, true)
		}
	}

	return diagram
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

func makeJoltageRequirements(line string) JoltageRequirement {
	var requirements []int

	leftCurlyBrace := strings.Index(line, "{")
	rightCurlyBrace := strings.Index(line, "}")
	reqString := line[leftCurlyBrace+1 : rightCurlyBrace]

	for _, char := range strings.Split(reqString, ",") {
		req, _ := strconv.Atoi(string(char))
		requirements = append(requirements, req)
	}

	return requirements
}

func parseInput(f *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
