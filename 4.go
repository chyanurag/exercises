package main

import "fmt"

func main() {

    var noun, verb, adj, adv string
    fmt.Print("Enter a noun: ")
    fmt.Scanln(&noun)
    fmt.Print("Enter a verb: ")
    fmt.Scanln(&verb)
    fmt.Print("Enter a adjective: ")
    fmt.Scanln(&adj)
    fmt.Print("Enter a adverb: ")
    fmt.Scanln(&adv)
    fmt.Println(fmt.Sprintf("Do you %s your %s %s %s?That's Hillarious!", verb, adj, noun, adv))

}
