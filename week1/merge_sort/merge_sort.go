package main

import "fmt"

func main() {
	unsorted := []int{5, 4, 1, 8, 7, 2, 6, 3}
	sorted := Sort(unsorted)
	fmt.Println(unsorted, " => ", sorted)
}

func Sort(x []int) []int {

	half := len(x) / 2
	left := x[:half]
	right := x[half:len(x)]

	if len(x) > 2 {
		left = Sort(left)
		right = Sort(right)
	}
	sorted := merge(left, right)
	return sorted
}

func merge(left, right []int) []int {
	var length int
	length = len(left)
	var sorted []int
	for i := 0; i < length; i++ {
		if left[i] < right[i] {
			sorted = append(sorted, left[i])
			sorted = append(sorted, right[i])
		} else {
			sorted = append(sorted, right[i])
			sorted = append(sorted, left[i])
		}
	}

	return sorted
}
