package main

import (
	"fmt"
	"time"
)

func askInt(prompt string) int {
    fmt.Printf(prompt)
    var user int
    fmt.Scanf("%d", &user)
    return user
}

func main() {

       age := askInt("What is your current age?")
       rage := askInt("At what age you want to retire?")
       t := rage - age
       fmt.Printf("You have %d years left until you can retire\n", t)
       currYear := time.Now().Year()
       fmt.Printf("It's %d, so you can retire in %d.\n", currYear, currYear + t)

}
