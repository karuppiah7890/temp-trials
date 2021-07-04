package solution_test

import (
	"fmt"
	"testing"

	"github.com/karuppiah7890/temp-trials/sample-coding-problems/longest-binary-gap/solution"
)

func TestSolution(t *testing.T) {
	// We could also use an array of structs with each struct containing input, expected output and a description and use the description in the test run name
	testCases := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 1, 8: 0, 9: 2, 20: 1, 33: 4, 529: 4, 1041: 5}

	for input, expectedOutput := range testCases {
		t.Run(fmt.Sprintf("Input: %d, Expected Output: %d", input, expectedOutput), func(t *testing.T) {
			actualOutput := solution.Solution(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected Output: %d, but got: %d", expectedOutput, actualOutput)
			}
		})
	}
}
