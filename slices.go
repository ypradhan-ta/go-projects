/*
   Slices in Go
*/

package main

import "fmt"

func main() {

	s := make([]string, 3)

	fmt.Println("empty:", s)

	s[0] = "apple"
	s[1] = "banana"
	s[2] = "chickoo"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "dessert")
	s = append(s, "egg", "fig")

	fmt.Println("append:", s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println("copy:", c)

	// slicing operations similar to python
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:6]
	fmt.Println("len:", len(l))

	fmt.Println("sl2:", l)

	l = s[3:]
	fmt.Println("sl3:", l)

	// inner slices in multi dim slices, can be of varying length
	twoD := make([][]int, 5)
	for i := 0; i < 5; i++ {
		inner := i + 1
		twoD[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
