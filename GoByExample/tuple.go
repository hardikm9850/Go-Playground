package main

import "fmt"

//Tuple assignment
func main() {
	var fibbonacci = fib(10)
	fmt.Print(fibbonacci)
	fmt.Println()
}

func fib(number int) int {
	x, y := 0, 1
	for i := 0; i < number; i++ {
		x, y = y, y+x
	}
	return x
}
