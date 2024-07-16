package main

import (
    "errors"
    "fmt"
)

func operation1() error {
    return errors.New("\nerror in operation 1")
}

func operation2() error {
    return errors.New("error in operation 2")
}

func main() {
    
    err1 := operation1()
    err2 := operation2()

    if err1 != nil && err2 != nil {
        combinedErr := errors.Join(err1, err2)
        fmt.Println("Combined error:", combinedErr)
    }

    specificErr := errors.New("error in operation 1")
    if errors.Is(err1, specificErr) {
        fmt.Println("err1 matches specificErr")
    } else {
        fmt.Println("err1 does not match specificErr")
    }
}

// --------- o/p -------
ERROR!
Combined error: 
error in operation 1
error in operation 2
err1 does not match specificErr
