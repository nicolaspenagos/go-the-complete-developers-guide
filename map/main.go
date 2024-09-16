package main

import "fmt"

func main() {

	// var colors map[string]string

	//colors := make(map[int]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#BF745",
		"white": "#fff",
	}
	// colors[10] = "#fff"
	// colors[11] = "#fff"
	delete(colors, "red")
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
