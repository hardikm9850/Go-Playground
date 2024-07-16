package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    p := reflect.ValueOf(&x) // Obtain a reflect.Value of a pointer
    v := p.Elem()            // Get the reflect.Value that the pointer points to

    if v.CanSet() {
        v.SetFloat(7.1) // This will work
    } else {
        fmt.Println("Cannot set value")
    }

    fmt.Println("Modified Value:", x)
}

// ------- some theory ---------

When we use reflection to obtain a reflect.Value of a variable, the value we get is 
read-only if we don't have a pointer to it. This is because Go's reflection system 
ensures that we can only modify values if we explicitly have permission to do so, which is indicated by having a pointer.

