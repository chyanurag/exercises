package main

import "fmt"

func main() {

    var age int
    fmt.Printf("Enter your age : ")
    fmt.Scanf("%d", &age)
    if age >= 18 {
	    fmt.Println("You are old enough to legally drive")
    } else {
	    fmt.Println("You are not old enough to legally drive")
    }


}
