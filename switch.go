/*
   Switch statement
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("i", i, "is")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case true:
		fmt.Println("afternoon")
	}

	myType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")

		case int:
			fmt.Println("integer")

		default:
			_ = t
		}
	}

	myType(5)
	myType(true)
}
