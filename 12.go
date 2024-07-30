package main

import "fmt"

func main() {

    var principal, rate float64
    var years int
    fmt.Print("Enter the principal : ")
    fmt.Scanf("%f", &principal)
    fmt.Print("Enter the rate : ")
    fmt.Scanf("%f", &rate)
    fmt.Print("Enter the years : ")
    fmt.Scanf("%d", &years)
    ints := (principal*rate*float64(years))/100
    fmt.Printf("After %d years at %.2f%%, the investment will be worth %.2f\n", years, rate, principal + ints)


}
