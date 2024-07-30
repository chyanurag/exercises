package main

import "fmt"

func main() {

    var name string
    fmt.Print("What is the input string?")
    fmt.Scanln(&name)
    fmt.Println(name + " has", len(name), "letters")


}
