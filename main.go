package main

import (
	"fmt"

	"github.com/jeff-donovan/advent-of-code/problems"
)

func main() {
	testCase := []int{1721, 979, 366, 299, 675, 1456}
	result := problems.ReportRepair(testCase)
	fmt.Println("result: ", result)
}
