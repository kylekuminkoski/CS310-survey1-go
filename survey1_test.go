//Co-authors: Kyle Kuminkoski and Grant Stumpf
//I confirm that all work follows the academic integrity policy.

package main

import "testing"

func TestHyper(t *testing.T) {
	factorialMap := make(map[int]int)

	var solution = [4][11]int{
		{1, 2, 4, 7, 11, 16, 22, 29, 37, 46, 56},
		{1, 2, 4, 8, 15, 26, 42, 64, 93, 130, 176},
		{1, 2, 4, 8, 16, 31, 57, 99, 163, 256, 386},
		{1, 2, 4, 8, 16, 32, 63, 120, 219, 382, 638},
	}

	var i, j int

	for i = 0; i < 4; i++ {
		for j = 0; j < 11; j++ {
			var k = i + 2
			Test := hypercake(j, k, factorialMap)
			Solution := solution[i][j]

			if Test != Solution {
				t.Errorf("Failure; n-value: %d, k-value: %d, got %d, wanted %d", j, k, Test, Solution)
			} else {
				t.Logf("Success; n-value: %d, k-value: %d, Solution: %d", j, k, Solution)
			}

		}
	}
}
