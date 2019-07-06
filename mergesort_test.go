package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name           string
		inLeft         []int
		inRight        []int
		wantSorted     []int
		wantInversions int
	}{
		// TODO: test cases
		{name: "empty", inLeft: []int{}, inRight: []int{}, wantSorted: []int{}, wantInversions: 0},
		{name: "two single", inLeft: []int{0}, inRight: []int{1}, wantSorted: []int{0, 1}, wantInversions: 0},
		{name: "single inverted", inLeft: []int{1}, inRight: []int{0}, wantSorted: []int{0, 1}, wantInversions: 1},
		{name: "single inverted", inLeft: []int{1, 1}, inRight: []int{0, 0}, wantSorted: []int{0, 0, 1, 1}, wantInversions: 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSorted, gotInversions := merge(test.inLeft, test.inRight)
			if len(gotSorted) != len(test.wantSorted) {
				t.Errorf("wrong length. got %v, wanted %v", len(gotSorted), len(test.wantSorted))
			}
			if len(test.wantSorted) > 0 && !reflect.DeepEqual(gotSorted, test.wantSorted) {
				t.Errorf("wrong array. got %v, wanted %v", gotSorted, test.wantSorted)
			}
			if gotInversions != test.wantInversions {
				t.Errorf("wrong num inversions. got %v, wanted %v", gotInversions, test.wantInversions)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		wantInversions int
	}{
		// TODO: test cases
		{name: "empty", input: []int{}, wantInversions: 0},
		{name: "two single", input: []int{0}, wantInversions: 0},
		{name: "single", input: []int{1}, wantInversions: 0},
		{name: "2 same", input: []int{1, 1}, wantInversions: 0},
		{name: "1 inv", input: []int{1, 0}, wantInversions: 1},
		{name: "4 inv", input: []int{2, 1, 3, 1, 2}, wantInversions: 4},
		{name: "4 inv", input: []int{5, 4, 3, 2, 1, 0}, wantInversions: 15},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSorted, gotInversions := mergeSort(test.input)
			sort.Ints(test.input)
			if len(test.input) > 0 && !reflect.DeepEqual(gotSorted, test.input) {
				t.Errorf("wrong sorting. got %v, wanted %v", gotSorted, test.input)
			}
			if gotInversions != test.wantInversions {
				t.Errorf("wrong num inversions. got %v, wanted %v", gotInversions, test.wantInversions)
			}
		})
	}
}
