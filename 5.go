package main

import (
	"fmt"
	"strconv"
)

func getNum(s string) int {
    v, err := strconv.Atoi(s)
    if err != nil {
        panic("please provide a number")
    }
    return v
}

func main() {

    var a,b string
    fmt.Print("What is the first number?")      
    fmt.Scanln(&a)
    fmt.Print("What is the second number?")      
    fmt.Scanln(&b)
    na, nb := getNum(a), getNum(b)
    fmt.Printf("%d + %d = %d\n", na, nb, na + nb)
    fmt.Printf("%d - %d = %d\n", na, nb, na - nb)
    fmt.Printf("%d * %d = %d\n", na, nb, na * nb)
    fmt.Printf("%d / %d = %d\n", na, nb, na / nb)

}
