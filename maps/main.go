package main

import "fmt"

func main() {
	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["white"] = "#ffffff"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"
	colors["black"] = "#000000"

	print(colors)
}

func print(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for ", key, " is ", value)
	}
}
