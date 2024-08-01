package main

import (
	"fmt"
	"math"
)

func askFloat(prompt string) float64 {
	var user float64
	fmt.Print(prompt)
	fmt.Scanf("%f", &user)
	return user
}

func getCompoundAmount(p, n, r float64) float64 {
	return p*math.Pow(1 + (r/100), n)
}

func main() {


	p := askFloat("What is the principal amount? ")
	n := askFloat("What is the rate? ")
	r := askFloat("What is the number of years? ")
	amount := getCompoundAmount(p, n, r)
	fmt.Printf("%.2f invested at %.2f%% for %.2f years is %.2f\n", p, r, n, amount)


}
