package main

import "fmt"

func add(x int, y int) (int, int) {

	var out = x +y
	var out2 = x -y
	return out, out2


}

func main() {

	num1 :=5
	num2 :=9

	result1, result2 := add(num2, num1)
	fmt.Print(result1, result2)


}
