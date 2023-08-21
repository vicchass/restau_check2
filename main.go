package main

import "fmt"

type Item struct {
	name  string
	price float64
	left  int
}

//VARIABLE
var Fish Item
var Pasta Item
var Chicken Item

var menu []Item

func main() {
	// set fish
	Fish.name = "grilled fish"
	Fish.price = 12.0
	Fish.left = 10

	// set pasta
	Pasta.name = "pesto pasta"
	Pasta.price = 10.0
	Pasta.left = 10

	// set chicken
	Chicken.price = 13.0
	Chicken.name = "chicken parmesan"
	Chicken.left = 3

	menu = append(menu, Fish, Pasta, Chicken)

	Print_items(menu)

}

func Print_items(items []Item) {
	for _, i := range items {
		fmt.Println("there is", i.name, " on the menu")

	}

}
