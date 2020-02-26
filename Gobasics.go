package main

import "fmt"

func main() {

	//var a int = 10
	a := 22
	if a < 20 {
		fmt.Println("A is less than 20")
	} else if a > 20 {
		fmt.Println("A is greater than 20")
	}
	var grade = ""

	var marks = 30
	switch marks {

	case 100:
		grade = "A"
	case 90:
		grade = "B"
	default:
		grade = "fail"
	}
	fmt.Println("Grade is " + grade)

}
