package math

import (
	"testing"
) 


func TestAdd(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    expected := 15
    result := Add(numbers)

    if result != expected {
        t.Errorf("Add(%v) = %d; want %d", numbers, result, expected)
    }
}

func TestSubtract(t *testing.T) {
	numbers := []int{10, 5}
		expected := 5
		result := Subtract(numbers)
	
		if result!= expected {
			t.Errorf("Subtract(%v) = %d; want %d", numbers, result, expected)
		}
}