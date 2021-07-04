package solution_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/karuppiah7890/temp-trials/sample-coding-problems/array-rotation/solution"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		array          []int
		k              int
		expectedOutput []int
	}{

		{
			array:          []int{},
			k:              1,
			expectedOutput: []int{},
		},
		{
			array:          []int{3, 8, 9, 7, 6},
			k:              1,
			expectedOutput: []int{6, 3, 8, 9, 7},
		},
		{
			array:          []int{3, 8, 9, 7, 6},
			k:              3,
			expectedOutput: []int{9, 7, 6, 3, 8},
		},
		{
			array:          []int{0, 0, 0},
			k:              1,
			expectedOutput: []int{0, 0, 0},
		},
		{
			array:          []int{1, 2, 3, 4},
			k:              4,
			expectedOutput: []int{1, 2, 3, 4},
		},
		{
			array:          []int{1, 2, 3, 4},
			k:              100,
			expectedOutput: []int{1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Array: %v, K: %d, Expected Output: %d", testCase.array, testCase.k, testCase.expectedOutput), func(t *testing.T) {
			actualOutput := solution.Solution(testCase.array, testCase.k)

			if !reflect.DeepEqual(actualOutput, testCase.expectedOutput) {
				t.Errorf("Expected Output: %d, but got: %d", testCase.expectedOutput, actualOutput)
			}
		})
	}
}
