package main

import "fmt"

func convertCurrency(amount, rate float64) float64 {
    ans := amount*rate
    return ans
}

func main() {

    var rupee, rate float64
    fmt.Print("How many rupees are you exchanging? ")   
    fmt.Scanf("%f", &rupee)
    fmt.Printf("What is the exchange rate? ")
    fmt.Scanf("%f", &rate)
    fmt.Printf("%.2f rupees at an exchange rate of %.2f is %.2f\n", rupee, rate, convertCurrency(rupee, rate))


}
