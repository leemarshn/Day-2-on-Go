package main

import "fmt"

func main() {
	i :=1

	for i<=5  {
		fmt.Println("Just print me several times")
		i++
	}

	for i :=1; i<=5; i++{
		fmt.Print("Cool Loop")
	}
}
