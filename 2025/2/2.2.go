package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	Start int
	End   int
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

func parseRanges(lines []string) []Range {
	input := lines[0]
	rangeStrings := strings.Split(input, ",")
	var ranges []Range
	for _, rs := range rangeStrings {
		startEnd := strings.Split(rs, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

// new definition for part 2:
//
//	Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice.
//	So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.
func isValidId(id int) bool {
	idString := strconv.Itoa(id)

	for n := 1; n <= len(idString)/2; n++ {
		numRepititions := len(idString) / n
		newString := strings.Repeat(idString[0:n], numRepititions)
		if idString == newString {
			return false
		}
	}

	return true
}

func main() {
	// f, err := os.Open("C:/code/advent-of-code/2025/2/day_2_input.txt")
	f, err := os.Open("C:/code/advent-of-code/2025/2/day_2_test.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close()
	lines, err := parseInput(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	start := time.Now()

	var invalidIds []int
	ranges := parseRanges(lines)
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if !isValidId(id) {
				invalidIds = append(invalidIds, id)
			}
		}
	}

	var total int
	for _, id := range invalidIds {
		fmt.Println(id)
		total += id
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}
