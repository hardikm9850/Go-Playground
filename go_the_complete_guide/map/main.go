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
	fmt.Println(_colors)
}
