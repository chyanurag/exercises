package main

import (
	"fmt"
)

func getFloat(prom string) float64 {
	fmt.Print(prom)
	var user float64
	fmt.Scanf("%f", &user)
	return user
}

func main() {

	amount := getFloat("What is the order amount? ")
	var state string
	fmt.Print("What is the state? ")
	fmt.Scanln(&state)
	if state == "WI" {
		tax := amount * 0.055
		fmt.Printf("The tax is %.2f\n", tax)
		amount += tax
	}
	fmt.Printf("The total is $%.2f\n", amount)


}
