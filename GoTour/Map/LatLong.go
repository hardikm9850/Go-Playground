package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//Declare
var mapOfLocation map[string]Vertex

func main() {
	// Intialise	
	mapOfLocation = make(map[string]Vertex)
	// Add element
	mapOfLocation["location"] = Vertex {
		1.33,3.44,
	}
	fmt.Println(mapOfLocation) // map[location:{1.33 3.44}]
	fmt.Println(mapOfLocation["location"]) // {1.33 3.44}
	fmt.Println(mapOfLocation["location_unknown"]) // {0 0}
}
