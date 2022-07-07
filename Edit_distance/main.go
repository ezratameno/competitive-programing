package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistance("intention", "execution"))
}

func minDistance(word1 string, word2 string) int {
	var arr [][]int

	for i := 0; i <= len(word1); i++ {
		arr = append(arr, make([]int, len(word2)+1))
	}

	for row := 0; row <= len(word1); row++ {
		for col := 0; col <= len(word2); col++ {
			// initialize the first row and col with values
			if col == 0 {
				arr[row][col] = row
			}
			if row == 0 {
				arr[row][col] = col
			}
			if row != 0 && col != 0 {
				if word1[row-1] != word2[col-1] {
					// arr[row-1][col] - insert
					// arr[row][col-1] - delete
					// arr[row-1][col-1] - replace
					// we find the min between the  options and add 1
					arr[row][col] = int(math.Min(math.Min(float64(arr[row-1][col]),
						float64(arr[row][col-1])), float64(arr[row-1][col-1]))) + 1
				} else {
					arr[row][col] = arr[row-1][col-1]
				}
			}

		}
	}

	return arr[len(word1)][len(word2)]
}
