package main

import "fmt"

func sayHello(name string) string {
    return "Hello " + name + ", nice to meet you!"
}

func main() {

    var name string
    fmt.Print("What is your name? ")
    fmt.Scanln(&name)
    fmt.Println(sayHello(name))

}
