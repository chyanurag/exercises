package main

import (
	"fmt"
	"os"
	"strconv"
)

// pizza slices

func askInt(prompt string) int {
    var user string
    fmt.Print(prompt)
    fmt.Scanln(&user)
    n, err := strconv.Atoi(user)
    if err != nil {
        fmt.Print("Please input a number\n")
        os.Exit(0)
    }
    return n
}

func main() {

    people := askInt("How many people?")   
    pizzas := askInt("How many pizzas do you have?")
    slices := pizzas * 8
    fmt.Printf("%d people with %d pizzas\n", people, pizzas)
    fmt.Printf("Each person gets %d pieces of pizza.\n", slices/people)
    fmt.Printf("There are %d leftover pieces", slices%people)

}
