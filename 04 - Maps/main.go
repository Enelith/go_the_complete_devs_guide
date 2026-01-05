package main

import "fmt"

func main() {
	// 1) Different ways to declare a map
	// The 0 value of a map is an empy map (map[])
	var colors0value_v1 map[string]string
	fmt.Println(colors0value_v1)
	// Ouput: map[]

	// "make" is a built-in function that returns a new initialized map
	colors0value_v2 := make(map[string]string)
	fmt.Println(colors0value_v2)
	// Ouput: map[]

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
	// Ouput: map[blue:#0000ff green:#00ff00 red:#ff0000]

	// 2) Add an element to a Map
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	// Ouput: map[blue:#0000ff green:#00ff00 red:#ff0000 white:#ffffff]

	// 3) Delete an element from a Map
	delete(colors, "blue")
	fmt.Println(colors)
	// Ouput: map[green:#00ff00 red:#ff0000 white:#ffffff]

	// 4) Iterate over a Map (through a function)
	printMap(colors)
	// Ouput:
	// green: #00ff00
	// white: #ffffff
	// red: #ff0000
}

// Helper function
// "colors" is the argument name. "map[string]string" is the type of the map / of the argument.
func printMap(colors map[string]string) {
	// color is the key for this step through the loop
	// hex is the value for this step through the loop
	for color, hex := range colors {
		fmt.Printf("%s: %s\n", color, hex)
	}
}
