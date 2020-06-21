package main

import "fmt"

func main() {

	num :=2

	if num <5 {

		fmt.Println("Hey", num)


	}else if num == 2{
		fmt.Println("Good Afternoon, ", num)

	}

	num2 :=2

	switch num2 {

	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("none")


	}




}
