package main

import "fmt"

func main() {

    fmt.Print("Enter number of items : ")
    var n float64
    fmt.Scanf("%f", &n)
    amt := float64(0)
    for i := 1; i < int(n+1); i++ {
        var price, quan float64
        fmt.Printf("Enter the price of item %d : ", i)
        fmt.Scanf("%f", &price)
        fmt.Printf("Enter the quantity of item %d : ", i)
        fmt.Scanf("%f", &quan)
        amt += price*quan
    }

    fmt.Printf("Subtotal: %.2f\n", amt)
    tax := amt * (5.5/100)
    fmt.Printf("Tax: %.2f\n", tax)
    fmt.Printf("Total : %.2f\n", amt+tax)

}
