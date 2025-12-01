package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func makeCodes(f *os.File) ([]string, error) {
	var codes []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return codes, nil
}

func main() {
	f, err := os.Open("C:/code/advent-of-code/2025/1/day_1_input.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}

	defer f.Close()
	codes, err := makeCodes(f)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	start := time.Now()

	for _, code := range codes {
		fmt.Println(code)
	}

	fmt.Println("took: ", time.Since(start))
}
