package gobyexample

import (
	"fmt"
	"math"
)

func main() {
	explainNew()
}

func explainNew() {
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}

func lifeTimeOfVariables() {
	for t := 0.0; t < 5; t += 1 {
		 x := math.Sin(t)
		 y := math.Cos(t)
	}
}
