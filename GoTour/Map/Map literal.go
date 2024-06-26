/**
A literal is a source code representation of a fixed value. Literals represent constant values directly in the code.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//If the top-level type is just a type name, in our case Vertex, you can omit it from the elements of the literal.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
var mapLiteral = map[string]Vertex {
	"Pune" : Vertex {
		1.1 , 2.2,
	},
	"BLR" : Vertex {
		21.1 , 32.2,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(mapLiteral)
}
