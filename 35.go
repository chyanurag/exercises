package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
)

func main() {


	names := []string{}

	for {
		fmt.Print("Enter a name : ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		name := sc.Text()
		if name == "" {
			break
		}
		names = append(names, name)
	}
	n := rand.Intn(len(names))
	fmt.Printf("The winner is %s\n", names[n])

}
