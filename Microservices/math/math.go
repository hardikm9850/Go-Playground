package math

func Add(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}


func Subtract(s []int) int {
	result := s[0] - s[1]
	return result
}

func Foo() {
	// This function does nothing
    // and serves as a placeholder for a future implementation
    // of other math functions
    
}