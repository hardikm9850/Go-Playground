package main

import "fmt"

// Note the return type of the adder function
func adder() func(int) int {
	sum := 0
  // The returned function is a closure because it captures and retains access to the sum variable from the enclosing scope (adder function).
	return func(x int) int {
		sum += x
		return sum
	}
}

// How Closures Work:
// Each closure retains its own sum variable.
// The pos closure adds i to its sum in each iteration.
// The neg closure adds -2*i to its sum in each iteration.

func main() {
  // Each call to adder returns a new closure with its own separate sum variable.
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
