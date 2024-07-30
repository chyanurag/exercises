package main

import (
	"fmt"
	"math"
)

const CONVERTION_FACT float64 = 0.092

func askInt(prompt string) int {
    fmt.Print(prompt)
    var user int
    fmt.Scanf("%d", &user)
    return user
}

func feetToMeter(n float64) float64 {
    return math.Sqrt(math.Pow(n, 2)*CONVERTION_FACT)   
}

func main() {

    length := askInt("What is the length of the room in feet?")
    width := askInt("What is the width of the room in feet?")
    fmt.Printf("You entered dimensions of %d feet by %d feet\n", length, width)
    fmt.Printf("The area is %d square feet and %.2f meters\n", length*width, feetToMeter(float64(length))*feetToMeter(float64(width)))

}
