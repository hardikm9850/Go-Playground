package main

import "fmt"
// Type Assertions and Type Switches
// We can use type assertions or type switches to determine the actual type stored in an empty interface

func main() {
    var i interface{} = 3.14
    
    switch v := i.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    case float64:
        fmt.Println("float64:", v)
    default:
        fmt.Println("unknown type")
    }

    var j interface{} = "hello"

    s, ok := j.(string)
    if ok {
        fmt.Println("string:", s)
    } else {
        fmt.Println("not a string")
    }
}
