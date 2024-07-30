package main

import (
	"fmt"
	"math"
)

func askInt(prompt string) int {
    fmt.Print(prompt)
    var user int
    fmt.Scanf("%d", &user)
    return user
}

func main() {

    length := askInt("Enter length in feet : ")
    width := askInt("Enter width in feet : ")
    area := float64(length*width)
    fmt.Printf("You need to purchase %.2f gallons of paint to  cover %.2f square feet\n", math.Ceil(float64(area/350)),area)

}
