package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// When you call fmt.Sprint(e) inside the Error method, fmt.Sprint checks if e implements the error interface. 
  // Since e does implement the error interface (through the Error method), fmt.Sprint 
  // will call the Error method on e to get its string representation leading to SOF

  // return fmt.Sprint(e) // causes stackoverflow
	
	return fmt.Sprintf("cannot Sqrt negative number: %f",e)
}

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return x,ErrNegativeSqrt(x)
	}
	var z float64 = 1
	 for z*z < x {
		z = z + 1
	 }
	return z-1, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
