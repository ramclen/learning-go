package main

import "fmt"

func main() {
	// colors := map[string] string {
	// 	"red": "ff0000",
	// }

	// var colors map[string] string
	
	colors := make(map[string]string)
	colors["red"] = "ff0000"
	// delete(colors, "red")
	printMap(colors)
	// fmt.Println(colors)
}

func printMap (c map[string]string){
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}