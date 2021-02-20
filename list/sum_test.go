package main

import "testing"

func TestSum(t *testing.T){
	numbers := [...]int{1,2,3}
	got := Sum(numbers)
	expected := 6
	if got != expected{
		t.Errorf("expected %d but got %d, numbers: %v", expected, got, numbers)
	}
}

func Sum(numbers [3]int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += numbers[i]
	}
	return sum
}