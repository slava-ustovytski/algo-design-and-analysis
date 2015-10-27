package main

import "fmt"

func main() {
	// unsorted := []int{5, 4, 1, 8, 7, 2, 6, 3}
	unsorted := []int{8, 7, 2, 1}
	sorted := Sort(unsorted)
	fmt.Println(unsorted, " => ", sorted)
}

func Sort(x []int) []int {
	length := len(x)
	if length <= 1 {
		return x
	}

	half := length / 2
	left := x[:half]
	right := x[half:]

	if length > 2 {
		left = Sort(left)
		right = Sort(right)
	}
	sorted := merge(left, right)
	
	return sorted
}

func merge(left, right []int) []int {
	var sorted []int
	i, j := 0, 0
	for i<len(left) && j<len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else  {
			sorted = append(sorted, right[j])
			j++
		}
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
	}

	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
	}

	return sorted
}
