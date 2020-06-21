package main

import "fmt"

var a = 9 //package level variable available in all functions

func demo(){
	var a  = 91025 //function level variable
	fmt.Println(a)
}

func main() {
	demo() // prints 91025

	fmt.Print(a)

	
}
