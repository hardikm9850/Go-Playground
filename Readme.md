### Learnings

##### Data types

| Data Type | Size (bits) | Example Value |
|-----------|-------------|---------------|
| `bool`    | 1           | `true`        |
| `byte`    | 8           | `0x1F`        |
| `rune`    | 32          | `'a'`         |
| `int`     | 32 or 64    | `42`          |
| `int8`    | 8           | `-128`        |
| `int16`   | 16          | `32767`       |
| `int32`   | 32          | `2147483647`  |
| `int64`   | 64          | `9223372036854775807` |
| `uint`    | 32 or 64    | `42`          |
| `uint8`   | 8           | `255`         |
| `uint16`  | 16          | `65535`       |
| `uint32`  | 32          | `4294967295`  |
| `uint64`  | 64          | `18446744073709551615` |
| `float32` | 32          | `3.14`        |
| `float64` | 64          | `3.141592653589793` |
| `complex64` | 64        | `1 + 2i`      |
| `complex128` | 128      | `1 + 2i`      |
| `string`  | variable    | `"Hello"`     |
| `array`   | variable    | `[3]int{1, 2, 3}` |
| `slice`   | variable    | `[]int{1, 2, 3}` |
| `map`     | variable    | `map[string]int{"one": 1}` |
| `struct`  | variable    | `struct{name string}{name: "Alice"}` |
| `interface` | variable  | `interface{}` |
| `chan`    | variable    | `make(chan int)` |
| `func`    | variable    | `func(x int) int { return x * 2 }` |


##### go.mod
The go.mod file is used to define modules and their dependencies in a Go project. 

To initialize a Go module, 
```
go mod init <module_name>
```
#### Interface
In Go, you implement an interface implicitly. Unlike languages like Java or C#, where you explicitly declare that a class implements an interface, in Go, a type satisfies an interface if it implements all the methods declared by that interface.
```

package main

import (
    "fmt"
)

// Define an interface named Shape
type Shape interface {
    Area() float64
}

// Define a struct named Circle
type Circle struct {
    Radius float64
}

// Define a method named Area for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Define a struct named Rectangle
type Rectangle struct {
    Width, Height float64
}

// Define a method named Area for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create a Circle object
    circle := Circle{Radius: 5}

    // Create a Rectangle object
    rectangle := Rectangle{Width: 4, Height: 3}

    // Declare a variable of type Shape
    var shape Shape

    // Assign the circle to the shape variable
    shape = circle
    fmt.Println("Circle Area:", shape.Area()) // Output: Circle Area: 78.5

    // Assign the rectangle to the shape variable
    shape = rectangle
    fmt.Println("Rectangle Area:", shape.Area()) // Output: Rectangle Area: 12
}

```

#### Function naming convention

Exported Methods: If a method name starts with a capital letter, it means the method is exported from the package and can be accessed by code outside of the package where it is defined.

Unexported Methods: If a method name starts with a lowercase letter, it means the method is unexported, and it is only accessible within the package where it is defined.
