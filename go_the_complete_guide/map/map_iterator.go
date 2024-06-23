package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	delete(colors, "red")
	fmt.Println(colors)

	_colors := make(map[string]string)
	_colors["whie"] = "#ffffff"
	printMap(_colors)
}

func printMap(colorMap map[string]string) {
	for color, value := range colorMap {
		fmt.Println("Hex code for", color, "is", value)
	}
}
