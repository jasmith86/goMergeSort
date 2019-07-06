package main

import "fmt"

// Merge two arrays, returning sorted array and number of inversions.
func merge(l, r []int) ([]int, int) {
	var srtd []int
	var inv int
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			srtd = append(srtd, l[i])
			i += 1
		} else {
			srtd = append(srtd, r[j])
			j += 1
			inv += len(l) - i
		}
	}
	srtd = append(srtd, l[i:]...)
	srtd = append(srtd, r[j:]...)
	return srtd, inv
}

// Use mergesort to get the sorted array and number of inversions.
func MergeSort(arr []int) ([]int, int) {
	if len(arr) < 2 {
		return arr, 0
	}
	l, li := MergeSort(arr[:len(arr)/2])
	r, ri := MergeSort(arr[len(arr)/2:])
	rv, ti := merge(l, r)
	return rv, li + ri + ti
}

func main() {
	a := []int{2, 4, 1, 0}
	sorted, inversions := MergeSort(a)
	fmt.Println(sorted)
	fmt.Println(inversions)
}
