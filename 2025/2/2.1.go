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

func isValidId(id int) bool {
	idString := strconv.Itoa(id)

	// numbers with an odd number of digits are by definition valid
	if (len(idString) % 2) != 0 {
		return true
	}

	return idString[0:len(idString)/2] != idString[len(idString)/2:]
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
		total += id
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}
