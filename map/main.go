package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(arr map[string]string) {
	for key, value := range arr {
		fmt.Println(key, value)
	}
}
