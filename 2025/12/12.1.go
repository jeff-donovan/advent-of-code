package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Present []string

type Tree struct {
	width         int
	length        int
	presentCounts []int
}

func makePresents(lines []string) []Present {
	var presents []Present

	i := 0
	isPresent := false
	var present Present
	for _, l := range lines {
		if isPresent && l == "" {
			presents = append(presents, slices.Clone(present))
			present = nil
			isPresent = false
			i++
			continue
		}

		if !isPresent && string(l[0]) == strconv.Itoa(i) {
			isPresent = true
			continue
		}

		if isPresent {
			present = append(present, l)
		}
	}

	return presents
}

func makeTrees(lines []string) []Tree {
	var trees []Tree

	for _, l := range lines {
		xIndex := strings.Index(l, "x")
		if xIndex != -1 {
			colonIndex := strings.Index(l, ":")
			widthString := l[:xIndex]
			lengthString := l[xIndex+1 : colonIndex]
			width, _ := strconv.Atoi(widthString)
			length, _ := strconv.Atoi(lengthString)
			presentIndexesString := l[colonIndex+1:]
			presentIndexStrings := strings.Split(presentIndexesString, " ")
			var pIndexes []int
			for _, pString := range presentIndexStrings {
				pIndex, err := strconv.Atoi(pString)
				if err == nil {
					pIndexes = append(pIndexes, pIndex)
				}
			}
			trees = append(trees, Tree{width, length, pIndexes})
		}
	}

	return trees
}

func algorithm12_1(lines []string) int {
	total := 0

	presents := makePresents(lines)
	fmt.Println("num presents: ", len(presents))
	for _, p := range presents {
		fmt.Println(p)
	}

	fmt.Println()

	trees := makeTrees(lines)
	fmt.Println("num trees: ", len(trees))
	for _, t := range trees {
		fmt.Println(t)
	}

	return total
}
