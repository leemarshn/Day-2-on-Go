package main

import (
	"fmt"
	"math"
)

func main() {

	num :=92.0

	var result = math.Sqrt(num) // sqrt only works with float numbers
	fmt.Println(result)
	fmt.Printf("The output is %.2f ", result)// printf used when you want to specify format
}
