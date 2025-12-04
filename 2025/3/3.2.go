package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func calculateLargestBankVoltage(bank string, numBatteries int) int {
	// plan
	//  - work backwards from the end of the string
	//  - calculate voltage for each index from 1 to numBatteries
	//  - for 2+ batteries you can only use values to the right
	//  - probably easiest to concatenate strings and then convert to ints

	// initialize
	var voltageMaps []map[int]int
	for _, char := range bank {
		digit, _ := strconv.Atoi(string(char))
		m := make(map[int]int)
		m[1] = digit
		voltageMaps = append(voltageMaps, m)
	}

	for n := 2; n <= numBatteries; n++ {
		for i := len(bank) - n; i >= 0; i-- {
			max := 0
			// starting from the end, which index had the best max for 1 fewer battery (than n)?
			for j := len(bank) - (n - 1); j > i; j-- {
				jMax := voltageMaps[j][n-1]
				iMax, _ := strconv.Atoi(string(bank[i]) + strconv.Itoa(jMax))
				if iMax > max {
					max = iMax
				}
			}
			voltageMaps[i][n] = max
		}
	}

	maxForNumBatteries := 0
	for _, vMap := range voltageMaps {
		val, ok := vMap[numBatteries]
		if ok && val > maxForNumBatteries {
			maxForNumBatteries = val
		}
	}
	return maxForNumBatteries
}

func main() {
	f, err := os.Open("C:/code/advent-of-code/2025/3/day_3_input.txt")
	// f, err := os.Open("C:/code/advent-of-code/2025/3/day_3_test.txt")
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

	total := 0
	for _, bank := range lines {
		total += calculateLargestBankVoltage(bank, 12)
	}

	fmt.Println("Answer: ", total)

	fmt.Println("took: ", time.Since(start))
}
