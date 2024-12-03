package main

import (
	"advent24/solutions/utils"
	"errors"
	"fmt"
	"sort"
)

// calculateDistance calculates the sum of absolute differences between two sorted slices.
// It returns an error if slices are nil or empty.
func calculateDistance(s1, s2 []int) (int, error) {
	// Handle nil slices
	if s1 == nil || s2 == nil {
		return 0, errors.New("input slices cannot be nil")
	}

	// Handle empty slices
	if len(s1) == 0 || len(s2) == 0 {
		return 0, errors.New("input slices cannot be empty")
	}

	// Sort slices only if not sorted
	if !isSorted(s1) {
		sort.Ints(s1)
	}
	if !isSorted(s2) {
		sort.Ints(s2)
	}

	// Determine the shorter slice length
	minLength := len(s1)
	if len(s2) < minLength {
		minLength = len(s2)
	}

	// Calculate the sum of absolute differences
	// skipped math library because math.Abs uses float64
	sum := 0
	for i := 0; i < minLength; i++ {
		diff := s1[i] - s2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum, nil
}

// isSorted checks if a slice is sorted in ascending order.
func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func main() {
	var s1, s2, err = utils.ReadFileIntoSlices("C:/Code/advent_2024/data/1.txt", "   ")
	utils.CheckError(err)

	t, err := calculateDistance(s1, s2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(t)
	}
}
