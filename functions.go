/*
   functions, closures
*/

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addMany(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func squareSeq() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	// sum := add(5, 10)

	// sum := addMany(1, 2, 3, 4, 5)
	// fmt.Println("sum:", sum)

	squares := squareSeq()

	fmt.Println(squares())
	fmt.Println(squares())
	fmt.Println(squares())
	fmt.Println(squares())
	fmt.Println(squares())

	newSq := squareSeq()
	fmt.Println(newSq())
	fmt.Println(newSq())

}
