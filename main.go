package main

import (
	"fmt"
	"os"
)

type Item struct {
	name  string
	id    string
	price float64
	left  int
}

var menu []Item

type Order struct {
	name    string
	id      string
	numbers int
	total   float64
}

var the_check []Order

// variable
var Fish Item
var Pasta Item
var Chicken Item

func main() {
	//TO_DO : create a name and name id , name decribe and name id is used to order stuff

	// set fish
	Fish.name = "fish with its potato"
	Fish.id = "fish1"
	Fish.price = 12.0
	Fish.left = 10

	// set pasta
	Pasta.name = "pasta a la parmesane"
	Pasta.id = "pasta1"
	Pasta.price = 10.0
	Pasta.left = 10

	// set chicken
	Chicken.name = "grilled chicken with fries"
	Chicken.id = "chicken1"
	Chicken.price = 13.0
	Chicken.left = 10

	// add element to the menu
	menu = append(menu, Fish, Pasta, Chicken)

	//START LOGIC
	var to_do string

	for {
		fmt.Println("what to do , a : print menu, b: order something ,c: print check, z : exit")
		fmt.Scan(&to_do)

		switch to_do {
		case "a":
			Print_items(menu)
		case "b":
			long_name, id_name, nb_pl, tot_price := Return_order(menu)
			//create type order
			var the_order Order
			the_order = Create_order(long_name, id_name, nb_pl, tot_price)
			the_check = append(the_check, the_order)
			fmt.Println(the_check)
		case "c":
			Print_check(the_check)

		case "z":
			os.Exit(3)
		}
	}

	// Return_order(menu)

}

func Print_items(items []Item) {
	for _, i := range items {
		//if there is no item left dont print it
		if i.left == 0 {
			continue
		}

		// print waring if there is not a lot of item left
		if i.left < 5 {
			fmt.Println(i.name, "cost :", i.price, " euros // id_name : ", i.id)
			fmt.Println("be carefull there is only", i.left, " left \n")
		} else {
			fmt.Println(i.name, "cost :", i.price, "euros // id_name : ", i.id, "\n")
		}

	}

}

func Create_order(the_name string, the_id string, the_numbers int, the_total float64) Order {
	//variable
	var val Order

	val.name = the_name
	val.id = the_id
	val.numbers = the_numbers
	val.total = the_total

	return val
}

// return element to be added to a order struct, name , numbers of plate , total price
func Return_order(the_menu []Item) (string, string, int, float64) {
	//create variable
	var name_plate string
	var long_name string
	var the_price float64

	for {
		Print_items(menu)
		fmt.Println("What do you want to order?")
		fmt.Scan(&name_plate)
		//
		exist := false
		for _, i := range the_menu {
			if i.id == name_plate {
				exist = true
				the_price = i.price
				long_name = i.name
			}
		}

		if exist {
			break
		} else {
			fmt.Println("could fint your order!!!")
		}

	}

	//ask number
	var number int
	fmt.Println("How many you want to order?")
	fmt.Scan(&number)

	//create total
	var total float64

	total = float64(number) * the_price

	// return the elements
	return long_name, name_plate, number, total

}

// print the final check
func Print_check(the_final_check []Order) {
	//create variable
	var the_final_total float64

	//print all element in check
	for _, e := range the_final_check {
		the_final_total += e.total
		fmt.Printf("%v : %v  %.1f  \n", e.name, e.numbers, e.total)
	}
	fmt.Println("total is :", the_final_total)

}
