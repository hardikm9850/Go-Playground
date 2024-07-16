package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)

    fmt.Println("Type:", t)
    fmt.Println("Value:", v)
    fmt.Println("Kind:", v.Kind())
}

// ------

Type: float64
Value: 3.4
Kind: float64
