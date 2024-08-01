package main

import (
	"fmt"
	"bufio"
	"os"
)

func findSlice(s []string, target string) int {
	for i, v := range s {
		if v == target {
			return  i
		}
	}
	return -1
}

func printEmployees(emps []string) {
	fmt.Printf("There are %d employees\n", len(emps))
	for _, v := range emps {
		fmt.Println(v)
	}
}

func main() {

	employees := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}
	for {
		printEmployees(employees)
		fmt.Print("Enter an employee name to remove : ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		emp := sc.Text()
		pos := findSlice(employees, emp)
		if pos == -1 {
			fmt.Println("Employee does not exists")
		} else {
			employees = append(employees[:pos], employees[pos+1:]...)
		}
	}
}
