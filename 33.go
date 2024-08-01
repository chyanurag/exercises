package main

import (
	"math/rand"
	"fmt"
	"os"
	"bufio"
)

func main() {

	fmt.Printf("What's your question? ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	answers := [4]string{"Yes", "No", "Maybe", "Ask again later"}
	n := rand.Intn(4)
	fmt.Println(answers[n])

}
