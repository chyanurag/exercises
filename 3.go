package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

    var quote, sayer string
    fmt.Print("What is the quote?")
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    quote = sc.Text()
    fmt.Print("Who said it?")
    sc.Scan()
    sayer = sc.Text()
    fmt.Println(sayer + " says " + "\"" + quote + "\"")


}
