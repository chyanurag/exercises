package main

import "fmt"

func main() {

    
	var user, pass string
	fmt.Print("Enter the username : ")
	fmt.Scanln(&user)
	fmt.Print("Enter the password : ")
	fmt.Scanln(&pass)
	if pass == "password" {
		fmt.Println("You are welcome!")
	} else {
		fmt.Println("I don't even know who you are")
	}

}
