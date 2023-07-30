package main

import "fmt"

func main() {
	var fruits = make([]string, 0, 5)
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "strawberry")

	// remove apple
	fruits = append(fruits[:0], fruits[1:]...)

	fmt.Printf("%+v\n", fruits)

}
